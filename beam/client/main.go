package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"flag"

	"github.com/google/uuid"
	"github.com/ldsec/lattigo/bfv"
)

// go run main.go  -mode=encrypt --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/encrypted.bin
// go run main.go  -mode=decrypt --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/encrypted.bin
func main() {
	mode := flag.String("mode", "", "(required for mode=mode) mode=encrypt|decrypt")
	pkFile := flag.String("pkFile", "", "(required for mode=genkey) Filename for the rider's publicKey")
	skFile := flag.String("skFile", "", "(required for mode=genkey) Filename for the rider's secrettKey")
	encryptedFile := flag.String("encryptedFile", "", "(required for mode=encrypt) Filename for to save the encryptedFile")

	flag.Parse()

	params := bfv.DefaultParams[bfv.PN12QP109]
	//params.T = //0x3ee0001 //65929217₁₀
	encoder := bfv.NewEncoder(params)

	if *mode == "encrypt" {
		kgen := bfv.NewKeyGenerator(params)
		riderSk, riderPk := kgen.GenKeyPair()
		pubBytes, err := riderPk.MarshalBinary()
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(*pkFile, pubBytes, 0640)
		if err != nil {
			panic(err)
		}

		secBytes, err := riderSk.MarshalBinary()
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(*skFile, secBytes, 0640)
		if err != nil {
			panic(err)
		}
		encryptorPk := bfv.NewEncryptorFromPk(params, riderPk)

		file, err := os.Create(*encryptedFile)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for i := 1; i < 5; i++ {

			uid, _ := uuid.NewUUID()

			XPlaintext := bfv.NewPlaintext(params)
			rX := make([]uint64, 1<<params.LogN)
			rX[0] = uint64(i)
			encoder.EncodeUint(rX, XPlaintext)
			XcipherText := encryptorPk.EncryptNew(XPlaintext)
			XcipherBytes, err := XcipherText.MarshalBinary()
			if err != nil {
				panic(err)
			}
			s := fmt.Sprintf("%v %s", uid.ID(), base64.RawStdEncoding.EncodeToString(XcipherBytes))
			_, err = io.WriteString(file, s+"\n")
			if err != nil {
				panic(err)
			}
		}

	}

	if *mode == "decrypt" {
		var sk bfv.SecretKey
		skBytes, err := ioutil.ReadFile(*skFile)
		if err != nil {
			panic(err)
		}
		sk.UnmarshalBinary(skBytes)

		decryptorSk := bfv.NewDecryptor(params, &sk)

		f, err := os.Open(*encryptedFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		rd := bufio.NewReader(f)
		for {
			line, err := rd.ReadString('\n')
			if err == io.EOF {
				fmt.Print(line)
				break
			}
			if err != nil {
				panic(err)
			}

			l := strings.Split(line, " ")

			encryptedBytes, err := base64.RawStdEncoding.DecodeString(l[1])
			if err != nil {
				panic(err)
			}
			var XcipherT bfv.Ciphertext
			err = XcipherT.UnmarshalBinary(encryptedBytes)
			if err != nil {
				panic(err)
			}
			distPlainT := bfv.NewPlaintext(params)
			decryptorSk.Decrypt(&XcipherT, distPlainT)
			x := encoder.DecodeUint(distPlainT)
			fmt.Printf("%s %v\n", l[0], x[0<<1])

		}

	}
}
