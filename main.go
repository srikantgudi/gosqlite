package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/srikantgudi/gosqlitetempl/db"
	"github.com/srikantgudi/gosqlitetempl/views"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", rootPage)
	r.HandleFunc("/products", productsPage)
	r.HandleFunc("/customers", customersPage)
	r.HandleFunc("/about", aboutPage)

	fmt.Println("Running on http://localhost:8090")
	s := &http.Server{
		Handler:        r,
		Addr:           ":8090",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func rootPage(w http.ResponseWriter, r *http.Request) {
	views.Base("Go+HTMX+Templ+SQLite").Render(r.Context(), w)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about page...")
	views.About().Render(r.Context(), w)
}

func productsPage(w http.ResponseWriter, r *http.Request) {
	data := db.GetProducts()
	views.Products(data).Render(r.Context(), w)
}

func customersPage(w http.ResponseWriter, r *http.Request) {
	customers := db.GetCustomers()
	views.Customers(customers).Render(r.Context(), w)
}
