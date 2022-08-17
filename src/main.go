package main

import (
	"os"

	"twodo.app/condo/service"
	"twodo.app/condo/utility"
)

func main() {
	service.DB.Migrate()
	os.Mkdir(utility.AttachmentPath, os.ModePerm)
	router := service.Router.Create()
	router.Run(":8080")
}
