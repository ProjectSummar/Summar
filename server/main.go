package main

func main() {
	apiServer := NewAPIServer("3001")
	apiServer.Run()
}
