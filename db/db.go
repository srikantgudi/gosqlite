package db

import (
	"fmt"
	"log"

	"github.com/srikantgudi/gosqlitetempl/models"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

var sqlProd string = "Select id, product_name, quantity_per_unit, unit_price, reorder_level From Products"
var sqlCustomers string = "SELECT c.id, c.company_name, c.city, COUNT(c.id) AS ocount FROM Customers c LEFT JOIN Orders o ON c.id=o.customer_id GROUP BY c.id HAVING ocount > 0"
var sqlCust string = "SELECT c.id, c.company, c.city FROM Customers c"
var sqlOrders string = "Select id, date_format(order_date, '%d-%b-%Y') orderdate, date_format(shipped_date, '%d-%b-%Y') shipdate from Orders"

func init() {
	// db = sqlx.MustConnect("mysql", "root:nimdaroot@tcp(localhost:3306)/northwind")
	db = sqlx.MustConnect("sqlite3", "./northwind.db")
}

func GetProducts() []models.Product {
	data := []models.Product{}
	fmt.Println("sqlProd:", sqlProd)
	err := db.Select(&data, sqlProd)
	if err != nil {
		log.Fatalln("error fetch:", err.Error())
	}
	return data
}

func SearchProducts(text string) []models.Product {
	data := []models.Product{}
	sql := fmt.Sprintf("%v Where concat(product_name,quantity_per_unit,list_price,reorder_level) regexp '%v'", sqlProd, text)
	db.Select(&data, sql)
	return data
}

func GetCustomers() []models.Customer {
	data := []models.Customer{}
	db.Select(&data, sqlCustomers)
	return data
}

func GetCustomer(custid string) models.Customer {
	data := models.Customer{}
	sqlStr := fmt.Sprintf("%v Where id=? Limit 1", sqlCust)
	db.Get(&data, sqlStr, custid)
	return data
}

func GetOrders(custid string) models.CustomerOrders {
	custOrders := models.CustomerOrders{}
	sqlStr := fmt.Sprintf("%v Where customer_id=?", sqlOrders)
	err := db.Select(&custOrders.Orders, sqlStr, custid)
	if err != nil {
		log.Fatalf("\n%v\n", err.Error())
	}
	custOrders.Cust = GetCustomer(custid)
	return custOrders
}

func GetOrder(orderid string) (models.Order, error) {
	data := models.Order{}
	err := db.Get(&data, fmt.Sprintf("%v Where id=? Limit 1", sqlOrders), orderid)
	return data, err
}

func GetOrderdetails(orderid string) ([]models.Orderdetail, error) {
	data := []models.Orderdetail{}
	err := db.Select(&data, "Select p.product_name productname, round(od.quantity,2) quantity, round(od.unit_price,2) unitprice, round(od.quantity * od.unit_price,2) linetotal From order_details od Join products p on p.id = od.product_id Where od.order_id = ?", orderid)
	return data, err
}
