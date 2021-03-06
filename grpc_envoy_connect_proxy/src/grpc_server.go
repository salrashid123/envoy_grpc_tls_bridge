package main

import (
	"echo"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	grpcport = flag.String("grpcport", ":50052", "grpcport")

	hs *health.Server

	conn *grpc.ClientConn
)

const (
	address string = ":50051"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *echo.EchoRequest) (*echo.EchoReply, error) {

	log.Println("Got rpc: --> ", in.Name)

	var h, err = os.Hostname()
	if err != nil {
		log.Fatalf("Unable to get hostname %v", err)
	}
	return &echo.EchoReply{Message: "Hello " + in.Name + "  from hostname " + h}, nil
}

func (s *server) SayHelloStream(in *echo.EchoRequest, stream echo.EchoServer_SayHelloStreamServer) error {

	log.Println("Got stream:  -->  ")
	stream.Send(&echo.EchoReply{Message: "Hello " + in.Name})
	stream.Send(&echo.EchoReply{Message: "Hello " + in.Name})

	return nil
}

type healthServer struct{}

func (s *healthServer) Check(ctx context.Context, in *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	log.Printf("Handling grpc Check request")
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (s *healthServer) Watch(in *healthpb.HealthCheckRequest, srv healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "Watch is not implemented")
}

func main() {

	flag.Parse()

	if *grpcport == "" {
		fmt.Fprintln(os.Stderr, "missing -grpcport flag (:50051)")
		flag.Usage()
		os.Exit(2)
	}

	// clientCaCert, err := ioutil.ReadFile("CA_crt.pem")
	// clientCaCertPool := x509.NewCertPool()
	// clientCaCertPool.AppendCertsFromPEM(clientCaCert)

	// certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	// if err != nil {
	// 	log.Fatalf("could not load server key pair: %s", err)
	// }

	// tlsConfig := tls.Config{
	// 	Certificates: []tls.Certificate{certificate},
	// }
	// creds := credentials.NewTLS(&tlsConfig)

	lis, err := net.Listen("tcp", *grpcport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	sopts := []grpc.ServerOption{grpc.MaxConcurrentStreams(10)}

	//sopts = append(sopts, grpc.Creds(creds))

	sopts = append(sopts)
	s := grpc.NewServer(sopts...)

	echo.RegisterEchoServerServer(s, &server{})

	healthpb.RegisterHealthServer(s, &healthServer{})

	log.Printf("Starting gRPC Server at %s", *grpcport)
	s.Serve(lis)

}
