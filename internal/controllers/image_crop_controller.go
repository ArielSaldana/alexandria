package controllers

import (
	"github.com/ArielSaldana/alexandria/internal/models"
	"github.com/ArielSaldana/alexandria/pkg/decoder"
	"github.com/ArielSaldana/alexandria/pkg/imageutils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/gographics/imagick.v3/imagick"
	"net/http"
)

func ImageCropController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)

	var imageConfig models.Operations
	var decoder decoder.QueryDecoder

	decoder.Decode(&imageConfig, r.URL.Query())

	switch imageConfig.Name {
	case "square":
		handleCropSquare(mw, imageConfig)
	default:
		handleCrop(mw, imageConfig)
	}

	w.Write(mw.GetImageBlob())
}

func handleCrop(wand *imagick.MagickWand, settings models.Operations) {
	var err = wand.CropImage(
		settings.Width,
		settings.Height,
		int(settings.X),
		int(settings.Y),
	)

	if err != nil {
		// handle err
		panic(err)
	}
}

func handleCropSquare(wand *imagick.MagickWand, imageConfig models.Operations) {
	var width uint

	width = imageutils.Smallest(wand.GetImageWidth(), wand.GetImageHeight())
	if imageConfig.Width != 0 && imageConfig.Width < width {
		width = imageConfig.Width
	}
	if imageConfig.Height != 0 && imageConfig.Height < width {
		width = imageConfig.Height
	}

	// width start position
	posX := imageutils.GetPos(wand.GetImageWidth(), width)
	// height start position
	posY := imageutils.GetPos(wand.GetImageHeight(), width)

	wand.CropImage(width, width, int(posX), int(posY))
}
