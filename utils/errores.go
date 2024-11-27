package utils

import (
	"fmt"
	"os"
)

func MiError(descripcion string, err error, cancelar bool) {
	if err != nil {
		fmt.Printf("Error: %s - %v\n", descripcion, err)
		if cancelar {
			os.Exit(1)
		}
	}
}
