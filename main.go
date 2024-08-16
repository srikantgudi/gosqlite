package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/srikantgudi/gosqlite/db"
	"github.com/srikantgudi/gosqlite/views"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", rootPage)
	r.HandleFunc("/products", productsPage)
	r.HandleFunc("/customers", customersPage)
	r.HandleFunc("/customer/{oid}/orders", ordersPage)
	http.HandleFunc("/order/{oid}/details", orderDetailsPage)
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
	tmpl := template.Must(template.ParseFiles("views/base.html"))
	tmpl.Execute(w, "Go + HTMX + SQLITE + TailwindCSS")
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

func ordersPage(w http.ResponseWriter, r *http.Request) {
	custid := r.PathValue("oid")
	data := db.GetOrders(custid)
	views.Orders(data).Render(r.Context(), w)
}
func orderDetailsPage(w http.ResponseWriter, r *http.Request) {
	orderid := r.PathValue("oid")
	data, err := db.GetOrderdetails(orderid)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("order details:", data)
	orderTotal := 0.0
	for _, o := range data {
		orderTotal += o.UnitPrice * o.Quantity
	}
	order, _ := db.GetOrder(orderid)
	views.Orderdetails(data, orderTotal, order).Render(r.Context(), w)
}
