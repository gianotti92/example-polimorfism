package main

import (
	"awesomeProject/router"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Iniciando web server")

	err := router.StartDefaultEngine()

	if err != nil {
		os.Exit(-1)
	}

}
