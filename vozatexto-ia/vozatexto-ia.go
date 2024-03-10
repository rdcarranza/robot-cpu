package vozatexto_ia

import (
	"fmt"
	"log"
	"os"
)

func Vozatexto_ia() {
	fmt.Println("vozatexto-ia: Constructor")
	if !verificarEnv() {
		log.Fatal("Error: archivo env NO localizado!")
	}

	arch_txt := "./vozatexto-ia/ingresotexto.txt"
	arch_mp3 := "../ingresovoz.mp3"

	if !archExiste(arch_mp3) {
		log.Fatal("Error: archivo ingresovoz.mp3 NO localizado!")
	}

	if verificarArchTxt(arch_txt) {

	}

}

func VaT_GetModulo() {
	fmt.Println("módulo: vozatexto-ia")
}

func VaT_Estado() bool {
	return true
}

func verificarArchTxt(arch string) bool {

	if !archExiste(arch) {
		fmt.Println("El archivo ingresotexto.txt no existe!")
		err := crearArchTxt(arch)
		if err != nil {
			log.Fatal(err)
		}

	}

	if archExiste(arch) {
		return true
	}

	return false

}

func archExiste(arch string) bool {
	if _, err := os.Stat(arch); os.IsNotExist(err) {
		return false
	}
	return true
}

func crearArchTxt(dir_arch string) error {

	_, err := os.OpenFile(dir_arch, os.O_RDWR|os.O_CREATE, 0775)
	if err != nil {
		return err
		//log.Fatal(err)
	} else {
		fmt.Println("Se genera env de Vat exitosamente!")
	}

	return nil
}
