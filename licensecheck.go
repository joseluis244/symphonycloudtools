package symphonycloudtools

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

func decryptLicense(ciphertextHex string, key string) (fingerprint string, r2 map[string]interface{}, firebase map[string]interface{}, err error) {
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return "", nil, nil, err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", nil, nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", nil, nil, fmt.Errorf("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	licenses := strings.Split(string(ciphertext), "||")
	Fingerprint := licenses[0]
	var R2 map[string]interface{}
	json.Unmarshal([]byte(licenses[1]), &R2)
	var Firebase map[string]interface{}
	json.Unmarshal([]byte(licenses[2]), &Firebase)
	return Fingerprint, R2, Firebase, nil
}
