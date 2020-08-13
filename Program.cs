using System;
using System.IO;
using Microsoft.Research.SEAL;

namespace HApp
{

    public class Program
    {
        private const int maxCoordinate = 1000;

        private static int randomNumber(int min, int max)
        {
            return _random.Next(min, max);
        }
        private static readonly Random _random = new Random();

        [STAThread]
        static void Main(string[] args)
        {
            int maxDrivers = 20;
            if (args.Length > 0) {
                maxDrivers = Convert.ToInt32(args[0]);
            }
            var c = new Program();

            string publicFile = "/tmp/public.bin";
            string secretFile = "/tmp/secret.bin";
            string encryptedXFile = "/tmp/encryptedXFile.bin";
            string encryptedYFile = "/tmp/encryptedYFile.bin";
            string distanceFile = "/tmp/distance.bin";

            // generate random spot for the rider
            int rx = randomNumber(0, maxCoordinate);
            int ry = randomNumber(0, maxCoordinate);

            // generate public, secret keys and encrypt the x,y location for the rider
            c.generateKeys(rx, ry, publicFile, secretFile, encryptedXFile, encryptedYFile);
            Console.WriteLine("Rider at ({0},{1})",rx,ry);
            Console.WriteLine();
            Console.WriteLine();

            ulong closest = ulong.MaxValue;
            int closestDriverIndex = 0;
            // for each driver
            for (int i = 0; i < maxDrivers; i++) {
                // generate coordinates on the grid
                int dx = randomNumber(0, maxCoordinate);
                int dy = randomNumber(0, maxCoordinate);
                Console.WriteLine("[{0}] Driver at ({1},{2})",i, dx,dy);

                // encrypt the location for this driver using the public key of the rider
                // save the encrypted location to a file as x,y coordinates
                string diversEncryptedXFile = encryptedXFile + "." + Convert.ToString(i);
                string diversEncryptedYFile = encryptedYFile + "." + Convert.ToString(i);
                string diversdistanceFile = distanceFile + "." + Convert.ToString(i);                
                c.encryptDriverLocation(dx, dy, publicFile,  diversEncryptedXFile,  diversEncryptedYFile);

                // Use the encrypted x,y coordinates of the rider and driver to derive the distance; always in encrypted form.
                // Save the distance to a file.
                c.calculateDistance(encryptedXFile, encryptedYFile, diversEncryptedXFile, diversEncryptedYFile, diversdistanceFile);

                // Use the rider secret key to decrypt the distance for a given driver
                var distance = c.decryptDistance(publicFile, secretFile, diversdistanceFile);

                // iterate and find the cloest driver
                Console.WriteLine("   distance from rider: {0}",distance);
                if (distance < closest) {
                    closest = distance;
                    closestDriverIndex = i;
                } 
            }
            Console.WriteLine();
            Console.WriteLine();            
            Console.WriteLine("Closest driver [{0}] at = {1}", closestDriverIndex, closest);
        }

        public void generateKeys(int x, int y, string publicFile, string secretFile, string encryptedXFile, string encryptedYFile)
        {

            //Console.WriteLine("Rider at ({0},{1})",x,y);

            using EncryptionParameters parms = new EncryptionParameters(SchemeType.BFV);
            ulong polyModulusDegree = 4096;
            parms.PolyModulusDegree = polyModulusDegree;
            parms.CoeffModulus = CoeffModulus.BFVDefault(polyModulusDegree);
            parms.PlainModulus = new Modulus(1024);


            using SEALContext context = new SEALContext(parms);
            //Console.WriteLine("Set encryption parameters and print");
            //Console.WriteLine(context);
            //Console.WriteLine("Parameter validation (success): {0}", context.ParameterErrorMessage());

            using KeyGenerator keygen = new KeyGenerator(context);

            using PublicKey riderPub = keygen.PublicKey;
            using SecretKey riderSec = keygen.SecretKey;


            using Evaluator evaluator = new Evaluator(context);
            using IntegerEncoder encoder = new IntegerEncoder(context);

            using Encryptor riderEncryptor = new Encryptor(context, riderPub);
            using Decryptor riderDecryptor = new Decryptor(context, riderSec);

            using Plaintext riderxPlain = encoder.Encode(x);
            using Plaintext rideryPlain = encoder.Encode(y);
            using Ciphertext riderxEncrypted = new Ciphertext();
            using Ciphertext rideryEncrypted = new Ciphertext();
            riderEncryptor.Encrypt(riderxPlain, riderxEncrypted);
            riderEncryptor.Encrypt(rideryPlain, rideryEncrypted);

            var fileStream = File.Create(publicFile);
            riderPub.Save(fileStream);
            fileStream.Close();

            fileStream = File.Create(secretFile);
            riderSec.Save(fileStream);
            fileStream.Close();

            fileStream = File.Create(encryptedXFile);
            riderxEncrypted.Save(fileStream);
            fileStream.Close();

            fileStream = File.Create(encryptedYFile);
            rideryEncrypted.Save(fileStream);
            fileStream.Close();

        }

        public void encryptDriverLocation(int x, int y, string publicFile, string encryptedXFile, string encryptedYFile)
        {

            //Console.WriteLine("Driver at ({0},{1})",x,y);

            using EncryptionParameters parms = new EncryptionParameters(SchemeType.BFV);
            ulong polyModulusDegree = 4096;
            parms.PolyModulusDegree = polyModulusDegree;
            parms.CoeffModulus = CoeffModulus.BFVDefault(polyModulusDegree);
            parms.PlainModulus = new Modulus(1024);

            using SEALContext context = new SEALContext(parms);

            using KeyGenerator keygen = new KeyGenerator(context);

            using PublicKey riderPub = new PublicKey();


            using Ciphertext xEncrypted = new Ciphertext();
            using (var sr = new StreamReader(publicFile))
            {
                riderPub.Load(context, sr.BaseStream);
            }

            using Evaluator evaluator = new Evaluator(context);
            using IntegerEncoder encoder = new IntegerEncoder(context);

            using Encryptor riderEncryptor = new Encryptor(context, riderPub);

            using Plaintext driverxPlain = encoder.Encode(x);
            using Plaintext driveryPlain = encoder.Encode(y);
            using Ciphertext driverxEncrypted = new Ciphertext();
            using Ciphertext driveryEncrypted = new Ciphertext();
            riderEncryptor.Encrypt(driverxPlain, driverxEncrypted);
            riderEncryptor.Encrypt(driveryPlain, driveryEncrypted);

            var fileStream = File.Create(encryptedXFile);
            driverxEncrypted.Save(fileStream);
            fileStream.Close();

            fileStream = File.Create(encryptedYFile);
            driveryEncrypted.Save(fileStream);
            fileStream.Close();

        }

        public void calculateDistance(string encryptedRiderXFile, string encryptedRiderYFile, string encryptedDriverXFile, string encryptedDriverYFile, string distanceFile)
        {

            using EncryptionParameters parms = new EncryptionParameters(SchemeType.BFV);
            ulong polyModulusDegree = 4096;
            parms.PolyModulusDegree = polyModulusDegree;
            parms.CoeffModulus = CoeffModulus.BFVDefault(polyModulusDegree);
            parms.PlainModulus = new Modulus(1024);

            using SEALContext context = new SEALContext(parms);

            using KeyGenerator keygen = new KeyGenerator(context);

            using PublicKey riderPub = new PublicKey();


            using Ciphertext xRiderEncrypted = new Ciphertext();
            using (var sr = new StreamReader(encryptedRiderXFile))
            {
                xRiderEncrypted.Load(context, sr.BaseStream);
            }

            using Ciphertext yRiderEncrypted = new Ciphertext();
            using (var sr = new StreamReader(encryptedRiderYFile))
            {
                yRiderEncrypted.Load(context, sr.BaseStream);
            }

            using Ciphertext xDriverEncrypted = new Ciphertext();
            using (var sr = new StreamReader(encryptedDriverXFile))
            {
                xDriverEncrypted.Load(context, sr.BaseStream);
            }

            using Ciphertext yDriverEncrypted = new Ciphertext();
            using (var sr = new StreamReader(encryptedDriverYFile))
            {
                yDriverEncrypted.Load(context, sr.BaseStream);
            }

            //Console.WriteLine("Calculate ( (x2-x1)^2 + (y2-y1)^2 )");

            using Evaluator evaluator = new Evaluator(context);
            using IntegerEncoder encoder = new IntegerEncoder(context);

            using RelinKeys relinKeys = keygen.RelinKeysLocal();

            using Ciphertext X2MinusX1 = new Ciphertext();
            evaluator.Sub(xDriverEncrypted, xRiderEncrypted, X2MinusX1);          // (x2-x1)

            using Ciphertext Y2MinusY1 = new Ciphertext();
            evaluator.Sub(yDriverEncrypted, yRiderEncrypted, Y2MinusY1);          // (y2-y1)

            using Ciphertext X2MinusX1Squared = new Ciphertext();
            evaluator.Square(X2MinusX1, X2MinusX1Squared);                        // (x2-x1)^2
            //evaluator.RelinearizeInplace(X2MinusX1Squared, relinKeys);

            using Ciphertext Y2MinusY1Squared = new Ciphertext();
            evaluator.Square(Y2MinusY1, Y2MinusY1Squared);                        // (x2-x1)^2
            //evaluator.RelinearizeInplace(Y2MinusY1Squared, relinKeys);

            using Ciphertext X2MinusX1SquaredPlusY2MinusY1Squared = new Ciphertext();
            evaluator.Add(X2MinusX1Squared, Y2MinusY1Squared, X2MinusX1SquaredPlusY2MinusY1Squared);   // (x2-x1)^2 + (y2-y1)^2

            // using Ciphertext SquareRootOfX2MinusX1SquaredPlusY2MinusY1Squared = new Ciphertext(); 
            // decimal vIn = 0.5M;
            // ulong vOut = Convert.ToUInt64(vIn);        
            // evaluator.Exponentiate(X2MinusX1SquaredPlusY2MinusY1Squared, vOut,relinKeys,SquareRootOfX2MinusX1SquaredPlusY2MinusY1Squared);

            var fileStream = File.Create(distanceFile);
            X2MinusX1SquaredPlusY2MinusY1Squared.Save(fileStream);
            fileStream.Close();

        }

        public ulong decryptDistance(string publicFile, string secretFile, string distanceFile)
        {

            using EncryptionParameters parms = new EncryptionParameters(SchemeType.BFV);
            ulong polyModulusDegree = 4096;
            parms.PolyModulusDegree = polyModulusDegree;
            parms.CoeffModulus = CoeffModulus.BFVDefault(polyModulusDegree);
            parms.PlainModulus = new Modulus(1024);

            using SEALContext context = new SEALContext(parms);

            using KeyGenerator keygen = new KeyGenerator(context);

            using PublicKey riderPub = new PublicKey();
            using (var sr = new StreamReader(publicFile))
            {
                riderPub.Load(context, sr.BaseStream);
            }

            using SecretKey riderSec = new SecretKey();
            using (var sr = new StreamReader(secretFile))
            {
                riderSec.Load(context, sr.BaseStream);
            }

            using Ciphertext X2MinusX1SquaredPlusY2MinusY1Squared = new Ciphertext();
            using (var sr = new StreamReader(distanceFile))
            {
                X2MinusX1SquaredPlusY2MinusY1Squared.Load(context, sr.BaseStream);
            }

            using Encryptor riderEncryptor = new Encryptor(context, riderPub);
            using Decryptor riderDecryptor = new Decryptor(context, riderSec);


            using Plaintext distance = new Plaintext();
            riderDecryptor.Decrypt(X2MinusX1SquaredPlusY2MinusY1Squared, distance);

            using IntegerEncoder encoder = new IntegerEncoder(context);

            //Console.WriteLine("Decrypted polynomial f(x) = {0}.", distance );
            //Console.WriteLine("     distance from rider {0}", encoder.DecodeUInt64(distance));
            return encoder.DecodeUInt64(distance); 

        }
    }
}