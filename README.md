## Simple RideSharing "helloworld" with Homomorphic Encryption

Just a simple C# application that calculates the closest driver to a rider using Homomorphic Encryption.

This sample uses both `C#` Microsoft SEAL Library and Golang `lattigo` to perform the distance calculation under encryption:

1. On startup a Rider is assigned a random `(x1, y1)` coordinate in a `1000x1000` grid
2. Rider generates homomorphic publickey, secretkey:  `pk`, `sk`
3. Rider encrypts the xy coordinates using `pk`:  `x1`->`x1'`, `y1`->`y1'`
4. Rider sends encrypted `x1'`, `y1'`  and `pk` to server
5. Server sends `pk` to 50 drivers each in a in random position in the grid
6. Each driver encrypts their respective location `x2`,`y2` using `pk`  -> `x2'`,`y2'`
7. Each driver each sends `x2'`, `y2'` to the server
8. Server calculates over the still encrypted data the distance from the rider to each driver

   `(x2'-x1')^2  + (y2'-y1')^2`  

9. Server sends the encrypted result back to the rider
10. Rider decrypts distance from each driver using the sk  
11. Rider picks the closest driver

The thing to note is the calculation done in step 8 is done using _encrypted data_....the server never decrypts any of the data its provided but nonetheless, it can calculate the distance value.

This repository also has a simple Apache Beam piepeline runner which will homomorphically add a given number to an already encrypted input.  You can extend that sample to make further modifications and create a simple streaming ridesharing app.

### C#

### Build and run

Use docker to build SEAL, compile the sample application and run:

The app will iterate over 50 random points, encrypt them and allow the Rider to decrypt the distance value:

```bash
docker build -t htest .

docker run -t htest

Rider at (281,399)

Driver at (137,340)
  distance from rider: 24217
Driver at (949,705)
  distance from rider: 539860
Driver at (730,973)
  distance from rider: 531077
Driver at (41,178)
  distance from rider: 106441
Driver at (212,203)
  distance from rider: 43177
Driver at (429,30)
  distance from rider: 158065
Driver at (446,821)
  distance from rider: 205309
Driver at (690,224)
  distance from rider: 197906
Driver at (488,105)
  distance from rider: 129285
Driver at (832,112)
  distance from rider: 385970
Driver at (199,64)
  distance from rider: 118949
Driver at (347,400)
  distance from rider: 4357        <<<<<<<<<<<<<
Driver at (804,181)
  distance from rider: 321053
Driver at (1,123)
  distance from rider: 154576
Driver at (520,514)
  distance from rider: 70346
Driver at (346,862)
  distance from rider: 218594
Driver at (962,107)
  distance from rider: 549025
Driver at (23,634)
  distance from rider: 121789
Driver at (808,673)
  distance from rider: 352805
Driver at (389,597)
  distance from rider: 50868
Cloest = 4357

```

## Golang

The go program by itself does not run in a loop but rather individually processes each step as a command line.  It's trivial to change this to iterate all 50 within a loop automatically as in the C# sample but i just left it as a "standalone" command set.

### Standalone

The following runs each step as a standalone command set which is later wrapped into the `runner.sh` and iterate over 50 drivers:
```bash
## Create rider public,secret key
go run main.go -mode=genkey --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin

## Then test encryption/decryption for a rider at (x,y)=(2,2)
go run main.go -mode=encrypt --pkFile=/tmp/rider/pk.bin --x=2 --y=2 --encryptedFile=/tmp/rider/rider.bin
go run main.go -mode=decrypt --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/rider/rider.bin

## Use the riders public key to encrypt a drivers location at (4,5)
go run main.go -mode=encrypt --pkFile=/tmp/rider/pk.bin --x=4 --y=5 --encryptedFile=/tmp/drivers/1.bin

## Derive the distance between the rider and driver given their _encrypted_ values
go run main.go -mode=distance --riderEncryptedFile=/tmp/rider/rider.bin --driverEncryptedFile=/tmp/drivers/1.bin --encryptedFile=/tmp/distance/1.bin

## Rider decrypts each drivers' location
go run main.go -mode=decryptdistance --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/distance/1.bin
```

To run the above commands as a schell script,
```bash
 ./runner.sh 
### RIder location and ID
2020/08/16 17:17:52 Decrypted: id: f1da6487-e005-11ea-baf5-e86a641d5560, (323,95)

## Driver location and ID
Driver [1] at (821, 145)
### Decrypted distance between driver and rider
  2020/08/16 17:17:54 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f26cc4e7-e005-11ea-bc51-e86a641d5560]  (53893)
Driver [2] at (679, 726)
  2020/08/16 17:17:56 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f35334df-e005-11ea-b9be-e86a641d5560]  (601)  <<<<<
Driver [3] at (412, 800)
  2020/08/16 17:17:57 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f4517dc8-e005-11ea-add5-e86a641d5560]  (46187)
Driver [4] at (180, 557)
  2020/08/16 17:17:59 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f53bb22b-e005-11ea-8e35-e86a641d5560]  (37282)
Driver [5] at (438, 703)
  2020/08/16 17:18:00 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f629c73a-e005-11ea-8e46-e86a641d5560]  (55204)
Driver [6] at (670, 888)
  2020/08/16 17:18:02 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f711526b-e005-11ea-b47f-e86a641d5560]  (28351)
Driver [7] at (724, 574)
  2020/08/16 17:18:03 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f7e10d6c-e005-11ea-a901-e86a641d5560]  (62557)
Driver [8] at (827, 107)
  2020/08/16 17:18:05 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f8bde225-e005-11ea-b0b7-e86a641d5560]  (57549)
Driver [9] at (718, 975)
  2020/08/16 17:18:06 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [f99c6359-e005-11ea-b50c-e86a641d5560]  (12907)
Driver [10] at (886, 514)
  2020/08/16 17:18:08 Distance: from Rider [f1da6487-e005-11ea-baf5-e86a641d5560] --> Driver [faa55930-e005-11ea-8d91-e86a641d5560]  (33771)
```


### References

- [A Homomorphic Encryption Illustrated Primer](https://blog.n1analytics.com/homomorphic-encryption-illustrated-primer/)
- [Microsoft SEAL](https://github.com/microsoft/SEAL)
- [SEAL Manual](https://www.microsoft.com/en-us/research/uploads/prod/2017/11/sealmanual-2-3-1.pdf)
- [SEAL Demo](https://github.com/microsoft/SEAL-Demo)
- [Lttigo](https://github.com/ldsec/lattigo)

