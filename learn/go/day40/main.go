package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/oz-fatma/agentic-ai-developer/learn/go/day40/config"
	"github.com/oz-fatma/agentic-ai-developer/learn/go/day40/middleware"
)

func main() {
	fmt.Println("Day 40: config + context + middleware + shutdown")

	cfg, err := config.LoadFromEnv()
	if err != nil {
		log.Fatal("config:", err)
	}
	fmt.Printf("Config: addr=%s log=%s\n", cfg.Addr(), cfg.LogLevel)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /info", infoHandler)

	handler := middleware.Chain(mux, middleware.Timeout(5*time.Second), middleware.RequestID, middleware.Logging)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal("listen:", err)
	}

	srv := &http.Server{
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	url := "http://" + ln.Addr().String()
	fmt.Println("Server at", url)

	go func() {
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Println("serve:", err)
		}
	}()

	resp, err := http.Get(url + "/health")
	if err == nil {
		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("GET /health -> %d %s (request-id: %s)\n", resp.StatusCode, string(body), resp.Header.Get("X-Request-ID"))
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go func() {
		<-ctx.Done()
		fmt.Println("Shutdown signal ready")
	}()

	time.Sleep(50 * time.Millisecond)
	fmt.Println("Shutting down gracefully...")
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("shutdown:", err)
	}
	fmt.Println("Done")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	id := middleware.RequestIDFromContext(r.Context())
	writeJSON(w, http.StatusOK, map[string]string{
		"request_id": id,
		"message":    "day 40 combined demo",
	})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
