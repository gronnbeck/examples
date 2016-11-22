package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gronnbeck/examples/scaling-redis-k8s/setupredis"
)

var (
	addr      = flag.String("addr", ":8080", "http addr the service should be bound to. E.g. host:port.")
	redisAddr = flag.String("redis-addr", "localhost:6379", "tcp addr to connect redis tor")
	redisPass = flag.String("redis-pass", "", "password to redis server")
)

func init() {
	flag.Parse()

	if addr == nil || *addr == "" {
		panic("Addr cannot be empty")
	}
	if redisAddr == nil || *redisAddr == "" {
		panic("Reids addr cannot be empty")
	}
}

func main() {
	log.Println("Starting application")

	log.Printf("Waiting for redis on %v\n", *redisAddr)
	redisClient, err := setupredis.NewWait(*redisAddr, *redisPass)

	if err != nil {
		panic("Couldn't setup redis")
	}

	log.Println("Redis setup complete")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		redisResp := redisClient.Cmd("GET", "known-key")

		if redisResp.Err != nil {
			panic(err)
		}

		str, err := redisResp.Str()

		if err != nil {
			panic(err)
		}

		w.Write([]byte(str))

	})

	log.Println("HTTP server running: " + *addr)
	http.ListenAndServe(*addr, nil)
}
