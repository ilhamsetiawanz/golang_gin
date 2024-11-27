package main

import (
	Bukucontroller "golang_gin/controller/BukuController"
	"golang_gin/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	database.ConnectDb()
	r.GET("/get-books", Bukucontroller.Read)
	r.GET("/get-books/:id", Bukucontroller.ReadById)
	r.POST("/create-book", Bukucontroller.Create)
	r.PUT("/update-book/:id", Bukucontroller.Update)
	r.DELETE("/delete-book/:id", Bukucontroller.Delete)

	r.Run()
}
