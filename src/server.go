package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	age := 23
	fmt.Println(age)

	router := mux.NewRouter()
}
