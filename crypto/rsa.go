package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// 公钥加密
func RsaEncryptByPub(pub_key_file string, origData []byte) ([]byte, error) {
	f, e := os.Open(pub_key_file)
	if e != nil {
		fmt.Println(e.Error())
		return nil, errors.New("public key file error")
	}
	bs, _ := ioutil.ReadAll(f)
	block, _ := pem.Decode(bs)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 私钥解密
func RsaDecryptByPri(pri_key_file string, ciphertext []byte) ([]byte, error) {
	f, e := os.Open(pri_key_file)
	if e != nil {
		fmt.Println(e.Error())
		return nil, errors.New("private key file error")
	}
	bs, _ := ioutil.ReadAll(f)
	block, _ := pem.Decode(bs)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
