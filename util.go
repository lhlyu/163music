package music163

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
)

func doGet(apiUrl string) (string, error) {
	resp, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func doPost(apiUrl, referer string, body io.Reader) (*gjson.Result, error) {
	req, err := http.NewRequest(http.MethodPost, apiUrl, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Referer", referer)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := gjson.ParseBytes(b)
	return &r, nil
}

func aesEncrypt(encStr string, key string) string {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "err"
	}
	src := _PKCS7Padding([]byte(encStr), block.BlockSize())
	blockModel := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(src))
	blockModel.CryptBlocks(cipherText, src)
	return base64.StdEncoding.EncodeToString(cipherText)
}

func _PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
