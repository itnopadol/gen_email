package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/gin-contrib/cors.v1"
	"github.com/itnopadol/gen_email/ctrl"
	)

//var (
//	templates *template.Template
//)

func main() {
	r := gin.New()
	r.Use(cors.Default())

	//r.GET("/email", ctrl.GenEmail)
	r.GET("/gentax", ctrl.GenTaxData)

	r.Run(":8099")
}