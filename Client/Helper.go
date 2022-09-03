package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func createFileInClient(path string, info string) {
	err := ioutil.WriteFile(path, []byte(info), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func calculateCheckSum(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
