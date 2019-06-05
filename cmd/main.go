package main

import (
    "context"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "os"
    "os/signal"
    "time"
	"../internal/pkg/streams"
)

var storedStr map[string]string

func main() {
        storedStr = make(map[string]string)

        r := mux.NewRouter()
         r.HandleFunc("/v1/streams/{streamId}", streams.GetStreamHandler).Methods("GET")

        srv := &http.Server{
                Addr: "0.0.0.0:80",
                Handler: r,
        }

        go func() {
                log.Println("Starting server::")
                if err := srv.ListenAndServe(); err != nil {
                        log.Println("error binding:", err)
                }
        }()

        c:=make(chan os.Signal, 1)
        signal.Notify(c, os.Interrupt)
        <-c

        ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
        defer cancel()
        srv.Shutdown(ctx)

        log.Println("Server shut down")
        os.Exit(0)
}
