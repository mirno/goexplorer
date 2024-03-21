package config

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"github.com/mirno/goexplorer/pkg/utils"
)

// Docs: https://pkg.go.dev/github.com/mitchellh/mapstructure
type Config struct {
    Server ServerConfig `mapstructure:"server"`
    CACerts []string    `mapstructure:"caCerts"`
    Proxy   bool      `mapstructure:"proxyRequired"`
}

type ServerConfig struct {
    InstanceUrl        string      `mapstructure:"instanceUrl"`
    Protocol           string      `mapstructure:"protocol"`
    Port               int         `mapstructure:"port"`
    ProxyRequired      bool        `mapstructure:"proxyRequired"`
    AuthenticationType AuthConfig  `mapstructure:"authenticationType"`
}

type AuthConfig struct {
    MTLS   MTLSConfig   `mapstructure:"MTLS"`
    OAuth2 OAuth2Config `mapstructure:"OAUTH2"`
}

type MTLSConfig struct {
    CertFile string `mapstructure:"certFile"`
    KeyFile  string `mapstructure:"keyFile"`
}

type OAuth2Config struct {
    ClientID     string `mapstructure:"clientID"`
    ClientSecret string `mapstructure:"clientSecret"`
    RefreshToken string `mapstructure:"RefreshToken"`
}


func NewCertificatePool(config *Config ) (*x509.CertPool) {
	if len(config.CACerts) == 0 {
		log.Fatalf("No CA Certs specified")

	}
	// Define pool after verification
	caCertPool := x509.NewCertPool()
	

	for _, caCertPath := range config.CACerts {
		caCertPath, err := os.ReadFile(caCertPath)
		utils.Assert(err) 

		caCertPool.AppendCertsFromPEM(caCertPath)
	}

	return caCertPool
}

func LoadClientCert(config *MTLSConfig) (*tls.Certificate, error) {
	if config.CertFile == "" || config.KeyFile == "" {
		return nil, fmt.Errorf("CertFile or KeyFile not defined")
	}
	clientCert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
	utils.Assert(err)

	return &clientCert, nil

}

func LoadTLSConfig(config *MTLSConfig) error {
	log.Print("Placeholder for loading TLS based on the config")
	return nil
}