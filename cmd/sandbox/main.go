package main

import "github.com/filipgorny/workshop/internal/application/console"

func main() {
	application := console.NewConsoleApplication()

	application.Run()
}
