package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lwithers/aa-sms-receiver/client"
)

var (
	addr = flag.String("addr", "peridot.lwithers.me.uk:8123",
		"server address to query")
)

func main() {
	cl := client.New(*addr)
	msg, err := cl.GetMessage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("OK: %q\n", msg)
}
