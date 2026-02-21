package main

import (
	"automation-service/internal/checker"
	"fmt"
)

func main() {
	result := checker.CheckURL("http://localhost:8081/ok")
	fmt.Println(result)
}
