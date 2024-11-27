package utils

func ExisteElemento(slice []string, item string) int {

	var ubicacion int = -1

	for i, v := range slice {

		if v == item {
			ubicacion = i
			break
		}

	}

	return ubicacion

}
