package deliveries

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"myfirstgosql/config"
	"myfirstgosql/repositories"
	"myfirstgosql/usecases"
	"net/http"
)

type ProductRoute struct {
	prefix  string
	useCase usecases.IProductUseCase
	rt      *mux.Router
}

func (pr *ProductRoute) InitRoute() {
	p := pr.rt.PathPrefix(pr.prefix).Subrouter()
	p.HandleFunc("", pr.getProductPagingHandler).Methods(http.MethodGet)
}
func (pr *ProductRoute) getProductPagingHandler(w http.ResponseWriter, r *http.Request) {
	result, err := pr.useCase.GetProductPaging(1, 3)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(result)
	w.Write([]byte(response))
}
func NewProductRoute(prefix string, sf *config.SessionFactory, rt *mux.Router) IAppRouter {
	repo := repositories.NewProductRepository(sf)
	usecase := usecases.NewProductUseCase(repo)
	return &ProductRoute{
		prefix,
		usecase,
		rt,
	}
}
