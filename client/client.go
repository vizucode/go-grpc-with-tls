package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "gotlsgrpc/proto"
)

func TLSconf() (certPool *x509.CertPool) {
	cert, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	certPool = x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		log.Fatal("cert pool not daijobu!")
	}

	return certPool
}

func ClientTLS() {
	CAs := TLSconf()
	tls := credentials.NewTLS(&tls.Config{
		RootCAs: CAs,
	})

	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(tls))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	rpc := pb.NewPostmanClient(conn)

	data := &pb.Newspaper{Title: "Hello", Description: "World"}

	result, err := rpc.SendNewspaper(context.Background(), data)
	if err != nil {
		log.Println(err)
	}

	resp, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(resp))
}
