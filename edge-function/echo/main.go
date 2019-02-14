package main

import (
	"fmt"
	"github.com/computes/ipfs-http-api/pubsub"
	"net/url"
	"os"
	"time"
)

func main() {
	ipfsURL, _ := url.Parse("http://localhost:5001")

	for {
		pubsub.Publish(ipfsURL, fmt.Sprintf("/%s/result", os.Getenv("EDGE_APP_ID")), "echo")
		time.Sleep(time.Second)
	}
}
