package main

func main() {
	r := registerRoutes()
	r.Run(":8080")
}
