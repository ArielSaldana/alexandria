package controllers

import (
	"github.com/ArielSaldana/alexandria/internal/models"
	"github.com/ArielSaldana/alexandria/pkg/decoder"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/gographics/imagick.v3/imagick"
	"net/http"
)

func ImageCropController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)

	var op models.Operations
	var decoder decoder.QueryDecoder

	decoder.Decode(&op, r.URL.Query())

	handleCrop(mw, op)

	w.Write(mw.GetImageBlob())
}

func handleCrop(wand *imagick.MagickWand, settings models.Operations) {
	wand.CropImage(settings.Width, settings.Height, int(settings.X), int(settings.Y))
}
