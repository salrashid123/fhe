## Simple RideSharing "helloworld" with Homomorphic Encryption

Just a simple C# application that calculates the closest driver to a rider using Homomorphic Encryption.

This sample uses `C#` Microsoft SEAL Library to perform the distance calculation under encryption:

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

### References

- [A Homomorphic Encryption Illustrated Primer](https://blog.n1analytics.com/homomorphic-encryption-illustrated-primer/)
- [Microsoft SEAL](https://github.com/microsoft/SEAL)
- [SEAL Manual](https://www.microsoft.com/en-us/research/uploads/prod/2017/11/sealmanual-2-3-1.pdf)
- [SEAL Demo](https://github.com/microsoft/SEAL-Demo)
- [Lttigo](https://github.com/ldsec/lattigo)

