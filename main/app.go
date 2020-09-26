package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
	"log"
	"myfirstgosql/config"
	"myfirstgosql/deliveries"
	"myfirstgosql/repositories"
	"myfirstgosql/usecases"
	"net/http"
	"os"
)

var (
	appName    = "My Go Project"
	appVersion = "0.0.1"
	appTag     = "GO Sql Sample Project"
)

type app struct {
	cfg *config.Config
}

func newApp(cliCtx *cli.Context) app {
	env := cliCtx.String("env")
	c := config.NewConfig(env)
	err := c.InitDb()
	if err != nil {
		panic(err)
	}
	myapp := app{
		cfg: c,
	}
	return myapp
}

func (a app) runApi() {
	fmt.Println("Run HTTP Mode")
	appRouter := mux.NewRouter()
	hostListen := fmt.Sprintf("%v:%v", a.cfg.HttpConf.Host, a.cfg.HttpConf.Port)
	deliveries.NewAppDelivery(appRouter, a.cfg.SessionFactory).Initialize()
	log.Printf("Ready to listen on %v", hostListen)
	if err := http.ListenAndServe(hostListen, appRouter); err != nil {
		log.Panic(err)
	}
}

func (a app) runMigration(mode string) {
	fmt.Println("Run Migration Mode")
	fmt.Printf("%s schema %s\n", a.cfg.GetEnv("dbhost", ""), a.cfg.GetEnv("dbschema", ""))
	migrationPath := a.cfg.GetEnv("DBMIGRATIONFILE", "")
	dbMigration(a.cfg.SessionFactory, migrationPath, mode)
}
func (a app) run() {
	fmt.Println("Run CLI Mode")
	repo := repositories.NewProductRepository(a.cfg.SessionFactory)
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
		Name:        appName,
		Version:     appVersion,
		Description: appTag,
		Action: func(c *cli.Context) error {
			switch c.String("mode") {
			case "cli":
				newApp(c).run()
			case "migration-up":
				newApp(c).runMigration("up")
			case "migration-down":
				newApp(c).runMigration("down")
			case "http":
				newApp(c).runApi()
			default:
				panic("Unknown mode")
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				EnvVars:     []string{"APPENV"},
				Usage:       "Application's environment",
				DefaultText: "dev",
			},
			&cli.StringFlag{
				Name:        "mode",
				Usage:       "Application's running mode, options are cli, migration-up,migration-down, http",
				Aliases:     []string{"m"},
				EnvVars:     []string{"APPMODE"},
				DefaultText: "http",
				Value:       "http",
			},
		},
	}
	err := appConfig.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
