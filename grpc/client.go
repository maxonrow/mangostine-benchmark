package grpc

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/mangostine-benchmark-go/grpc/mangostine_v0"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:26657"
	defaultName = "world"
)

//create the instance of grpc client
func GRPClient(tx []byte) {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewRpcClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	log.Printf("transaction name %v", name)
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SendRawTx(ctx, &pb.RawTransaction{RawTx: tx})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("transaction completed %v", r)
}
