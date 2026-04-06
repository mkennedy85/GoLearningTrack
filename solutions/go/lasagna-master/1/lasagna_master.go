package lasagnamaster

func PreparationTime(layers []string, timePerLayer int) int {
    if timePerLayer == 0 {
        return len(layers) * 2
    }
    return len(layers) * timePerLayer
}

func Quantities(layers []string) (noodleQuantity int, sauceQuantity float64) {
    for _, layer := range layers {
        switch layer {
            case "noodles":
            	noodleQuantity += 50
            case "sauce":
            	sauceQuantity += 0.2
        }
    }
    return
}

func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(perPortion []float64, portions int) []float64 {
    var scaledQuantities []float64
    for i := 0; i < len(perPortion); i++ {
        newPortion := perPortion[i] * (float64(portions) / 2)
        scaledQuantities = append(scaledQuantities, newPortion)
    }
    return scaledQuantities
}
