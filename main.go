package main

func main() {
	db, dbClose := InitializeDB()
	defer dbClose()

	srv := InitializeServer(db)
	srv.ListenAndServe()
}
