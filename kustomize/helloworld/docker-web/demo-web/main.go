package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	risky := os.Getenv("ENABLE_RISKY")
	var alt string
	if risky == "true" {
		alt = os.Getenv("ALT_GREEITMG")
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello astaxie: %s", alt)) // 这个写入到 w 的是输出到客户端的
}

func main() {
	http.HandleFunc("/", sayhelloName)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
