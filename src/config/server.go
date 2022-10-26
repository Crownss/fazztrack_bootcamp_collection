package config

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/routers"
	"github.com/spf13/cobra"
	"github.com/rs/cors"
)

var ServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	RunE:  server,
}

func corsHandler() *cors.Cors {
	t := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000","http://localhost:3000/"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return t
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "0.0.0.0:8000"
		if run := os.Getenv("PORT"); run != "" {
			addrs = ":" + run
		}
		corss := cors.AllowAll()
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      corss.Handler(mainRoute),
		}
		log.Println("running on", addrs)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
