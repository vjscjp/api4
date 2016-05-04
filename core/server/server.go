package server

import (
	"net/http"

	"github.com/vjscjp/api4/core/controllers"
	"github.com/vjscjp/api4/core/controllers/app"
	"github.com/vjscjp/api4/core/controllers/host"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	Default    = "/"
	Apps       = "/apps/{id}"
	ListApps   = "/apps"
	HostPorts  = "/hostport"
	LoginPath  = "/login"
	LogoutPath = "/logout"
	HostPort   = "/hostport/{id}/{port}"
)

func InitRoutes() *negroni.Negroni {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc(Default, Status).Methods("GET")
	muxRouter.HandleFunc(LoginPath, LogIn).Methods("POST")
	muxRouter.HandleFunc(LogoutPath, LogOut).Methods("GET")
	muxRouter.HandleFunc(Apps, app.GetApp).Methods("GET")
	muxRouter.HandleFunc(ListApps, app.ListApps).Methods("GET")
	muxRouter.HandleFunc(HostPorts, host.GetHostPorts).Methods("GET")
	muxRouter.HandleFunc(HostPort, host.GetHostPort).Methods("GET")
	
	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "CONNECT"},
		AllowedHeaders:   []string{"*"},
	})
	
	n := negroni.New(
		negroni.NewStatic(http.Dir("public")),
		cor,
		&CookieAuth{},
	)
	n.UseHandler(muxRouter)
	return n
}

func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	controllers.ServeJsonResponseWithCode(w, map[string]string{"Status": "OK"}, http.StatusOK)
}
