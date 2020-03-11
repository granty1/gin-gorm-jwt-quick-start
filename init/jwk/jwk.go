package jwk

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/gin-cli/init/log"
	"io/ioutil"
	"os"
)

var (
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func Init() {
	err := GetRsaKey(2048)
	if err != nil {
		log.Log.Fatal("get rsa key failed, error:" + err.Error())
	}
}

//RSA公钥私钥产生
func GetRsaKey(bits int) error {
	// 如果已经生成过 Key ，则直接拿已经生成的 Key
	if isExistPath("./init/jwk/private.pem") {
		// 读取私钥

		privateFile, err := os.Open("./init/jwk/private.pem")
		if err != nil {
			return err
		}
		defer privateFile.Close()

		data, err := ioutil.ReadAll(privateFile)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(data)

		PrivateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return err
		}
		PublicKey = &PrivateKey.PublicKey

		return nil
	}

	// 生成私钥文件
	PrivateKey, _ = rsa.GenerateKey(rand.Reader, bits)

	derStream := x509.MarshalPKCS1PrivateKey(PrivateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("./init/jwk/private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	PublicKey = &PrivateKey.PublicKey
	derPKIX, err := x509.MarshalPKIXPublicKey(PublicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPKIX,
	}
	file, err = os.Create("./init/jwk/public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}

func isExistPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
