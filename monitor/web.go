package monitor

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mbndr/mowos"
)

// starts a web server
func startWebServer() error {
	r := mux.NewRouter()

	// static files (html, css, js)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/web")))

	// api routes (called by frontend over ajax)
	// TODO

	addr := config.Monitor.ListenIP + ":" + config.Monitor.ListenPort

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	mowos.Log.Info("listening on http " + addr)
	return srv.ListenAndServe()
}
