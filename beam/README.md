## Simple math with Beam DirectRunner and Homomorphic Encryption

Sample Beam pipeline that will read a file in which each line represents a homomorphically encrypted number.


Beam will read each file, use the same public key to add "2" to that number and then save the output into a target file.

At no time does Beam actually "know" what number is encoded within the source file....it just runs arithmetic against it blindly.

First Generate the encrypted values.

The `/client/` application will generate a public/secret key pair combo, then use the pubic key to encrypt sequention number `[1,2,3,4,5]` numbers and save those numbers and an arbitrary `uid` filed into an output text file `/tmp/encrypted.bin`. 

Once encrypted, you can run Beam which reads the encrypted file and public key, then for each line adds "2" to that value homomorphically, then saves each line into an output file

Finally, use the secret key to decode the encrypted file that added that number to it `/processed.bin`


```bash
cd client/
go run main.go  -mode=encrypt --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/encrypted.bin
```

```bash
cd ../
go run main.go --input=/tmp/encrypted.bin  --output=/tmp/processed.bin  --pkFile=/tmp/rider/pk.bin 
```

```bash
cd client/
go run main.go  -mode=decrypt --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/processed.bin
```

