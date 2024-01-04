package vozatexto_ia

import (
	"fmt"
	"io"
	"log"
	"os"
)

func verificarEnv() bool {
	/*
		ex, _ := os.Executable()
		fmt.Println("La ruta del ejecutable es: " + ex)
		exPath := filepath.Dir(ex)
		fmt.Println("El directorio del ejecutable es:" + exPath)
	*/

	env := "./vozatexto-ia/vat.env"
	if !envExiste(env) {
		fmt.Println("El archivo env de VaT no existe!")
		copia_env, err := os.Open("./vozatexto-ia/env.copia")
		if err != nil {
			log.Fatal(err)
		}
		_env, err := os.OpenFile(env, os.O_RDWR|os.O_CREATE, 0775)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(_env, copia_env)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("Se genera env de Vat exitosamente!")
		}

	}

	if envExiste(env) {
		return true
	} else {
		return false
	}
}

func envExiste(arch_env string) bool {
	if _, err := os.Stat(arch_env); os.IsNotExist(err) {
		return false
	}
	return true

}
