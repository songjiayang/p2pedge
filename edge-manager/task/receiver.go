package task

import (
	"edge-manager/util"
	"github.com/computes/ipfs-http-api/pubsub"
	"log"
)

func Listen(id *util.IpfsIdentity) {
	sub, err := pubsub.Subscribe(util.IpfsURL, id.ID+"/tasks")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Start edge node with id", id.ID)

	sub.Handle(new(receiver))
}

type receiver struct{}

func (_ *receiver) Recv(msg *pubsub.SubscriptionMessage) {
	msgStr, _ := msg.DataAsString()
	t := NewTask(msgStr)
	log.Printf("New task with cid: %s \n", msgStr)

	if err := t.Load(); err != nil {
		log.Println(err)
	} else {
		go func() {
			util.CheckError(t.Run())
		}()
	}
}
