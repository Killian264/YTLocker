package main

func main() {
	scope := "https://www.googleapis.com/auth/youtube"

	filePath := "../../secrets/"

	getClient(scope, filePath)
}
