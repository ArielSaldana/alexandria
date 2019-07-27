package imageutils

import "math"

type ORIENTATION int

const (
	PORTRAIT  ORIENTATION = 1
	LANDSCAPE ORIENTATION = 2
	SQUARE    ORIENTATION = 3
)

func GetOrientation(width int, height int) ORIENTATION {
	if IsPortrait(width, height) {
		return PORTRAIT
	} else if IsLandScape(width, height) {
		return LANDSCAPE
	}
	return SQUARE
}

func IsPortrait(width int, height int) bool {
	if height > width {
		return true
	}
	return false
}

func IsLandScape(width int, height int) bool {
	if width > height {
		return true
	}
	return false
}

func IsSquare(width int, height int) bool {
	if width == height {
		return true
	}
	return false
}

func GetAspectRatio(width uint, height uint) float64 {
	aspectRatio := float64(width / height)
	return aspectRatio
}

func GetAdjustedWidth(width uint, height uint, adjustedHeight uint) uint {
	adjustedWidth := math.Ceil(float64((width / height) * adjustedHeight))
	return uint(adjustedWidth)
}

func GetAdjustedHeight(width uint, height uint, adjustedWidth uint) uint {
	adjustedHeight := math.Ceil(float64(adjustedWidth / (width / height)))
	return uint(adjustedHeight)
}

func Biggest(x uint, y uint) uint {
	if x > y {
		return x
	}
	return y
}

func Smallest(x uint, y uint) uint {
	if x > y {
		return y
	}
	return x
}

func GetPos(originalWidth, newWidth uint) uint {
	return (originalWidth - newWidth) / 2
}
