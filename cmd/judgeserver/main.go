/*
Copyright 2018 The Judge Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/gearnode/judge/pkg/apiserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 5053, "Runs judgeserver on the specified port")
	addr := flag.String("addr", "127.0.0.1", "Binds judgeserver to the specified IP")
	crt := flag.String("tls-crt", "", "The certificate for the TLS mode")
	key := flag.String("tls-key", "", "The certificate key for TLS mode")
	help := flag.Bool("help", false, "Shows this help message")
	tls := flag.Bool("tls", false, "Enable TLS mode")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Starting Judge apiserver on %s:%d\n", *addr, *port)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *addr, *port))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	srv := grpc.NewServer()

	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*crt, *key)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		srv = grpc.NewServer(grpc.Creds(creds))
	}

	reflection.Register(srv)
	apiserver.Register(srv)

	srv.Serve(lis)
}
