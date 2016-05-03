package host

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vjscjp/api4/core/controllers"
)

type HOST_PORT struct {
	Host string
	Port int
}

func GetHostPort(w http.ResponseWriter, r *http.Request) {
	client, err := controllers.NewMarathonClient(w, r)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	status := http.StatusOK
	vars := mux.Vars(r)
	id := vars["id"]
	var port int
	if p := vars["port"]; len(p) < 1 {
		http.Error(w, "Input param Port is missing", http.StatusInternalServerError)
		return
	} else {
		if port, err = strconv.Atoi(p); err != nil {
			http.Error(w, "Invalid Port input param, required a valid interger port number", http.StatusInternalServerError)
			return
		}
	}
	resp, err := client.ListTasks()
	if err != nil {
		status = http.StatusServiceUnavailable
	}

	if resp == nil {
		http.Error(w, fmt.Sprintf("Host Port Error: Found Nil Response."), http.StatusInternalServerError)
		return
	}

	if resp.Tasks == nil {
		http.Error(w, fmt.Sprintf("Host Port Error: Found Nil Response.Tasks"), http.StatusInternalServerError)
		return
	}
	var flag bool
	for _, t := range resp.Tasks {
		if t.Host == id {
			for _, p := range t.Ports {
				//fmt.Println("DEBUG port ", p)
				if p == port {
					resp, err := client.GetApp(t.AppID)
					if err != nil {
						http.Error(w, "Hostport Error:"+err.Error(), http.StatusInternalServerError)
						return
					}
					if response, er := controllers.ParseResponse(resp); er != nil {
						http.Error(w, fmt.Sprintf("App Error: %s for id %s", err.Error(), id), http.StatusInternalServerError)
						return
					} else {
						controllers.ServeJsonResponseWithCode(w, response, status)
						return
					}
					flag = true
				}
			}
		}
	}
	if !flag {
		http.Error(w, "No Record Found", http.StatusInternalServerError)
		return
	}

}

func GetHostPorts(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotFound)
	return
}
