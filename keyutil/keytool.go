package main

import (
	"cryptopasta"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Generating a new key, writing it to private.key")
	key := cryptopasta.NewEncryptionKey()
	ioutil.WriteFile("private.key",key[:], os.ModeAppend | os.ModePerm)
}