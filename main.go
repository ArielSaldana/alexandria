package main

import (
	"github.com/ArielSaldana/alexandria/internal/controllers"
	"github.com/codegangsta/negroni"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/gographics/imagick.v3/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mux := httprouter.New()
	mux.GET("/image", controllers.ImageController)
	mux.GET("/image/crop", controllers.ImageCropController)
	mux.GET("/image/resize", controllers.ImageResizeController)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}
