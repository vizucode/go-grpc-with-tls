package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "gotlsgrpc/proto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type newspaper struct {
	pb.UnimplementedPostmanServer
}

func NewPostman() pb.PostmanServer {
	return &newspaper{}
}

func (nw *newspaper) SendNewspaper(ctx context.Context, req *pb.Newspaper) (*pb.RespNewspaper, error) {

	resp := &pb.RespNewspaper{
		ResponseMessage: "Send Successfully",
		Data:            []*pb.Newspaper{req},
	}

	return resp, nil
}

func TLSConf() (ca tls.Certificate, crtPool *x509.CertPool) {
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}

	crtPool = x509.NewCertPool()
	certByte, err := os.ReadFile("server.crt")
	if err != nil {
		log.Fatal(err)
	}

	if ok := crtPool.AppendCertsFromPEM(certByte); !ok {
		log.Fatal("failed to append certificate")
	}

	return cert, crtPool
}

func NewServer() {
	tcpCon, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	cert, caPool := TLSConf()
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caPool,
	})

	server := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterPostmanServer(server, NewPostman())

	fmt.Println("gRPC is running on port 50051 [secured with TLS]")

	err = server.Serve(tcpCon)
	if err != nil {
		log.Fatal(err)
	}
}
