package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"myfirstgosql/config"
	"myfirstgosql/deliveries"
	"myfirstgosql/repositories"
	"myfirstgosql/usecases"
)

type app struct {
	sf *config.SessionFactory
}

func newApp() app {
	c := config.NewConfig()
	err := c.InitDb()
	if err != nil {
		panic(err)
	}
	myapp := app{
		sf: c.SessionFactory,
	}
	return myapp
}

func (a app) run() {
	repo := repositories.NewProductRepository(a.sf)
	usecase := usecases.NewProductUseCase(repo)
	productDelivery := deliveries.NewProductDelivery()
	//
	//newProduct, err := usecase.RegisterNewProduct(models.Product{
	//	ProductCode: "GLG",
	//	ProductName: "Sabun Mandi",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//productDelivery.PrintOneProduct(newProduct)

	//_, err := usecase.RegisterNewProductWithPrice(models.ProductWithPrice{
	//	Product: models.Product{
	//		ProductCode: "BVV",
	//		ProductName: "Penanak Nasi",
	//	},
	//	ProductPrice: models.ProductPrice{
	//		Price: 137000,
	//	},
	//})
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("======= Simple Query =======")
	result, err := usecase.GetProductPaging(1, 3)
	if err != nil {
		panic(err)
	}
	productDelivery.PrintProduct(result)
	//
	//fmt.Println("======= Query with parameter =======")
	//queryFilter := "%Meja%"
	//result, err = usecase.GetProductByNameLike([]string{queryFilter})
	//if err != nil {
	//	panic(err)
	//}
	//productDelivery.PrintProduct(result)
	//
	//fmt.Println("======= Single Query Aggregation  =======")
	//result1, err := usecase.GetTotalProduct()
	//if err != nil {
	//	panic(err)
	//}
	//productDelivery.PrintTotalProduct(result1)
	//
	//fmt.Println("======= Multi Query  =======")
	//queryFilter2 := []string{"%Meja%", "G%"}
	//result, err = usecase.GetProductByNameLike(queryFilter2)
	//if err != nil {
	//	panic(err)
	//}
	//productDelivery.PrintProduct(result)

}
func main() {
	newApp().run()
}
