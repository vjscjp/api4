package app

import (
	"fmt"
	"net/http"

	"github.com/vjscjp/api4/core"
	"github.com/vjscjp/api4/core/controllers"
	"github.com/gorilla/mux"
)

func GetApp(w http.ResponseWriter, r *http.Request) {
	
	fmt.Printf("%+v", r)
	 
	client, err := controllers.NewMarathonClient(w, r)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	status := http.StatusOK
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("Input Id", id)
	if len(id) < 1 {
		http.Error(w, "Input param id is missing", http.StatusInternalServerError)
		return
	}
	resp, err := client.GetApp(id)
	if err != nil {
		fmt.Println("Invalid Details : " + err.Error())
		status = http.StatusServiceUnavailable
	}

	if response, er := controllers.ParseResponse(resp); er != nil {
		http.Error(w, fmt.Sprintf("App Error: %s for id %s", err.Error(), id), http.StatusInternalServerError)
	} else {
		controllers.ServeJsonResponseWithCode(w, response, status)
	}

}

func ListApps(w http.ResponseWriter, r *http.Request) {

	client, err := controllers.NewMarathonClient(w, r)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	status := http.StatusOK

	resp, err := client.ListApps()
	if err != nil {
		fmt.Println("Invalid Details : " + err.Error())
		status = http.StatusServiceUnavailable
	}
	var apps []core.ListApp
	for _, a := range resp.Apps {
		apps = append(apps, core.ListApp{Id: a.ID, Display: a.ID})
	}
	controllers.ServeJsonResponseWithCode(w, apps, status)
}
