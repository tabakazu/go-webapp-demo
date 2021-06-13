package main

func main() {
	srv := InitializeServer()
	srv.ListenAndServe()
}
