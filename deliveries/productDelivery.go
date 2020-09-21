package deliveries

import (
	"fmt"
	"myfirstgosql/models"
)

type ProductDelivery struct {
}

func (pd *ProductDelivery) PrintOneProduct(result *models.Product) {
	fmt.Printf("%v\n", result)
}
func (pd *ProductDelivery) PrintProduct(result []*models.Product) {
	for _, p := range result {
		fmt.Println(p.Id, p.ProductCode, p.ProductName)
	}
}
func (pd *ProductDelivery) PrintTotalProduct(result int64) {
	fmt.Printf("Total product item %d\n", result)
}

func NewProductDelivery() *ProductDelivery {
	return &ProductDelivery{}
}
