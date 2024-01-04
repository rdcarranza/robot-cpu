package vozatexto_ia

import (
	"fmt"
	"log"
)

func Vozatexto_ia() {
	fmt.Println("vozatexto-ia: Constructor")
	if !verificarEnv() {
		log.Fatal("Error: archivo env NO localizado!")
	}
}

func VaT_GetModulo() {
	fmt.Println("módulo: vozatexto-ia")
}

func VaT_Estado() bool {
	return true
}
