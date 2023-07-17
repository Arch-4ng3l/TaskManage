package main

func main() {

	psql := NewPostgres()
	psql.Init()
	NewAPIServer(":3000", psql).Run()
}
