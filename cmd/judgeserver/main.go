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
	"os"

	"github.com/gearnode/judge/pkg/apiserver"
	"google.golang.org/grpc/credentials"
)

func main() {
	port := flag.Int("port", 8443, "The port")
	addr := flag.String("addr", "127.0.0.1", "The addr")
	crt := flag.String("tls-crt", "", "The cert")
	key := flag.String("tls-key", "", "The key")
	help := flag.Bool("help", false, "Help")
	tls := flag.Bool("tls", false, "Enable TLS")

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	srv := apiserver.Server{Port: *port, Addr: *addr}

	if *tls {
		creds, err := credentials.NewServerTLSFromFile(*crt, *key)
		if err != nil {
			fmt.Printf("could not load TLS keys: %s\n", err)
			os.Exit(1)
		}

		srv.Creds = &creds
	}

	srv.Start()
}