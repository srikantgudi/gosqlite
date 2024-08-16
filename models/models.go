package models

type Product struct {
	Id              int64   `db:"id" json:"id"`
	ProductName     string  `db:"product_name" json:"productname"`
	QuantityPerUnit string  `db:"quantity_per_unit" default:"-"`
	UnitPrice       float32 `db:"unit_price" json:"unitprice"`
	ReorderLevel    int64   `db:"reorder_level" json:"reorderlevel"`
}

type Customer struct {
	Id          string `db:"id"`
	CompanyName string `db:"company_name"`
	City        string `db:"city"`
	Ocount      int    `db:"ocount"`
}

type Order struct {
	Id        int     `db:"id" json:"id"`
	OrderDate string  `db:"orderdate" json:"orderdate"`
	ShipDate  *string `db:"shipdate" json:"shipdate"`
}

type CustomerOrders struct {
	Cust   Customer
	Orders []Order
}

type Orderdetail struct {
	Id          string  `db:"id"`
	ProductName string  `db:"productname"`
	UnitPrice   float64 `db:"unitprice"`
	Quantity    float64 `db:"quantity"`
	LineTotal   float64 `db:"linetotal"`
}
