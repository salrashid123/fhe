package main

import (
	"io/ioutil"
	"log"

	"flag"

	"github.com/google/uuid"
	"github.com/ldsec/lattigo/bfv"
	"github.com/salrashid123/fhe/rideshare"
	"google.golang.org/protobuf/proto"
)

func genKey(pubFile string, secFile string) error {

	// BFV parameters (128 bit security)
	params := bfv.DefaultParams[bfv.PN12QP109]
	//params.T = //0x3ee0001 //65929217₁₀

	kgen := bfv.NewKeyGenerator(params)
	riderSk, riderPk := kgen.GenKeyPair()

	pubBytes, err := riderPk.MarshalBinary()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(pubFile, pubBytes, 0640)
	if err != nil {
		return err
	}

	secBytes, err := riderSk.MarshalBinary()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(secFile, secBytes, 0640)
	if err != nil {
		return err
	}

	return nil
}

func encryptwithPub(id string, x uint64, y uint64, pubFile string, encryptedFile string) error {

	// BFV parameters (128 bit security)
	params := bfv.DefaultParams[bfv.PN12QP109]
	//params.T =  0x3ee0001 //65929217₁₀

	encoder := bfv.NewEncoder(params)

	var pk bfv.PublicKey

	pkBytes, err := ioutil.ReadFile(pubFile)
	if err != nil {
		return err
	}

	pk.UnmarshalBinary(pkBytes)

	encryptorPk := bfv.NewEncryptorFromPk(params, &pk)

	XPlaintext := bfv.NewPlaintext(params)
	rX := make([]uint64, 1<<params.LogN)
	rX[0] = x
	encoder.EncodeUint(rX, XPlaintext)
	XcipherText := encryptorPk.EncryptNew(XPlaintext)
	XcipherBytes, err := XcipherText.MarshalBinary()
	if err != nil {
		return err
	}

	YPlaintext := bfv.NewPlaintext(params)
	rY := make([]uint64, 1<<params.LogN)
	rY[0] = y
	encoder.EncodeUint(rY, YPlaintext)
	YcipherText := encryptorPk.EncryptNew(YPlaintext)
	YcipherBytes, err := YcipherText.MarshalBinary()
	if err != nil {
		return err
	}

	eLoc := &rideshare.EncryptedCoordinate{
		Id: id,
		X:  XcipherBytes,
		Y:  YcipherBytes,
		Pk: pkBytes,
	}

	cipherBytes, err := proto.Marshal(eLoc)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(encryptedFile, cipherBytes, 0640)
	if err != nil {
		return err
	}

	return nil
}

func decryptwithSec(secFile string, encryptedFile string) (rideshare.DecryptedCoordinate, error) {

	// BFV parameters (128 bit security)
	params := bfv.DefaultParams[bfv.PN12QP109]
	//params.T =  0x3ee0001 //65929217₁₀

	encoder := bfv.NewEncoder(params)

	var sk bfv.SecretKey
	skBytes, err := ioutil.ReadFile(secFile)
	if err != nil {
		return rideshare.DecryptedCoordinate{}, err
	}

	sk.UnmarshalBinary(skBytes)

	// *****

	encBytes, err := ioutil.ReadFile(encryptedFile)
	if err != nil {
		return rideshare.DecryptedCoordinate{}, err
	}

	decryptorSk := bfv.NewDecryptor(params, &sk)

	serialized := &rideshare.EncryptedCoordinate{}
	err = proto.Unmarshal(encBytes, serialized)
	if err != nil {
		return rideshare.DecryptedCoordinate{}, err
	}

	id := serialized.Id

	var XcipherT bfv.Ciphertext
	err = XcipherT.UnmarshalBinary(serialized.X)
	if err != nil {
		return rideshare.DecryptedCoordinate{}, err
	}
	XplainT := bfv.NewPlaintext(params)
	decryptorSk.Decrypt(&XcipherT, XplainT)
	x := encoder.DecodeUint(XplainT)

	var YcipherT bfv.Ciphertext
	err = YcipherT.UnmarshalBinary(serialized.Y)
	if err != nil {
		return rideshare.DecryptedCoordinate{}, err
	}
	YplainT := bfv.NewPlaintext(params)
	decryptorSk.Decrypt(&YcipherT, YplainT)
	y := encoder.DecodeUint(YplainT)

	ret := rideshare.DecryptedCoordinate{
		Id: id,
		X:  x[0<<1],
		Y:  y[0<<1],
	}
	return ret, nil
}

func distance(riderEncryptedFile string, driverEncryptedFile string, encryptedFile string) error {

	// BFV parameters (128 bit security)
	params := bfv.DefaultParams[bfv.PN12QP109]
	//params.T =  0x3ee0001 //65929217₁₀

	evaluator := bfv.NewEvaluator(params)
	//encoder := bfv.NewEncoder(params)

	rLoc := &rideshare.EncryptedCoordinate{}

	protoBytes, err := ioutil.ReadFile(riderEncryptedFile)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(protoBytes, rLoc)
	if err != nil {
		return err
	}

	rX := &bfv.Ciphertext{}
	rY := &bfv.Ciphertext{}

	err = rX.UnmarshalBinary(rLoc.X)
	if err != nil {
		return err
	}

	err = rY.UnmarshalBinary(rLoc.Y)
	if err != nil {
		return err
	}

	dX := &bfv.Ciphertext{}
	dY := &bfv.Ciphertext{}

	dLoc := &rideshare.EncryptedCoordinate{}

	protoBytes, err = ioutil.ReadFile(driverEncryptedFile)
	if err != nil {
		return err
	}

	err = proto.Unmarshal(protoBytes, dLoc)
	if err != nil {
		return err
	}

	err = dX.UnmarshalBinary(dLoc.X)
	if err != nil {
		return err
	}

	err = dY.UnmarshalBinary(dLoc.Y)
	if err != nil {
		return err
	}

	pK := &bfv.PublicKey{}

	err = pK.UnmarshalBinary(rLoc.Pk)
	if err != nil {
		return err
	}

	//encryptorPk := bfv.NewEncryptorFromPk(params, &pK)

	X2MinusX1 := bfv.NewCiphertext(params, 1)
	Y2MinusY1 := bfv.NewCiphertext(params, 1)
	X2MinusX1Squared := bfv.NewCiphertext(params, 1)
	Y2MinusY1Squared := bfv.NewCiphertext(params, 1)
	distEnc := bfv.NewCiphertext(params, 1)

	evaluator.Sub(dX, rX, X2MinusX1)
	evaluator.Sub(dY, rY, Y2MinusY1)

	X2MinusX1Squared = evaluator.MulNew(X2MinusX1, X2MinusX1)
	Y2MinusY1Squared = evaluator.MulNew(Y2MinusY1, Y2MinusY1)

	distEnc = evaluator.AddNew(X2MinusX1Squared, Y2MinusY1Squared)

	distBytes, err := distEnc.MarshalBinary()
	if err != nil {
		return err
	}
	dst := &rideshare.Distance{
		Rid:  rLoc.Id,
		Did:  dLoc.Id,
		Dist: distBytes,
	}

	distanceBytes, err := proto.Marshal(dst)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(encryptedFile, distanceBytes, 0640)
	if err != nil {
		return err
	}

	return nil
}

func decryptDistanceWithSec(secFile string, encryptedFile string) (rid string, did string, d uint64, err error) {

	// BFV parameters (128 bit security)
	params := bfv.DefaultParams[bfv.PN12QP109]
	//params.T =  0x3ee0001 //65929217₁₀

	encoder := bfv.NewEncoder(params)

	var sk bfv.SecretKey
	skBytes, err := ioutil.ReadFile(secFile)
	if err != nil {
		return "", "", 0, err
	}
	sk.UnmarshalBinary(skBytes)
	decryptorSk := bfv.NewDecryptor(params, &sk)

	encBytes, err := ioutil.ReadFile(encryptedFile)
	if err != nil {
		return "", "", 0, err
	}

	serialized := &rideshare.Distance{}
	err = proto.Unmarshal(encBytes, serialized)
	if err != nil {

	}

	var XcipherT bfv.Ciphertext
	err = XcipherT.UnmarshalBinary(serialized.Dist)
	if err != nil {
		return "", "", 0, err
	}
	distPlainT := bfv.NewPlaintext(params)
	decryptorSk.Decrypt(&XcipherT, distPlainT)
	x := encoder.DecodeUint(distPlainT)
	return serialized.Rid, serialized.Did, x[0<<1], nil
}

func main() {
	mode := flag.String("mode", "", "(required) genkey, encrypt, decrypt, calc")
	pkFile := flag.String("pkFile", "", "(required for mode=genkey) Filename for the rider's publicKey")
	skFile := flag.String("skFile", "", "(required for mode=genkey) Filename for the rider's secrettKey")
	id := flag.String("id", "", "(required for mode=encrypt)")
	x := flag.Uint64("x", 0, "(required for mode=encrypt)")
	y := flag.Uint64("y", 0, "(required for mode=encrypt)")
	encryptedFile := flag.String("encryptedFile", "", "(required for mode=encrypt) Filename for to save the encryptedFile")
	riderEncryptedFile := flag.String("riderEncryptedFile", "", "(required for mode=distance) Filename for to save the riderEncryptedFile")
	driverEncryptedFile := flag.String("driverEncryptedFile", "", "(required for mode=distance) Filename for to save the encryptedFile")

	flag.Parse()
	if *mode == "" {
		log.Fatalf("mode must be set")
	}
	if *mode == "genkey" {
		if *pkFile == "" || *skFile == "" {
			log.Fatalf("--pkFile and --skFile must be set for --mode=riderencrypt")
		}
		err := genKey(*pkFile, *skFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}
	if *mode == "encrypt" {
		if *pkFile == "" || *encryptedFile == "" {
			log.Fatalf("--pkFile, x,y, encryptedFile must be set for --mode=encrypt")
		}
		if *id == "" {
			uid, _ := uuid.NewUUID()
			*id = uid.String()
		}
		err := encryptwithPub(*id, *x, *y, *pkFile, *encryptedFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}
	if *mode == "distance" {
		if *riderEncryptedFile == "" || *driverEncryptedFile == "" || *encryptedFile == "" {
			log.Fatalf("--riderEncryptedFile, driverEncryptedFile, encryptedFile must be set for --mode=distance")
		}
		err := distance(*riderEncryptedFile, *driverEncryptedFile, *encryptedFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
	}

	if *mode == "decrypt" {
		if *skFile == "" || *encryptedFile == "" {
			log.Fatalf("--skFile, encryptedFile must be set for --mode=decrpt")
		}
		r, err := decryptwithSec(*skFile, *encryptedFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		log.Printf("Decrypted: id: %s, (%d,%d)\n", r.Id, r.X, r.Y)
	}

	if *mode == "decryptdistance" {
		if *skFile == "" || *encryptedFile == "" {
			log.Fatalf("--skFile, encryptedFile must be set for --mode=decrpt")
		}
		rid, did, r, err := decryptDistanceWithSec(*skFile, *encryptedFile)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		log.Printf("Distance: from Rider [%s] --> Driver [%s]  (%d)\n", rid, did, r)
	}

}
