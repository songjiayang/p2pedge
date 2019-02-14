package main

import (
	"encoding/json"
	"fmt"
	"github.com/computes/ipfs-http-api"
	"github.com/computes/ipfs-http-api/pubsub"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

var (
	recver *Recver
)

type Recver struct{}

func (_ *Recver) Recv(msg *pubsub.SubscriptionMessage) {
	msgStr, _ := msg.DataAsString()
	fmt.Println(msgStr)
}

func main() {
	if len(os.Args) < 3 || (os.Args[1] == "add" && len(os.Args) < 4) {
		log.Panic("miss args")
	}

	cmd := os.Args[1]

	if cmd == "add" {
		resp, err := http.DefaultClient.Get(fmt.Sprintf("http://localhost:5001/api/v0/pubsub/pub?arg=%s/tasks&arg=%s", os.Args[2], os.Args[3]))
		checkError(err)
		defer resp.Body.Close()
		log.Printf("Add successful, you can use `edge-ctl data %s` to get the result.", os.Args[3])
		return
	}

	if cmd == "data" {
		ipfsURL, _ := url.Parse("http://localhost:5001")

		body, err := ipfs.Cat(ipfsURL, os.Args[2])
		checkError(err)
		defer body.Close()

		buf, _ := ioutil.ReadAll(body)

		var cfg map[string]string
		checkError(json.Unmarshal(buf, &cfg))

		sub, err := pubsub.Subscribe(ipfsURL, fmt.Sprintf("/%s/result", cfg["id"]))
		if err != nil {
			log.Panic(err)
		}
		sub.Handle(recver)

		sigs := make(chan os.Signal, 1)
		done := make(chan bool, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		go func() {
			<-sigs
			done <- true
		}()
		<-done
	}
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
