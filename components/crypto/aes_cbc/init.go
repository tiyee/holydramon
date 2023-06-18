package aes_cbc

import "github.com/tiyee/holydramon/components/crypto"

const _iv = "abc"
const _key = "def"

var iv = _iv
var key = _key
var cryptor crypto.ICrypto

func Reset(newKey, newIv string) {
	cryptor = NewAesCbc(newKey, newIv)
}
func Decrypt(cipherText []byte) ([]byte, error) {
	return cryptor.Decrypt(cipherText)
}
func Encrypt(textBytes []byte) ([]byte, error) {
	return cryptor.Encrypt(textBytes)
}
