package main

import (
	"echo"
	"flag"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//healthpb "google.golang.org/grpc/health/grpc_health_v1"
	//"google.golang.org/grpc/status"
)

const ()

var (
	conn *grpc.ClientConn
)

func main() {

	address := flag.String("host", "localhost:8080", "host:port of gRPC server")
	// cacert := flag.String("cacert", "CA_crt.pem", "CACert for server")
	// serverName := flag.String("servername", "server.domain.com", "CACert for server")
	// flag.Parse()

	// var err error

	// caCert, err := ioutil.ReadFile(*cacert)
	// caCertPool := x509.NewCertPool()
	// caCertPool.AppendCertsFromPEM(caCert)

	// tlsConfig := tls.Config{
	// 	ServerName: *serverName,
	// 	RootCAs:    caCertPool,
	// }

	// creds := credentials.NewTLS(&tlsConfig)

	//conn, err = grpc.Dial(*address, grpc.WithTransportCredentials(creds))
	conn, err := grpc.Dial(*address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := echo.NewEchoServerClient(conn)
	ctx := context.Background()

	r, err := c.SayHello(ctx, &echo.EchoRequest{Name: "unary RPC msg "})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("RPC Response: %s", r)

}
