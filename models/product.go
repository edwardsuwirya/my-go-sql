package models

type Product struct {
	Id          string
	ProductCode string
	ProductName string
}

type TotalProduct struct {
	Count int64
}

type ProductWithPrice struct {
	Product
	ProductPrice
}
