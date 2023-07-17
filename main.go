package main

import (
	"github.com/Arch-4ng3l/TaskManage/api"
	"github.com/Arch-4ng3l/TaskManage/storage"
)

func main() {
	psql := storage.NewPostgres()
	psql.Init()
	api.NewAPIServer(":3000", psql).Run()
}
