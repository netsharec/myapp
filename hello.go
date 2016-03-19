package hello

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"github.com/pressly/chi"
)

func init() {
	var router = chi.NewRouter()
	router.Get("/", handler)
	router.Get("/:id", getId)
	http.Handle("/", router)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func getId(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParams(ctx)["id"]
	fmt.Fprintf(w, "id: %v", id)
}
