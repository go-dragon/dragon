package dragoncrypto

//just use AesDecrypt and AesDecrypt func to encode and decode, it'll return a base64 encode string
import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

//aes  link https://www.jianshu.com/p/b63095c59361
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData, key []byte) (string, error) {
	block, err := aes.NewCipher(key) //key is 16,24,32 bit, and origData's size is multiple of key, we use pkcs7 to fill up the size
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	result := base64.StdEncoding.EncodeToString(crypted) //use base64 encode
	return result, nil
}

func AesDecrypt(encryptData string, key []byte) ([]byte, error) {
	crypted, err := base64.StdEncoding.DecodeString(encryptData)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
