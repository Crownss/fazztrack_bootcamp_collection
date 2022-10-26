package config

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/crownss/fazztrack_bootchamp/week_9/src/routers"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "serve",
	Short: "run server",
	RunE:  server,
}

func server(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.New(); err == nil {
		var addrs string = "0.0.0.0:8000"
		if run := os.Getenv("PORT"); run != "" {
			addrs = ":" + run
		}
		srv := &http.Server{
			Addr:         addrs,
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			IdleTimeout:  time.Minute,
			Handler:      mainRoute,
		}
		log.Println("running on", addrs)
		srv.ListenAndServe()
		return nil
	} else {
		return err
	}
}
