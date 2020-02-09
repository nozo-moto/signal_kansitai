package main

import (
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logfile, err := os.OpenFile("./signal.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic("cannnot create log :" + err.Error())
	}
	defer logfile.Close()
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	c := make(chan os.Signal)
	d := make(chan int)
	t := time.NewTicker(60 * time.Second)
	signal.Notify(c)

	go func() {
		for {
			s := <-c
			log.Println("signal :", s)
			if s == syscall.SIGTERM {
				d <- 1
				return
			}
		}
	}()

	for {
		select {
		case <-d:
			return
		case <-t.C:
			log.Println("alive :")
		}

	}

}
