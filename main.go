package main

func main() {
	r := setupRouter()
	r.Run("localhost:8080")
}
