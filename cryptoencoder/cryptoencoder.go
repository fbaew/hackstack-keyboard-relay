package cryptoencoder

import (
    "cryptopasta"
    "fmt"
    "io/ioutil"
    "log"
    "encoding/hex"
)

/*
func generateKeyFile() {
    ioutil.WriteFile("private.key", cryptopasta.NewEncryptionKey()[:], os.ModeAppend | os.ModePerm)
}
*/

func LoadKey(keyfile string) *[32]byte {
    fmt.Printf("Loading %s\n", keyfile)
    keyData, keyLoadError := ioutil.ReadFile(keyfile)
    if keyLoadError != nil { log.Fatal(keyLoadError) }
    if len(keyData) != 32 {
        fmt.Printf("Read %d bytes...\n",len(keyData))
        log.Fatal("Key file is unexpected size")
    }

    fmt.Printf("Success... Read %d bytes...\n",len(keyData))
    fmt.Printf("Key: %s\n",hex.EncodeToString(keyData))

    var key [32]byte
    copy(key[:],keyData[:])

    return &key
}

func Encode(plainText string, key *[32]byte) string {
    cipherText, encryptionError := cryptopasta.Encrypt([]byte(plainText),key)
    if encryptionError != nil { log.Fatal("Failed to encrypt text") }
    fmt.Println("Encrypted successfully!")

    hexCipherText := hex.EncodeToString(cipherText)

    fmt.Printf("Encrypted data: %s\n",hexCipherText)
    return hexCipherText + "\n"

}

func Decode(cipherText []byte, key *[32]byte) string {
    hexBytes := make([]byte, hex.DecodedLen(len(cipherText)-1))
    n, decodingError := hex.Decode(hexBytes, cipherText[:len(cipherText)-1])
    if decodingError != nil { log.Fatal(decodingError) }
    fmt.Printf("Decoded string successfully...  (%d bytes)\n", n)
    fmt.Println("Decrypting string...")
    plainText, decryptionError := cryptopasta.Decrypt(hexBytes,key)
    if decryptionError != nil { log.Fatal(decryptionError) }
    return string(plainText)
}

func main() {
//    generateKeyFile()

    sampleString := "Hello, world!"
    encodedHex := Encode(sampleString, LoadKey("private.key"))
    fmt.Printf("Encoded data: %s\n", encodedHex)
    fmt.Println("")
    decodedText := Decode([]byte(encodedHex), LoadKey("private.key"))
    fmt.Printf("Decoded data: %s\n",decodedText)

}
