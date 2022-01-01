package helper

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/erikjiang/grpc_demo/certs"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func GetServerCreds() credentials.TransportCredentials {
	// load server key & pem
	cert, err := tls.LoadX509KeyPair(certs.Path("server/server.pem"), certs.Path("server/server.key"))
	if err != nil {
		log.Fatalf("load x509 cert file failed, err: %s\n", err.Error())
	}
	// create cert pool and  append ca pem
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(certs.Path("ca/ca.pem"))
	if err != nil {
		log.Fatalf("read ca file failed, err: %s\n", err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	// set credentials
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
}

func GetClientCreds() credentials.TransportCredentials {
	// load client key & pem
	cert, err := tls.LoadX509KeyPair(certs.Path("client/client.pem"), certs.Path("client/client.key"))
	if err != nil {
		log.Fatalf("load x509 cert file failed, err: %s\n", err.Error())
	}
	// create cert pool and  append ca pem
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(certs.Path("ca/ca.pem"))
	if err != nil {
		log.Fatalf("read ca file failed, err: %s\n", err.Error())
	}
	certPool.AppendCertsFromPEM(ca)
	// set credentials
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost", // 与证书中的"CN"即:Common Name相同
		RootCAs:      certPool,
	})
}
