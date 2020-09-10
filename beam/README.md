## Simple math with Beam DirectRunner and Homomorphic Encryption

Sample Beam pipeline that will read a file in which each line represents a homomorphically encrypted number.


Beam will read each line in the file, use the same public key to add `2` to that encrypted value number and then save the output into a target file line by line.

>> At no time does Beam actually "know" what number is encoded within the source file....it just runs arithmetic against it blindly.

First Generate the encrypted values.

The `/client/` application will generate a public/secret key pair combo, then use the pubic key to encrypt sequential number `[1,2,3,4]` numbers and save those numbers and an arbitrary `uid` filed into an output text file `/tmp/encrypted.bin`. 

Once encrypted, you can run Beam which reads the encrypted file and public key, then for each line adds "2" to that value homomorphically, then saves each line into an output file

Finally, use the secret key to decode the encrypted file that added that number to it `/processed.bin`
The output will show  `[3,4,5,6]` 


- Encrypt

First encrypt a sequence
```bash
cd client/
go run main.go  -mode=encrypt --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/encrypted.bin
```

`/tmp/encryted.bin` is a file where the first column is just q unique ID and the second is the encrypted form of a number

```
1049460993 Encrypted(1)
1049499542 Encrypted(2)
1049539794 Encrypted(3)
1049578153 Encrypted(4) 
```

- Process

Run a dataflow localRunner that adds `2` to each encrypted value

```bash
cd ../
go run main.go --input=/tmp/encrypted.bin  --output=/tmp/processed.bin  --pkFile=/tmp/rider/pk.bin 
```


```golang
	encoder.EncodeUint(rY, YPlaintext)
	y := encryptorPk.EncryptNew(YPlaintext)

	xPlusY := bfv.NewCiphertext(params, 1)
	evaluator.Add(x, y, xPlusY)
```

- Decrypt

Now decrypt the output fule `/tmp/processed.bin` using the `sk.bin`

```bash
cd client/
go run main.go  -mode=decrypt --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/processed.bin
```

The output will be the number that was originally encrypted but larger by 2
```
1049460993 3
1049499542 4
1049539794 5
1049578153 6

```

