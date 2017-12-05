package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const connectTimeout = 5 * time.Second

// EINVAL Invalid argument
const EINVAL = 22

func grpcConnect(addr string) {
	fmt.Printf("connecting to %v\n", addr)
	con, err := grpc.Dial(addr,
		grpc.WithTimeout(connectTimeout),
		grpc.WithInsecure(),
		grpc.WithBlock())
	if err != nil {
		fmt.Println("connection could not be established")
		os.Exit(1)
	}

	fmt.Println("connection established")

	// if the connection is immediately closed after it's opened,
	// a "broken pipe" error message will be logged on the server-side.
	// Wait a little bit to not cause this error message getting logged on
	// the server.
	time.Sleep(100 * time.Millisecond)
	con.Close()
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		name := path.Base(os.Args[0])

		fmt.Printf("Usage: %v <GRPC-Host:Port>...\n", name)
		fmt.Println()
		fmt.Println("grpc-healthcheck checks if GRPC services are reachable.")
		fmt.Println("It tries to establish a connection to each of the passed addresses.")
		fmt.Println("When a connection can't be established it terminates.")
		fmt.Println()
		fmt.Printf("\nExit Codes:\n")
		fmt.Printf(" 0  - connection established\n")
		fmt.Printf(" 1  - connection error\n")
		fmt.Printf(" %d - invalid commandline parameters\n", EINVAL)
		fmt.Println()
		os.Exit(EINVAL)
	}

	grpclog.SetLogger(log.New(ioutil.Discard, "", 0))

	for _, addr := range os.Args[1:] {
		grpcConnect(addr)
	}

	os.Exit(0)
}
