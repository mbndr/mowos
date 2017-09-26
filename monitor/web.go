package monitor

import (
	"mime"
	"net/http"
	"path"
	"time"

	"github.com/mbndr/mowos"
)

// starts a web server
func startWebServer() error {
	api := http.NewServeMux()
	// TODO register API routes
	api.HandleFunc("/", apiHandler)

	router := http.NewServeMux()
	router.Handle("/api/", api)
	router.HandleFunc("/", assetHandler)

	addr := config.Monitor.ListenIP + ":" + config.Monitor.ListenPort

	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	mowos.Log.Info("listening on http " + addr)
	return srv.ListenAndServe()
}

// try to get the data of the compiled assets (from go-bindata)
func assetHandler(w http.ResponseWriter, r *http.Request) {
	p := r.URL.Path
	// use index file on root
	if p == "/" {
		p = "/index.html"
	}
	// only assets in web folder can be accessed
	data, err := mowos.Asset("web" + p)
	if err != nil {
		mowos.Log.Error(err)
		http.NotFound(w, r)
		return
	}
	// serve file
	mime := mime.TypeByExtension(path.Ext(p))
	w.Header().Set("Content-Type", mime)
	w.Write(data)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API Handler"))
}
