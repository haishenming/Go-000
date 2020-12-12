package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g.Go(func() error {
		return runServer(ctx)
	})
	g.Go(func() error {
		return RunSignal(ctx)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}

}

func runServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()

	return server.ListenAndServe()
}

func RunSignal(ctx context.Context) error {
	c := make(chan os.Signal)

	signal.Notify(c)

	select {
	case s := <-c:
		return errors.New(s.String())
	case <-ctx.Done():
		return ctx.Err()
	}
}
