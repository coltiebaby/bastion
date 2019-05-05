// Utilities sets up the transport for the clients to use
// Will default to InsecureSkipVerify
// Enable it by setting the environmental variable `LCU_SSH`
// to the actual path or by placing the cert in your in your documents.
//
// unix/mac ~/Documents/riot.pem
// windows C:\Documents\
//
// Checks env -> file -> doesn't care

package clientutil

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Adds a certificate in memory to the config.
func addTLSCert(pemfile string) (*tls.Config, error) {

	certs, err := x509.SystemCertPool()
	if err != nil {
		return &tls.Config{}, err
	}

	cert_data, err := ioutil.ReadFile(pemfile)
	if err != nil {
		return &tls.Config{}, err
	}

	if ok := certs.AppendCertsFromPEM(cert_data); !ok {
		fmt.Println("No certs appended, using system certs only")
		return &tls.Config{}, errors.New("Could not append cert!")
	}

	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            certs,
	}

	return config, nil
}

func ignoreVerify() *tls.Config {
	return &tls.Config{
		InsecureSkipVerify: true,
	}
}

func getTLSConfig() (*tls.Config, error) {
	if fp := os.Getenv(envkey); fp != "" {
		return addTLSCert(fp)
	} else if _, err := os.Stat(DEFAULT_PEMFILE); err == nil {
		return addTLSCert(DEFAULT_PEMFILE)
	} else {
		config := ignoreVerify()
		return config, nil
	}
}

func newHttpClient() *http.Client {
	certs, err := getTLSConfig()
	if err != nil {
		panic(INVALID_CERT_ERR)
	}

	transport := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       time.Second * 90,
		TLSHandshakeTimeout:   time.Second * 10,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       certs,
	}

	return &http.Client{Transport: transport}
}

const (
	envkey string = `LCU_SSH`
)

var (
	INVALID_CERT_ERR error = errors.New("Failed to load in the riot cert")
	HTTP_CLIENT_ERR  error = errors.New("Failed to create an http client")
)

var (
	// Add a do once to this when config gets added
	HttpClient *http.Client = newHttpClient()
)
