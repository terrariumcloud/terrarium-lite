package ssl

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func Write(path string, filename string, data []byte) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		if err != nil {
			return err
		}
	}
	fileOut := fmt.Sprintf("%s/%s", path, filename)
	if os.IsNotExist(err) {
		err := os.WriteFile(fileOut, data, 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetPrivateKey(key *rsa.PrivateKey) []byte {
	bytes := x509.MarshalPKCS1PrivateKey(key)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: bytes,
	})
}

func GetPublicKey(key *rsa.PublicKey) []byte {
	bytes := x509.MarshalPKCS1PublicKey(key)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	})
}

func GetCertificate(cert []byte) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert,
	})
}
