package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"saltextio"

	"github.com/apache/beam/sdks/go/pkg/beam"

	// "github.com/apache/beam/sdks/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/go/pkg/beam/x/beamx"
	"github.com/ldsec/lattigo/bfv"
)

var (
	input  = flag.String("input", "/tmp/encrypted.bin", "File(s) to read.")
	pkFile = flag.String("pkFile", "/tmp/rider/pk.bin ", "publicKey File read.")
	output = flag.String("output", "", "Output file (required).")

	params      *bfv.Parameters
	pk          bfv.PublicKey
	encoder     bfv.Encoder
	encryptorPk bfv.Encryptor
	evaluator   bfv.Evaluator
)

var (
	operandAdd uint64 = 2
)

func extractFn(ctx context.Context, line string, emit func(string, string)) {
	kv := strings.Split(line, " ")
	if len(kv) == 0 {
		return
	}
	x := bfv.NewCiphertext(params, 1)
	encBytes, err := base64.RawStdEncoding.DecodeString(kv[1])
	if err != nil {
		panic(err)
	}

	x.UnmarshalBinary(encBytes)

	YPlaintext := bfv.NewPlaintext(params)
	rY := make([]uint64, 1<<params.LogN)
	rY[0] = operandAdd
	encoder.EncodeUint(rY, YPlaintext)
	y := encryptorPk.EncryptNew(YPlaintext)

	xPlusY := bfv.NewCiphertext(params, 1)
	evaluator.Add(x, y, xPlusY)

	xPlusYBytes, err := xPlusY.MarshalBinary()
	if err != nil {
		panic(err)
	}

	emit(kv[0], base64.RawStdEncoding.EncodeToString(xPlusYBytes))
}

func formatFn(w string, c string) string {
	return fmt.Sprintf("%s %s", w, c)
}

func Compute(s beam.Scope, lines beam.PCollection) beam.PCollection {
	s = s.Scope("Compute")
	return beam.ParDo(s, extractFn, lines)
}

func main() {
	flag.Parse()
	beam.Init()

	if *output == "" {
		log.Fatal("No output provided")
	}

	if *pkFile == "" {
		log.Fatal("No pkFile provided")
	}

	// ctx := context.Background()
	// c, err := gcsx.NewClient(ctx, storage.ScopeReadOnly)
	// if err != nil {
	// 	log.Fatal("Could not create GCS context")
	// }

	// buckets, object, err := gcsx.ParseObject(*pkFile) //  gs://mineral-minutia-820-rider-1/pk.bin
	// if err != nil {
	// 	log.Fatal("No pkFile provided")
	// }

	// ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	// defer cancel()

	// pkBytes, err := gcsx.ReadObject(ctx, c, buckets, object)
	// if err != nil {
	// 	log.Fatal("Could not read pkFile provided")
	// }
	// pk.UnmarshalBinary(pkBytes)

	// BFV parameters (128 bit security)
	params = bfv.DefaultParams[bfv.PN12QP109]
	encoder = bfv.NewEncoder(params)
	evaluator = bfv.NewEvaluator(params)
	pkBytes, err := ioutil.ReadFile(*pkFile)
	if err != nil {
		panic(err)
	}

	pk.UnmarshalBinary(pkBytes)

	encryptorPk = bfv.NewEncryptorFromPk(params, &pk)

	p := beam.NewPipeline()
	s := p.Root()

	lines := saltextio.Read(s, *input)

	counted := Compute(s, lines)
	formatted := beam.ParDo(s, formatFn, counted)
	saltextio.Write(s, *output, formatted)

	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}
