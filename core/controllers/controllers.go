package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	f "github.com/MustWin/gomarathon"
	gcontext "github.com/gorilla/context"
	"github.com/vjscjp/api4/core"
)

func NewMarathonClient(w http.ResponseWriter, r *http.Request) (client *f.Client, err error) {

	user := GetCurrentUser(r)
	if user == nil {
		ServeJsonResponseWithCode(w, map[string]interface{}{"StatusCode": 401, "Status": "User not Logged in"}, http.StatusOK)
		return nil, fmt.Errorf("User Not Logged in.")
	}
	auth := f.HttpBasicAuth{user.ID, user.Pass}
	client, err = f.NewClient(user.Api, &auth, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		fmt.Println("Invalid Details : " + err.Error())
	}
	return
}

func GetCurrentUser(r *http.Request) *core.User {
	if rv := gcontext.Get(r, "user"); rv != nil {
		if user, ok := rv.(*core.User); ok {
			return user
		}
	}
	return nil
}
func ServeJsonResponseWithCode(w http.ResponseWriter, responseBodyObj interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(responseBodyObj); err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func ParseResponse(resp *f.Response) (core.App, error) {
	var out core.App
	fmt.Println("Error ", resp)
	if resp == nil {
		return out, fmt.Errorf("Found Empty Response")
	}
	if resp.App == nil {
		return out, fmt.Errorf("No App is found")
	}

	out.Id = resp.App.ID
	if len(resp.App.Labels) > 0 {
		if pn, k := resp.App.Labels["project_name"]; k {
			out.ProjectName = pn
		}
		if pi, k := resp.App.Labels["project_id"]; k {
			out.ProjectId = pi
		}
		if en, k := resp.App.Labels["env_name"]; k {
			out.EnvName = en
		}
		if ei, k := resp.App.Labels["env_id"]; k {
			out.EnvId = ei
		}

		if sn, k := resp.App.Labels["service_name"]; k {
			out.ServiceName = sn
		}
		if si, k := resp.App.Labels["service_id"]; k {
			out.ServiceId = si
		}
	}

	if resp.App.Tasks != nil {
		var tasks []*core.Task
		for _, tsks := range resp.App.Tasks {
			task := new(core.Task)
			task.Host = tsks.Host
			task.Ports = tsks.Ports
			task.Id = tsks.ID
			task.AppId = tsks.AppID

			tasks = append(tasks, task)
		}
		if len(tasks) > 0 {
			out.Tasks = tasks
		}
	}
	return out, nil
}
