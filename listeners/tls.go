package listeners

import (
	"crypto/tls"
	"log"
)

// NewTCP initializes and returns a new TCP listener, listening on an address.
func NewTLS(config Config) *TCP {
	cert, err := tls.LoadX509KeyPair(config.CertPath, config.KeyPath)
	if err != nil {
		log.Fatal(err)
	}
	// Basic TLS Config
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return NewTCP(Config{
		ID:        config.ID,
		Address:   config.Address,
		TLSConfig: tlsConfig,
	})
}
