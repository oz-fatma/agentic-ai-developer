package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Day 39: graceful shutdown pattern (short server lifecycle)")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "running")
	})

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	srv := &http.Server{Handler: mux}
	url := "http://" + ln.Addr().String()
	fmt.Println("Server started at", url)

	go func() {
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			fmt.Println("serve error:", err)
		}
	}()

	resp, err := http.Get(url)
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("GET / -> %d %q\n", resp.StatusCode, string(body))
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutdown signal received")
	}()

	time.Sleep(50 * time.Millisecond)
	fmt.Println("Initiating graceful shutdown...")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("shutdown error:", err)
		return
	}
	fmt.Println("Server stopped cleanly")
}
