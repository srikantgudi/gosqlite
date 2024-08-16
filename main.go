package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"gosqlite/db"
	"gosqlite/models"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", rootPage)
	r.HandleFunc("/products", productsPage)
	r.HandleFunc("/customers", customersPage)
	r.HandleFunc("/customer/{oid}/orders", ordersPage)
	r.HandleFunc("/order/{oid}/details", orderDetailsPage)
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
	tmpl := template.Must(template.ParseFiles("views/about.html"))
	tmpl.Execute(w, nil)
}

func productsPage(w http.ResponseWriter, r *http.Request) {
	data := db.GetProducts()
	tmpl := template.Must(template.ParseFiles("views/products.html"))
	tmpl.Execute(w, data)
}

func customersPage(w http.ResponseWriter, r *http.Request) {
	customers := db.GetCustomers()
	tmpl := template.Must(template.ParseFiles("views/customers.html"))
	tmpl.Execute(w, customers)
}

func ordersPage(w http.ResponseWriter, r *http.Request) {
	custid := r.PathValue("oid")
	data := db.GetOrders(custid)
	fmt.Println("orders page:", data)
	tmpl := template.Must(template.ParseFiles("views/orders.html"))
	tmpl.Execute(w, data)
}
func orderDetailsPage(w http.ResponseWriter, r *http.Request) {
	orderid := r.PathValue("oid")
	fmt.Println("order-id:", orderid)
	data, err := db.GetOrderdetails(orderid)
	if err != nil {
		log.Fatalln("error fetching details:", err.Error())
	}
	fmt.Printf("order details: %#v", data)
	orderTotal := 0.0
	for _, o := range data {
		orderTotal += o.LineTotal
	}
	order, _ := db.GetOrder(orderid)
	fmt.Println("order-details-page: order = ", order)
	tmpl := template.Must(template.ParseFiles("views/orderdetails.html"))
	type detail struct {
		Data  []models.Orderdetail
		Total float64
	}
	dtl := detail{Data: data, Total: orderTotal}
	tmpl.Execute(w, dtl)
}
