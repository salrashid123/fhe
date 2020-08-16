#!/bin/bash



rm -rf /tmp/rider /tmp/drivers /tmp/distance
mkdir -p /tmp/rider
mkdir -p /tmp/drivers
mkdir -p /tmp/distance

randX=`shuf -i 0-1000 -n 1`
randY=`shuf -i 0-1000 -n 1`

## first just test encryption/decryption
go run main.go -mode=genkey --pkFile=/tmp/rider/pk.bin --skFile=/tmp/rider/sk.bin
go run main.go -mode=encrypt --pkFile=/tmp/rider/pk.bin --x=$randX --y=$randY --encryptedFile=/tmp/rider/rider.bin
go run main.go -mode=decrypt --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/rider/rider.bin

for i in {1..10}
do
   randX=`shuf -i 0-1000 -n 1`
   randY=`shuf -i 0-1000 -n 1`
   echo "Driver [$i] at ($randX, $randY)"
   go run main.go -mode=encrypt --pkFile=/tmp/rider/pk.bin --x=$randX --y=$randY --encryptedFile=/tmp/drivers/$i.bin
   go run main.go -mode=distance --riderEncryptedFile=/tmp/rider/rider.bin --driverEncryptedFile=/tmp/drivers/$i.bin --encryptedFile=/tmp/distance/$i.bin
   go run main.go -mode=decryptdistance --skFile=/tmp/rider/sk.bin --encryptedFile=/tmp/distance/$i.bin
done