package main

var counter int64

func init() {
	counter = set_counter()
}
func main() {

	r := registerRoutes()
	r.Run(":8080")

}
