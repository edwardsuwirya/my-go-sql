package deliveries

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"myfirstgosql/config"
	"net/http"
	"time"
)

const (
	LOGIN_ROUTE        = "/login"
	LOGOUT_ROUTE       = "/logout"
	HEALTH_CHECK_ROUTE = "/healthcheck"
)

type AuthRoute struct {
	prefix string
	rt     *mux.Router
	store  *sessions.CookieStore
}

func (pr *AuthRoute) InitRoute() {
	p := pr.rt.PathPrefix(pr.prefix).Subrouter()
	p.HandleFunc(LOGIN_ROUTE, pr.loginHandler).Methods(http.MethodPost)
	p.HandleFunc(LOGOUT_ROUTE, pr.logoutHandler).Methods(http.MethodGet)
	p.HandleFunc(HEALTH_CHECK_ROUTE, pr.healthCheckHandler).Methods(http.MethodGet)
}
func (pr *AuthRoute) loginHandler(w http.ResponseWriter, r *http.Request) {
	var users = map[string]string{"edo": "P@ssw0rd", "admin": "password"}
	session, _ := pr.store.Get(r, "session.id")
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Please pass the data as URL form encoded", http.StatusBadRequest)
		return
	}
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")
	if originalPassword, ok := users[username]; ok {
		if password == originalPassword {
			session.Values["authenticated"] = true
			session.Save(r, w)
		} else {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "User is not found", http.StatusNotFound)
		return
	}
	w.Write([]byte("Logged In successfully"))
}
func (pr *AuthRoute) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := pr.store.Get(r, "session.id")
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte(""))
}

func (pr *AuthRoute) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := pr.store.Get(r, "session.id")
	if (session.Values["authenticated"] != nil) && session.Values["authenticated"] != false {
		w.Write([]byte(time.Now().String()))
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}
func NewAuthRoute(prefix string, sf *config.SessionFactory, rt *mux.Router) IAppRouter {
	store := sessions.NewCookieStore([]byte("SESSION_SECRET"))
	return &AuthRoute{
		prefix,
		rt,
		store,
	}
}
