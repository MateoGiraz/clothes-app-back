package startup

import (
	"github.com/matoegiraz/clothes-app/models"
)

func CreateColors() []models.AvailableColorType {
	stringColors := []string{
		"blue",
		"red",
		"green",
		"black",
		"white",
	}

	var modelColors []models.AvailableColorType

	for _, color := range stringColors {
		newColor := models.AvailableColorType{color}
		modelColors = append(modelColors, newColor)
	}

	return modelColors
}

func CreateClothes() []models.AvailableClothesType {
	stringClothes := []string{
		"t-shirts",
		"shirts",
		"hoodie",
		"jacket",
		"jeans",
		"pants",
		"shorts",
		"shoes",
		"hats",
	}

	var modelClothes []models.AvailableClothesType

	for _, name := range stringClothes {
		newName := models.AvailableClothesType{name}
		modelClothes = append(modelClothes, newName)
	}

	return modelClothes
}
