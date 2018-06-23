package main

import (
	"fmt"
	"net/http"
	"io"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Println("req", req)
	io.WriteString(
		res,
		`<DOCTYPE html>
		<html>
			<head>
				<title>Hello from go</title>
			</head>
			<body>
				<h1>Hello from go-lang!</h1>
			</body>
		</html>`,
	)
}

func main(){
	http.HandleFunc("/hello", hello)
	http.Handle(
		"/assets/",
		http.StripPrefix(
			"/assets/",
			http.FileServer(http.Dir("assets")),
		),
	)
	http.ListenAndServe(":9000", nil)
}