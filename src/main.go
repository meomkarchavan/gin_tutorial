package main

func main() {

	r := registerRoutes()
	r.Run(":8080")

	// result, err := find_user(22)
	// if err != nil {
	// 	fmt.Println("NOt found")
	// } else {
	// 	fmt.Printf("Found a single document: %+v\n", result)
	// }
}
