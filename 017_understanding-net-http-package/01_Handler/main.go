package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServceHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Any code you want in this func")
}

func main() {

}
