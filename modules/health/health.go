package health

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattmattox/supportability-collector/modules/cli"

	"github.com/mattmattox/supportability-collector/modules/logging"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var log = logging.SetupLogging()

var gitCommit string
var gitBranch string

func PrintVersion() {
	log.Printf("Current build version: %s", gitCommit)
	log.Printf("Current build branch: %s", gitBranch)
}

func StartHealthServer(cli.Cli) {
	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/healthz", HealthHandler)
		router.HandleFunc("/version", VersionHandler)
		router.Handle("/metrics", promhttp.Handler())
		address := "0.0.0.0:" + cli.Settings().HealthCheckPort
		if err := http.ListenAndServe(address, router); err != nil {
			log.Fatal(err)
		} else {
			log.Infoln("Health check server started")
		}
	}()
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Version: " + gitCommit))
}
