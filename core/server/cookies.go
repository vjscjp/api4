package server

import (
	"fmt"
	http "net/http"
	"time"
	
	"github.com/gorilla/securecookie"
)

const (
	COOKIE_NAME = "m_auth"
	HASH_KEY    = "52250f1f44b4460cfbbddd51f13193216c85c655de5722b05f5822ab6aa8bsdwe"
	BLOCK_KEY   = "9732eac67a6d40d173abf1be8bdc5e76"
)

var s = securecookie.New([]byte(HASH_KEY), []byte(BLOCK_KEY))

func SetCookie(w http.ResponseWriter, value map[string]string) {
	duration := time.Now().Add(time.Hour * 10) // 3 min
	setCookieWithName(w, COOKIE_NAME, duration, value)
	fmt.Println("------")
	fmt.Println(value)
	fmt.Println("------")
}

func ReadCookie(r *http.Request) (map[string]string, error) {
	return readCookieWithName(r, COOKIE_NAME)
}

func DeleteCookie(w http.ResponseWriter) {
	duration := time.Now().Add(-time.Minute * 1) // 1 min ago
	setCookieWithName(w, COOKIE_NAME, duration, nil)
}

func setCookieWithName(w http.ResponseWriter, cookieName string, duration time.Time, value map[string]string) {
	if cookie := buildCookie(cookieName, duration, value); cookie != nil {
		fmt.Println(cookie)
		//w.Header().Add("X-Cookie",cookie.String())
		//This function seems doesn't set Cookie in Header
		http.SetCookie(w, cookie)
	}
}

func readCookieWithName(r *http.Request, cookieName string) (map[string]string, error) {
	value := make(map[string]string)
	if cookie, err := r.Cookie(cookieName); err == nil {
		err = s.Decode(cookieName, cookie.Value, &value)
		return value, err
	} else {
		return value, err
	}
}

func buildCookie(cookieName string, duration time.Time, value map[string]string) *http.Cookie {
	if encoded, err := GetEncodedVal(cookieName, value); err == nil {
		return &http.Cookie{
			Name:    cookieName,
			Value:   encoded,
			Path:    "/",
			Expires: duration,
		}
	} else {
		fmt.Println(err)
		return nil
	}
}


func GetEncodedVal(cookieName string, value map[string]string) (string, error) {
	return s.Encode(cookieName, value)
}


func GetDecodeVal(cookieName string, val string) (map[string]string, error) {
	value := make(map[string]string)
	err := s.Decode(cookieName, val, &value)
	return value, err
}
