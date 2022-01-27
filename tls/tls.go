package tls

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"github.com/ebauman/simpleca/file"
	"github.com/ebauman/simpleca/parse"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	CACertFileName = "ca.pem"
	CAKeyFileName  = "ca.key"

	CertFileName = "cert.pem"
	KeyFileName  = "key.pem"

	PEMCertificate = "CERTIFICATE"
	PEMKey         = "RSA PRIVATE KEY"
)

type CertConfig struct {
	Name               string
	Path               string
	Passphrase         string
	Country            string
	State              string
	Locality           string
	Organization       string
	OrganizationalUnit string
	CommonName         string
	IPAddresses        []net.IP
	DNSNames           []string
	EmailAddresses     []string
	URIs               []string
	ExpireIn           string
}

func FullCAPath(conf *CertConfig) string {
	return fmt.Sprintf("%s/%s", conf.Path, conf.Name)
}

func FullCertPath(conf *CertConfig, caName string) string {
	return fmt.Sprintf("%s/%s/%s", conf.Path, caName, conf.Name)
}

func GenerateCA(conf *CertConfig) error {
	cert, err := genCert(conf)
	if err != nil {
		return err
	}

	cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}
	cert.IsCA = true
	cert.KeyUsage = x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign

	caPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, cert, cert, &caPrivKey.PublicKey, caPrivKey)
	if err != nil {
		return err
	}

	caPEM := new(bytes.Buffer)
	err = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})
	if err != nil {
		return err
	}

	caPrivKeyPEM := new(bytes.Buffer)
	err = pem.Encode(caPrivKeyPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(caPrivKey),
	})
	if err != nil {
		return err
	}

	// now, write these out to file so we can use them in the future
	var fullPath = FullCAPath(conf)
	var writePath = fmt.Sprintf("%s/%s", fullPath, CACertFileName)

	err = os.WriteFile(writePath, caPEM.Bytes(), 0700)
	if err != nil {
		return err
	}

	writePath = fmt.Sprintf("%s/%s", fullPath, CAKeyFileName)
	err = os.WriteFile(writePath, caPrivKeyPEM.Bytes(), 0700)

	return err
}

func LoadCert(certPath string, keyPath string) (*rsa.PrivateKey, *x509.Certificate, error) {
	certPem, err := os.ReadFile(certPath)
	if err != nil {
		return nil, nil, err
	}

	cert, err := decodeCertificate(certPem)
	if err != nil {
		return nil, nil, err
	}

	keyPem, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, nil, err
	}

	key, err := decodeKey(keyPem)
	if err != nil {
		return nil, nil, err
	}

	return key, cert, nil
}

func LoadCA(certPath string, keyPath string) (*rsa.PrivateKey, *x509.Certificate, error) {
	key, cert, err := LoadCert(certPath, keyPath)
	if err != nil {
		return nil, nil, err
	}

	if !cert.IsCA {
		return nil, nil, fmt.Errorf("certificate is not an authority")
	}

	return key, cert, nil
}

func SignCert(conf *CertConfig, caName string) error {
	caPrivPath := fmt.Sprintf("%s/%s/%s", conf.Path, caName, CAKeyFileName)
	caCertPath := fmt.Sprintf("%s/%s/%s", conf.Path, caName, CACertFileName)
	caPriv, caCert, err := LoadCA(caCertPath, caPrivPath)
	if err != nil {
		return err
	}

	cert, err := genCert(conf)
	if err != nil {
		return err
	}

	cert.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}
	cert.KeyUsage = x509.KeyUsage(x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment)
	cert.SubjectKeyId = []byte(strconv.FormatInt(time.Now().Unix(), 10)) // HACK

	certPrivKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, cert, caCert, &certPrivKey.PublicKey, caPriv)
	if err != nil {
		return err
	}

	certPEM, err := PEMEncodeCert(certBytes)
	if err != nil {
		return err
	}

	keyPEM, err := PEMEncodeKey(certPrivKey)
	if err != nil {
		return err
	}

	var fullPath = FullCertPath(conf, caName)

	if err = file.CheckPath(fullPath); err != nil {
		return err
	}

	var writePath = fmt.Sprintf("%s/%s", fullPath, CertFileName)

	err = os.WriteFile(writePath, certPEM, 0700)
	if err != nil {
		return err
	}

	writePath = fmt.Sprintf("%s/%s", fullPath, KeyFileName)
	err = os.WriteFile(writePath, keyPEM, 0700)

	return err
}

func decodeCertificate(data []byte) (*x509.Certificate, error) {
	pBlock, _ := pem.Decode(data)

	if pBlock == nil {
		return nil, fmt.Errorf("error decoding pem data from file")
	}

	cert, err := x509.ParseCertificate(pBlock.Bytes)
	if err != nil {
		fmt.Println(string(data))
		return nil, fmt.Errorf("error parsing certificate from pem block: %s", err)
	}

	return cert, nil
}

func decodeKey(data []byte) (*rsa.PrivateKey, error) {
	pBlock, _ := pem.Decode(data)

	if pBlock == nil {
		return nil, fmt.Errorf("error decoding pem data from file")
	}

	if pBlock.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("invalid private key, not rsa private key")
	}

	key, err := x509.ParsePKCS1PrivateKey(pBlock.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing rsa private key: %v", err)
	}

	return key, nil
}

func genCert(conf *CertConfig) (*x509.Certificate, error) {
	uris, err := parse.ParseURIs(conf.URIs)
	if err != nil {
		return nil, err
	}

	var notAfter time.Time
	if conf.ExpireIn != "" {
		d, err := parse.ParseDuration(conf.ExpireIn)
		if err != nil {
			return nil, err
		}
		notAfter = *d
	} else {
		notAfter = time.Now().AddDate(10, 0, 0)
	}

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(time.Now().Unix()),
		Subject: pkix.Name{
			Country:            []string{conf.Country},
			Province:           []string{conf.State},
			Locality:           []string{conf.Locality},
			Organization:       []string{conf.Organization},
			OrganizationalUnit: []string{conf.OrganizationalUnit},
			CommonName:         conf.CommonName,
		},
		NotBefore:             time.Now(),
		NotAfter:              notAfter,
		BasicConstraintsValid: true,
		DNSNames:              conf.DNSNames,
		EmailAddresses:        conf.EmailAddresses,
		IPAddresses:           conf.IPAddresses,
		URIs:                  uris,
	}

	return cert, nil
}

func PEMEncodeCert(cert []byte) ([]byte, error) {
	certPEM := new(bytes.Buffer)
	err := pem.Encode(certPEM, &pem.Block{
		Type:  PEMCertificate,
		Bytes: cert,
	})
	if err != nil {
		return nil, err
	}

	return certPEM.Bytes(), nil
}

func PEMEncodeKey(key *rsa.PrivateKey) ([]byte, error) {
	keyPEM := new(bytes.Buffer)
	err := pem.Encode(keyPEM, &pem.Block{
		Type:  PEMKey,
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	if err != nil {
		return nil, err
	}

	return keyPEM.Bytes(), nil
}
