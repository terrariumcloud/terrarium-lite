package ssl

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"io"
	"math/big"
	"time"
)

type CA interface {
	GenerateRootCA()
	CreateClientCertificate()
}

type CertificateAuthority struct {
	Bits          int
	reader        io.Reader
	caTemplate    *x509.Certificate
	CAKey         *rsa.PrivateKey
	CACertificate *x509.Certificate
}

func (ca *CertificateAuthority) generatePrivateKey(bits int) error {
	privateKey, err := rsa.GenerateKey(ca.reader, bits)
	if err != nil {
		return err
	}
	ca.CAKey = privateKey
	return nil
}

func (ca *CertificateAuthority) generateRootCA() error {
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
	ca.caTemplate = caTemplate
	cert, err := x509.CreateCertificate(ca.reader, caTemplate, caTemplate, &ca.CAKey.PublicKey, ca.CAKey)
	if err != nil {
		return err
	}
	certEncoded, err := x509.ParseCertificate(cert)
	if err != nil {
		return err
	}
	ca.CACertificate = certEncoded
	return nil
}

func (ca *CertificateAuthority) GenerateRootCA(path string) error {
	err := ca.generatePrivateKey(ca.Bits)
	if err != nil {
		return err
	}

	err = ca.generateRootCA()
	if err != nil {
		return err
	}
	err = Write(path, "ca.key", GetPrivateKey(ca.CAKey))
	if err != nil {
		return err
	}
	err = Write(path, "ca.pub", GetPublicKey(&ca.CAKey.PublicKey))
	if err != nil {
		return err
	}
	err = Write(path, "ca.pem", GetCertificate(ca.CACertificate.Raw))
	if err != nil {
		return err
	}
	return nil
}

func (ca *CertificateAuthority) CreateClientCertificate(dns []string, bits int) (*x509.Certificate, error) {
	key, err := rsa.GenerateKey(ca.reader, bits)
	if err != nil {
		return nil, err
	}
	certTemplate := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Terrarium"},
		},
		DNSNames:              dns,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  false,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}
	cert, err := x509.CreateCertificate(ca.reader, certTemplate, ca.caTemplate, &key.PublicKey, ca.CAKey)
	if err != nil {
		return nil, err
	}
	certEncoded, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, err
	}
	return certEncoded, nil
}

func NewCA(bits int) *CertificateAuthority {
	ca := &CertificateAuthority{
		reader: rand.Reader,
		Bits:   bits,
	}
	return ca
}
