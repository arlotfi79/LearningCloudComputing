package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

func createRandomString(n int) []byte {
	var baseString = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, n)
	for i := range b {
		b[i] = baseString[rand.Intn(len(baseString))]
	}
	return b
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

// @returns: checksum of the file
func createRandomFile(path string) {
	info := createRandomString(1024)
	err := ioutil.WriteFile(path, info, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
