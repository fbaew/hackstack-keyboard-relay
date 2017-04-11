package cryptoencoder

import (
    "cryptopasta"
    "fmt"
    "io/ioutil"
    "log"
    "encoding/hex"
    "errors"
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

func Encode(plainText string, key *[32]byte) (string, error) {
    cipherText, encryptionError := cryptopasta.Encrypt([]byte(plainText),key)
    if encryptionError != nil {
        return "", errors.New("Failed to encrypt data")
    }
    fmt.Println("Encrypted successfully!")

    hexCipherText := hex.EncodeToString(cipherText)

    fmt.Printf("Encrypted data: %s\n",hexCipherText)
    return hexCipherText + "\n",nil

}
func Decode(cipherText []byte, key *[32]byte) (string, error) {
    hexBytes := make([]byte, hex.DecodedLen(len(cipherText)-1))
    n, decodingError := hex.Decode(hexBytes, cipherText[:len(cipherText)-1])
    if decodingError != nil {
        fmt.Println("Unable to decode received string")
        return "", errors.New("Unable to decode received string")
    }
    fmt.Printf("Decoded string successfully...  (%d bytes)\n", n)
    fmt.Println("Decrypting string...")

    plainText, decryptionError := cryptopasta.Decrypt(hexBytes,key)
    if decryptionError != nil {
        fmt.Println("Unable to decrypt decoded value")
        return "", errors.New("Unable to decrypt decoded value")
    }
    return string(plainText), nil
}

func main() {
//    generateKeyFile()

    sampleString := "Hello, world!"
    encodedHex, encodingError := Encode(sampleString, LoadKey("private.key"))
    if encodingError != nil {}
    fmt.Printf("Encoded data: %s\n", encodedHex)
    fmt.Println("")
    decodedText, decodingError := Decode([]byte(encodedHex), LoadKey("private.key"))
    if decodingError != nil {}
    fmt.Printf("Decoded data: %s\n",decodedText)


}
