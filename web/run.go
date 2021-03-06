package web

import (
	dsc "github.com/jeffjen/go-discovery/info"
	api "github.com/jeffjen/podd/web/api"

	log "github.com/Sirupsen/logrus"

	"net/http"
)

func init() {
	s := api.GetServeMux()

	// Register static polymer assets
	asset_repo := api.Dir{
		http.Dir("html/bower_components/"),
		http.Dir("html/custom_components/"),
	}
	s.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(asset_repo)))

	// Register static html resources
	s.Handle("/css", http.FileServer(http.Dir("html/www/css/")))
	s.Handle("/js", http.FileServer(http.Dir("html/www/js/")))
	s.Handle("/", http.FileServer(http.Dir("html/www/")))

	// API for cluster
	s.HandleFunc("/cluster/list", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/www/cluster-list.json")
	})
	// API for service
	s.HandleFunc("/service/list", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/www/service-list.json")
	})

	s.HandleFunc("/info", dsc.Info)
}

func RunAPIEndpoint(addr string) {
	server := api.GetServer()
	server.Addr = addr
	log.Error(server.ListenAndServe())
}
