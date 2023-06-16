package main

import (
	"fmt"

	"github.com/njfamirm/sitatus/pkg/uptime"
)

func main() {
	responseTime, err := uptime.Uptime("https://www.google.com")

	if err != nil {
		fmt.Println("Error:", err)
		panic(err)
	}

	fmt.Println("Response Time:", responseTime)
}
