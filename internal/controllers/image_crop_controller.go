/*
 * TODO decide on either cropping the image or clamping the position /
 * to pos <= fullWdith - width
 */

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
	if imageConfig.Width == 0 || imageConfig.Width > wand.GetImageWidth() {
		imageConfig.Width = wand.GetImageWidth()
	}

	if imageConfig.Height == 0 || imageConfig.Height > wand.GetImageHeight() {
		imageConfig.Height = wand.GetImageHeight()
	}

	imageConfig.Width = imageutils.Smallest(imageConfig.Width, imageConfig.Height)
	imageConfig.Height = imageConfig.Width

	// width start position
	posX := imageutils.GetPos(wand.GetImageWidth(), imageConfig.Width)
	// height start position
	posY := imageutils.GetPos(wand.GetImageHeight(), imageConfig.Width)

	wand.CropImage(imageConfig.Width, imageConfig.Width, int(posX), int(posY))
}
