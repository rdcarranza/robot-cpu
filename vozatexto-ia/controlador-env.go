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
	copia_env := "./vozatexto-ia/env.copia"
	if !envExiste(env) {
		fmt.Println("El archivo env de VaT no existe!")
		err := crearEnv(env, copia_env)
		if err != nil {
			log.Fatal(err)
		}

	}

	if envExiste(env) {
		return true
	}

	return false

}

func crearEnv(dir_env string, c_dir_env string) error {
	copia_env, err := os.Open(c_dir_env)
	if err != nil {
		log.Fatal(err)
	}
	_env, err := os.OpenFile(dir_env, os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		return err
		//log.Fatal(err)
	}
	_, err = io.Copy(_env, copia_env)
	if err != nil {
		return err
		//log.Fatal(err)

	} else {
		fmt.Println("Se genera env de Vat exitosamente!")
	}

	return nil
}

func envExiste(arch_env string) bool {
	if _, err := os.Stat(arch_env); os.IsNotExist(err) {
		return false
	}
	return true

}
