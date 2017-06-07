//Dave Cheney of go forum did helps me with this example
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		htmp := `<!doctype html>

    	<html lang="en">
	<head>
	  <meta charset="utf-8">
	  <title>MAIN</title>
	</head>

	<body>
	  <form action="/" enctype="multipart/form-data" method="POST">
		  	Photo: <input type="file"  id="filedata" name="filedata">
	 	  	<input type="submit"	value="Upload Photo">
	  </form>
	</body>
	</html>`
		switch req.Method {
		case "POST":
			fmt.Println(req.ContentLength)
			if req.ContentLength > 50*1024 {
				x := req.ContentLength - (50 * 1024)
				fmt.Printf("Bigger than allowed by %v bytes", x)
				http.Redirect(res, req, "/bad", http.StatusSeeOther)
				return
			}
			http.Redirect(res, req, "/ok", http.StatusSeeOther)
		case "GET":
			io.WriteString(res, htmp)
		default:
			http.Error(res, "method not allowed: "+req.Method, 405)
		}

	})
	http.HandleFunc("/bad", func(res http.ResponseWriter, req *http.Request) {
		htmp := `<!doctype html>

	<html lang="en">
	<head>
	  <meta charset="utf-8">
	  <title>BAD</title>
	</head>
	<body>
	  <h1>BAD</h1>
	</body>
	</html>`
		io.WriteString(res, htmp)
	})
	http.HandleFunc("/ok", func(res http.ResponseWriter, req *http.Request) {
		htmp := `<!doctype html>

	<html lang="en">
	<head>
	  <meta charset="utf-8">
	  <title>OK</title>
	</head>
	<body>
	  <h1>OK</h1>
	</body>
	</html>`
		io.WriteString(res, htmp)
	})

	http.ListenAndServe(":8080", nil)
}
