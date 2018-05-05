package main

import (
	"log"

	"github.com/gearnode/judge/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err = grpc.Dial(":7777", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewJudgeClient(conn)

	response, err := c.SayHello(context.Background(), &api.PingMessage{Greeting: "foo"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Greeting)

	acl, err := c.Authorize(context.Background(), &api.AuthorizeRequest{Resource: "foo", Action: "foo"})
	if err != nil {
		log.Fatalf("Error when calling Authorize: %s", err)
	}
	log.Printf("Response from server: %T", acl.Permitted)
}
