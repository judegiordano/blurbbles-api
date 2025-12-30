package main

import "blurble.com/internal"

// just for running locally
func main() {
	app := internal.Server()
	app.Listen(":3000")
}
