package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("index.gohtml")
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			file, _, err := req.FormFile("my-file")
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer file.Close()

			src := io.LimitReader(file, 400)

			dst, err := os.Create(filepath.Join(".", "file.txt"))
			if err != nil {
				http.Error(res, err.Error(), 500)
				return
			}
			defer dst.Close()

			io.Copy(dst, src)
		}
		tpl.Execute(res, nil)
		file, _ := ioutil.ReadFile("file.txt")
		fmt.Fprint(res, strings.Split(string(file), "\n"))

		// 	res.Header().Set("Content-Type", "text/html")
		// 	io.WriteString(res, `
		//     <form method="POST" enctype="multipart/form-data">
		//       <input type="file" name="my-file">
		//       <input type="submit">
		//     </form>
		//     `)
	}))
}
