package main

import "blurbbles/internal"

// just for running locally
func main() {
	app := internal.Server()
	app.Listen(":3000")
}
