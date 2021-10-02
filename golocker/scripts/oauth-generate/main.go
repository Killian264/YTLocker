package main

func main() {
	scopes := []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/youtube.readonly", "https://www.googleapis.com/auth/youtube"}

	getClient(scopes)
}
