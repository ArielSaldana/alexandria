package controllers

import (
	"github.com/ArielSaldana/alexandria/pkg/imageutils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/gographics/imagick.v3/imagick"
	"net/http"
)

func ImageController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)
	w.Write(mw.GetImageBlob())
}

func ImageCropController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)

	mw.CropImage(10, 10, 10, 10)
	w.Write(mw.GetImageBlob())
}

func ImageResizeController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)

	var width = uint(400)
	println(mw.GetImageWidth())
	println(mw.GetImageHeight())
	println(width)

	adjustedHeight := imageutils.GetAdjustedHeight(mw.GetImageWidth(), mw.GetImageHeight(), width)
	mw.ResizeImage(uint(width), uint(adjustedHeight), 0)
	w.Write(mw.GetImageBlob())
}

func getMagickWandUsingQueryParam(r *http.Request) *imagick.MagickWand {
	queryValues := r.URL.Query()
	url := queryValues.Get("image")

	mw := imagick.NewMagickWand()

	if err := mw.ReadImage(url); err != nil {
		panic(err)
	}

	return mw
}
