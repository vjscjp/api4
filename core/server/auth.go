package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	gcontext "github.com/gorilla/context"
	"github.com/vjscjp/api4/core"
	"github.com/vjscjp/api4/core/controllers"
)

const (
	CONTEXT_USER     = "user"
	CONTEXT_USER_KEY = "pass"
	CONTEXT_API      = "api"
)

func GetUser(r *http.Request) *core.User {
	if rv := gcontext.Get(r, CONTEXT_USER); rv != nil {
		if user, ok := rv.(*core.User); ok {
			return user
		}
	}
	return nil
}

func SetUser(r *http.Request, val *core.User) {
	gcontext.Set(r, CONTEXT_USER, val)
}

func UnsetUser(r *http.Request) {
	gcontext.Delete(r, CONTEXT_USER)
}

type Login struct {
	StatusCode int
	Status     string
	Apps       interface{}
	Token      string
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	cookie, _ := ReadCookie(r)
	b, e := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	status := http.StatusOK
	result := new(Login)
	if e != nil {
		result.Status = "Login Payload payload. " + e.Error()
		controllers.ServeJsonResponseWithCode(w, result, status)
		return
	}
	var user core.User
	if e = json.Unmarshal(b, &user); e != nil {
		result.Status = "Not a Valid JSON payload. " + e.Error()
		controllers.ServeJsonResponseWithCode(w, result, status)
		return
	}

	if len(user.ID) < 1 || len(user.Pass) < 1 || len(user.Api) < 1 {
		result.Status = "Either one of the parameter is missing (id, pass or api)"
		controllers.ServeJsonResponseWithCode(w, result, status)
		return
	}

	cookie[CONTEXT_USER] = user.ID
	cookie[CONTEXT_USER_KEY] = user.Pass
	cookie[CONTEXT_API] = user.Api

	SetCookie(w, cookie)
	SetUser(r, &user)

	//Check Marathon Api call to validate User credentials
	client, err := controllers.NewMarathonClient(w, r)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}
	//Little slow we may need to look for other
	resp, err := client.ListApps()
	if err != nil {
		controllers.ServeJsonResponseWithCode(w, &Login{401, "Invalid Credentials", nil, ""}, status)
		return
	}
	var apps []core.ListApp
	result.StatusCode = 200
	result.Status = "OK"
	for _, a := range resp.Apps {
		apps = append(apps, core.ListApp{Id: a.ID, Display: a.ID})
	}
	result.Apps = apps
	//temp solution as UI is unable to read Cookies
	if c, e := GetEncodedVal(COOKIE_NAME, cookie); e == nil {
		result.Token = fmt.Sprintf("%s=%s;", COOKIE_NAME, c)
	}

	controllers.ServeJsonResponseWithCode(w, result, status)

}

func LogOut(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w)
	UnsetUser(r)
}

func IsLoggedIn(r *http.Request) bool {
	return GetUser(r) != nil
}

// The rest of this is auth related middlewear that sets up the above helper functions on the request

type CookieAuth struct {
}

// Negroni compatible interface
func (cookieAuth *CookieAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	cookie, err := ReadCookie(r)
	if err != nil {
		fmt.Println(err)
	}
	var currentUser core.User
	//fmt.Println("COOKIES Pass: ", cookie[CONTEXT_USER_KEY])
	fmt.Println("COOKIES User: ", cookie[CONTEXT_USER])
	currentUser.ID = cookie[CONTEXT_USER]
	currentUser.Pass = cookie[CONTEXT_USER_KEY]
	currentUser.Api = cookie[CONTEXT_API]
	gcontext.Set(r, CONTEXT_USER, &currentUser)
	fmt.Println("CURRENT USER: ", currentUser)
	next(w, r)
}
