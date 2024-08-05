package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"joseluis244/symphonycloudtools"
	"os"
)

func main() {
	fingerprint, err := symphonycloudtools.HashGenerator()
	if err != nil {
		panic(err)
	}
	keys, err := os.ReadFile("test/licensegen/keys.json")
	if err != nil {
		panic(err)
	}
	var M map[string]map[string]interface{}
	json.Unmarshal(keys, &M)
	R2map := M["R2"]
	jr2, err := json.Marshal(R2map)
	if err != nil {
		panic(err)
	}
	Firebasemap := M["Firebase"]
	jfirebase, err := json.Marshal(Firebasemap)
	if err != nil {
		panic(err)
	}
	combinedMaps := ""
	combinedMaps = fmt.Sprintf("%s||%v||%v", fingerprint, string(jr2), string(jfirebase))
	encrypted, err := encryptString(combinedMaps, "MedicareSoft203$")
	if err != nil {
		panic(err)
	}
	fmt.Println("license...", encrypted, len(encrypted))
}

func encryptString(plaintext string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return hex.EncodeToString(ciphertext), nil
}
