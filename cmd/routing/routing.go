package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/*
BaseRouter Returns only business routes
*/
func BaseRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/home", HomeHandler)
	r.HandleFunc("/products/{key}", ProductHandler)
	r.HandleFunc("/articles/{category}/", ArticlesCategoryHandler)

	return r
}

/*
HomeHandler Home Handler
*/
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Home")
	log.Printf("/home called")
}

/*
ProductHandler Product Handler
*/
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product: %v\n", vars["key"])
	log.Printf("/Looked up a product")
}

/*
ArticlesCategoryHandler Article Handler
*/
func ArticlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Articles of category: %v\n", vars["category"])
	log.Printf("/Looked for articles of a category")
}
