package ssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

type CA interface {
	GetPublicKey() []byte
	GetPrivateKey() []byte
	GetCertificate() []byte
	Write(path string) error
}

type CertificateAuthority struct {
	CAKey         *rsa.PrivateKey
	CACertificate *x509.Certificate
}

func (ca *CertificateAuthority) GetPrivateKey() []byte {
	bytes := x509.MarshalPKCS1PrivateKey(ca.CAKey)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: bytes,
	})
}

func (ca *CertificateAuthority) GetPublicKey() []byte {
	bytes := x509.MarshalPKCS1PublicKey(&ca.CAKey.PublicKey)
	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: bytes,
	})
}

func (ca *CertificateAuthority) GetCertificate() []byte {
	return pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ca.CACertificate.Raw,
	})
}

func (ca *CertificateAuthority) Write(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.Mkdir(path, 0700)
		if err != nil {
			return err
		}
	}
	certPath := fmt.Sprintf("%s/ca.pem", path)
	certKeyPath := fmt.Sprintf("%s/ca.key", path)
	certPubPath := fmt.Sprintf("%s/ca.pub", path)
	_, err = os.Stat(certPath)
	if os.IsNotExist(err) {
		err := os.WriteFile(certPath, ca.GetCertificate(), 0600)
		if err != nil {
			return err
		}
	}
	_, err = os.Stat(certKeyPath)
	if os.IsNotExist(err) {
		err := os.WriteFile(certKeyPath, ca.GetPrivateKey(), 0600)
		if err != nil {
			return err
		}
	}
	_, err = os.Stat(certPubPath)
	if os.IsNotExist(err) {
		err := os.WriteFile(certPubPath, ca.GetPublicKey(), 0600)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewCA(bits int) (*CertificateAuthority, error) {
	randomReader := rand.Reader
	privateKey, err := rsa.GenerateKey(randomReader, bits)
	if err != nil {
		return nil, err
	}
	caTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Terrarium"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(randomReader, caTemplate, caTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, err
	}
	certEncoded, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, err
	}
	return &CertificateAuthority{
		CAKey:         privateKey,
		CACertificate: certEncoded,
	}, nil
}
