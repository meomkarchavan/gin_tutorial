package main

import "os"

func main() {
	os.Setenv("ACCESS_SECRET", "omkar")

	r := registerRoutes()
	r.Run(":8080")

}
