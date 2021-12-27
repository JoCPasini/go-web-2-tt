package main

import (
	"fmt"
	"os"

	"github.com/JosePasiniMercadolibre/go-web-2-tt/cmd/server/handler"
	"github.com/JosePasiniMercadolibre/go-web-2-tt/internal/transacciones"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//_ = godotenv.Load()

	err := godotenv.Load()
	fmt.Println(os.Getenv("TOKEN"))
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("************************", os.Getenv("TOKEN"))
	}

	repo := transacciones.NewRepository()
	service := transacciones.NewService(repo)
	handler := handler.NewTransaccion(service)

	r := gin.Default()
	pr := r.Group("/transacciones")
	{
		pr.GET("/getAll", handler.GetAll())
		pr.POST("/store", handler.Store())
		pr.PUT("/update/:id", handler.Update())
		pr.PATCH("/updateCodigoMonto/:id", handler.UpdateCodigoMonto())
		pr.DELETE("/delete/:id", handler.Delete())
	}
	r.Run()
}
