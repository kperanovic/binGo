package main

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"github.com/kperanovic/tombola/internal/engine"
	"github.com/kperanovic/tombola/internal/socketio"
	"github.com/kperanovic/tombola/internal/store/memory"
	"go.uber.org/zap"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx := context.TODO()

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	log.Info("starting service")
	defer log.Info("service closed")

	// Initiate engine
	e := engine.NewEngine(ctx, log, memory.NewMemoryStore())

	// Initiate socketio
	socket := socketio.NewSocketIO(ctx, log, e)

	if err := socket.Init(); err != nil {
		log.Fatal("error initializing socketio", zap.Error(err))
	}

	go socket.Server.Serve()
	defer socket.Server.Close()

	// r := mux.NewRouter()

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "null")

		socket.Server.ServeHTTP(w, r)
	})

	log.Info("http server started", zap.String("address", ":8000"))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("error starting http server", zap.Error(err))
	}
}
