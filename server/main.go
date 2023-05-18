package main

func main() {
	apiServer := NewAPIServer("8000")
	apiServer.Run()
}
