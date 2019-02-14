package main

import (
	"edge-manager/task"
	"edge-manager/util"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	id, err := util.GetIpfsNodeID()
	if err != nil {
		log.Panic(err)
	}

	task.Listen(id)

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()
	<-done
}
