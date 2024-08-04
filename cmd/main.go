package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"osquerything/pkg/osquery/table"

	"github.com/osquery/osquery-go"
)

func main() {
	fmt.Println("initializing...")
	socket := flag.String("socket", "", "Path to osquery socket file")
	flag.Parse()
	if *socket == "" {
		log.Fatalf(`Usage: %s --socket SOCKET_PATH`, os.Args[0])
	}

	server, err := osquery.NewExtensionManagerServer("software_update", *socket)
	if err != nil {
		log.Fatalf("Error creating extension: %s\n", err)
	}

	table.SoftwareUpdate(server)

	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
