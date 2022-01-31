package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/socheatsok78-lab/dedebme"
	"github.com/socheatsok78-lab/go-embed-frontend/dist"
)

func getFrontendAssets() http.FileSystem {
	return http.FS(dist.Assets)
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "The port to listen on")
	flag.Parse()

	httpFS := getFrontendAssets()
	handler := dedebme.Server(httpFS)

	http.Handle("/", handler)

	addr := fmt.Sprintf("localhost:%d", port)
	fmt.Printf("Serving app at http://%s\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
