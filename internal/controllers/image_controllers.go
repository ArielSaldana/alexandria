/*
 * TODO: move the reflection code over to a utility
 */

package controllers

import (
	"github.com/ArielSaldana/alexandria/internal/models"
	"github.com/ArielSaldana/alexandria/pkg/decoder"
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
	var op models.Operations
	var decoder decoder.QueryDecoder
	decoder.Decode(&op, r.URL.Query())

	mw.CropImage(op.Width, op.Height, int(op.X), int(op.Y))
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
		var queryValue = queryValues.Get(key)

		if len(queryValue) == 0 {
			continue
		}

		var capitalizedKey = strings.Title(key)
		var value interface{}

		dynamicVariable := reflect.ValueOf(operation).Elem().FieldByName(capitalizedKey)

		if !dynamicVariable.IsValid() || !dynamicVariable.CanSet() {
			continue
		}

		elementType := reflect.PtrTo(dynamicVariable.Type())
		valueOfType := elementType.Elem().Elem().Kind().String()

		switch valueOfType {
		case "uint":
			numericValue64, err := strconv.ParseUint(queryValue, 10, 32)
			if err != nil {
				panic(err)
			}
			value = uint(numericValue64)
			break

		case "string":
			value = queryValue
			break

		default:
			continue
		}

		dynamicVariable.Set(ptr(reflect.ValueOf(value)))
	}

	return models.Operations{}
}

func ptr(v reflect.Value) reflect.Value {
	pt := reflect.PtrTo(v.Type()) // create a *T type.
	pv := reflect.New(pt.Elem())  // create a reflect.Value of type *T.
	pv.Elem().Set(v)              // sets pv to point to underlying value of v.
	return pv
}
