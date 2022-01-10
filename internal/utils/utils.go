package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
)

func GetVar(req *http.Request) {
	vars := mux.Vars(req)
	fmt.Print(vars)
}

func GetID(req *http.Request) int {
	vars := mux.Vars(req)
	if id, ok := vars["ID"]; ok {
		id, _ := strconv.ParseInt(id, 10, 32)
		return int(id)
	}
	return -1
}

func GetCookie(req *http.Request, name string) string {
	cookie, err := req.Cookie(name)
	if err != nil {
		log.Fatal(err)
	}
	return cookie.Value
}

var apiHost = "http://192.168.192.50:8080/api/"

func MakeGETRequest(path string, token string) []byte {
	reqURL, err := url.Parse(apiHost + path)
	if err != nil {
		log.Fatal(err)
	}
	reqURL.Query().Add("token", token)

	res, err := http.Get(reqURL.String())
	if err != nil {
		log.Fatal(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return resBody
}
