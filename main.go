package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var (
	pprof        = flag.String("pprof", "", "start a pprof http server (https://golang.org/pkg/net/http/pprof/)")
	freeosmemory = flag.Bool("freeosmemory", true, "aggressively free memory back to the OS")
	addr         = flag.String("addr", ":4389", "server listen address (tcp or websocket)")
)

func freeOSMemory() {
	for {
		time.Sleep(1 * time.Second)
		//		debug.FreeOSMemory()
	}
}

func main() {
	fmt.Println("vim-go")
	flag.Parse()
	// Start pprof server, if desired.
	if *pprof != "" {
		go func() {
			log.Println(http.ListenAndServe(*pprof, nil))
		}()
	}

	if *freeosmemory {
		go freeOSMemory()
	}

	listen := func(addr string) (*net.Listener, error) {
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Could not bind to address %s: %v", addr, err)
			return nil, err
		}
		if os.Getenv("TLS_CERT") != "" && os.Getenv("TLS_KEY") != "" {
			cert, err := tls.X509KeyPair([]byte(os.Getenv("TLS_CERT")), []byte(os.Getenv("TLS_KEY")))
			if err != nil {
				return nil, err
			}
			listener = tls.NewListener(listener, &tls.Config{
				Certificates: []tls.Certificate{cert},
			})
		}
		return &listener, nil
	}
	lis, err := listen(*addr)
	if err != nil {
		panic(err)
	}
	defer (*lis).Close()
	log.Println("langserver-go: listening for TCP connections on", *addr)

	//connectionCount := 0

	for {
		//conn, err := (*lis).Accept()
		_, err := (*lis).Accept()
		if err != nil {
			panic(err)
		}
		log.Println("langserver-go: connected")
	}
}
