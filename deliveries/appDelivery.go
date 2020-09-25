package deliveries

import (
	"github.com/gorilla/mux"
	"myfirstgosql/config"
)

const (
	AUTH_MAIN_ROUTE    = "/auth"
	PRODUCT_MAIN_ROUTE = "/product"
)

type IAppRouter interface {
	InitRoute()
}
type AppDelivery struct {
	rt *mux.Router
	sf *config.SessionFactory
}

func NewAppDelivery(rt *mux.Router, sf *config.SessionFactory) *AppDelivery {
	return &AppDelivery{rt, sf}
}
func (a *AppDelivery) Initialize() {
	//repo := repositories.NewProductRepository(a.sf)
	//usecase := usecases.NewProductUseCase(repo)
	//productDelivery := deliveries.NewProductDelivery()
	//
	routerList := []IAppRouter{
		NewProductRoute(PRODUCT_MAIN_ROUTE, a.sf, a.rt),
		NewAuthRoute(AUTH_MAIN_ROUTE, a.sf, a.rt),
	}

	for _, rt := range routerList {
		rt.InitRoute()
	}
}

//func (a *AppDelivery) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	a.rt.HandleFunc(path, f).Methods("GET")
//}
//
//func (a *AppDelivery) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	a.rt.HandleFunc(path, f).Methods("POST")
//}
//
//func (a *AppDelivery) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	a.rt.HandleFunc(path, f).Methods("PUT")
//}
//
//func (a *AppDelivery) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
//	a.rt.HandleFunc(path, f).Methods("DELETE")
//}
