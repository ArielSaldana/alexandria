package controllers

import (
	"github.com/ArielSaldana/alexandria/internal/models"
	"github.com/ArielSaldana/alexandria/pkg/imageutils"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/gographics/imagick.v3/imagick"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ImageController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)
	w.Write(mw.GetImageBlob())
}

func ImageCropController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)

	var keys []string
	keys = make([]string, 4)
	keys[0] = "width"
	keys[1] = "height"
	keys[2] = "x"
	keys[3] = "y"

	operation := new(models.Operations)
	queryParamsToOperations(r, operation, keys)

	// logic
	if operation.Width == nil {
		*operation.Width = mw.GetImageWidth()
	}

	if operation.Height == nil {
		*operation.Height = mw.GetImageHeight()
	}

	if operation.X == nil {
		*operation.X = 0
	}

	if operation.Y == nil {
		*operation.Y = 0
	}

	mw.CropImage(*operation.Width, *operation.Height, int(*operation.X), int(*operation.Y))
	w.Write(mw.GetImageBlob())
}

func ImageResizeController(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mw := getMagickWandUsingQueryParam(r)

	var width = uint(400)

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

func queryParamsToOperations(r *http.Request, operation *models.Operations, keys []string) models.Operations {
	queryValues := r.URL.Query()

	for _, key := range keys {
		var value = queryValues.Get(key)
		var capitalizedKey = strings.Title(key)

		if len(value) == 0 {
			continue
		}

		numericValue64, err := strconv.ParseUint(value, 10, 32)
		numericValue := uint(numericValue64)

		if err != nil {
			panic(err)
		}

		dynamicVariable := reflect.ValueOf(operation).Elem().FieldByName(capitalizedKey)

		if !dynamicVariable.IsValid() || !dynamicVariable.CanSet() {
			continue
		}

		dynamicVariable.Set(ptr(reflect.ValueOf(numericValue)))
	}

	return models.Operations{}
}

func ptr(v reflect.Value) reflect.Value {
	pt := reflect.PtrTo(v.Type()) // create a *T type.
	pv := reflect.New(pt.Elem())  // create a reflect.Value of type *T.
	pv.Elem().Set(v)              // sets pv to point to underlying value of v.
	return pv
}
