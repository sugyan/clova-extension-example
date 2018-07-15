package main

import (
	"io/ioutil"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf(ctx, "error: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	log.Infof(ctx, "body (%d): %s", len(b), string(b))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
