package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/nu7hatch/gouuid"
)

type User struct {
	Name string
	Age  string
}

func serve(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	name := req.FormValue("Name:")
	age := req.FormValue("Age:")
	data := datF(name, age)
	code := getCode(data)
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session-fino",
			Value: id.String() + "|" + name + age + "|" + data + "|" + code,
			// Secure: true,
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	err = tpl.Execute(res, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func datF(name string, age string) string {
	user := User{
		Name: name,
		Age:  age,
	}
	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	// info := base64.URLEncoding.EncodeToString(bs)
	return base64.URLEncoding.EncodeToString(bs)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	http.HandleFunc("/", serve)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
