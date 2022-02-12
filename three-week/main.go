package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	timeOutCtx, _ := context.WithTimeout(ctx, 5*time.Second)
	g, errCtx := errgroup.WithContext(timeOutCtx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello world1"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	channel := make(chan os.Signal, 1)
	signal.Notify(channel)

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				fmt.Println("server.Shutdown", errCtx.Err())
				return server.Shutdown(errCtx)
			case <-channel:
				cancel()
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println("Service shutdown err: ", err)
	}
}
