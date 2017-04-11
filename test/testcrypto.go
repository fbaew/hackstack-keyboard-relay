package main

import (
    "cryptopasta"
    "fmt"
    "encoding/hex"
    "io/ioutil"
    "os"
)

func main() {
    key := cryptopasta.NewEncryptionKey()
    var keyslice []byte = key[0:32]
    plaintext := []byte("Hello, secret world!")

    ciphertext, err := cryptopasta.Encrypt(plaintext, key)
    if err != nil { fmt.Println("There was a problem encrypting the data") }

    fmt.Println("Encryption Key (hex):")
    fmt.Println(hex.EncodeToString(keyslice))


    fmt.Println("Plaintext:")
    fmt.Println(hex.EncodeToString(plaintext))
    fmt.Println("")
    fmt.Println("Ciphertext:")
    fmt.Println(hex.EncodeToString(ciphertext))
    fmt.Println("")
    fmt.Println("Plaintext (decrypted):")
    
    decrypted, err := cryptopasta.Decrypt(ciphertext, key)
    if err != nil { fmt.Println("Unable to decrypt data") }
    fmt.Println(hex.EncodeToString(decrypted))

    writeToFile(key)
}

func writeToFile(key *[32]byte) {
    ioutil.WriteFile("key.key",[]byte(hex.EncodeToString(key[0:32]) + "\n"),os.ModeAppend|os.ModePerm)
}

