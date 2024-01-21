package main

import (
	"fmt"
	"net/http"

	"goProject/dictionary"
	"goProject/route"
)

const filePath = "dictionary.txt"

func main() {
	d := dictionary.New(filePath)
	defer d.Close()

	router := route.NewRouter(d)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", router)
}
