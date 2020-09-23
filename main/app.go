package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/urfave/cli/v2"
	"myfirstgosql/config"
	"myfirstgosql/deliveries"
	"myfirstgosql/repositories"
	"myfirstgosql/usecases"
	"os"
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

func (a app) runMigration() {
	dbMigration(a.sf)
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
	ascii :=
		`
 _______         __                       
|    ___|.-----.|__|.-----.--------.---.-.
|    ___||     ||  ||  _  |        |  _  |
|_______||__|__||__||___  |__|__|__|___._|
                    |_____|

My Go SQL

`
	fmt.Println(ascii)
	appConfig := &cli.App{
		Name:        config.GetEnv("appname", "My Go Project"),
		Version:     config.GetEnv("appversion", "0.0.0"),
		Description: config.GetEnv("apptag", ""),
		Action: func(c *cli.Context) error {
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "cli",
				Aliases: []string{"c"},
				Usage:   "Run console based application",
				Action: func(c *cli.Context) error {
					newApp().run()
					return nil
				},
			}, {
				Name:    "migration",
				Aliases: []string{"d"},
				Usage:   "Run database migration",
				Action: func(c *cli.Context) error {
					newApp().runMigration()
					return nil
				},
			},
		},
	}
	err := appConfig.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
