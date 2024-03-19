package vozatexto_ia

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func Vozatexto_ia() {
	fmt.Println("vozatexto-ia: Constructor")
	if !verificarEnv() {
		log.Fatal("Error: archivo env NO localizado!")
	}

	arch_txt := "./vozatexto-ia/ingresotexto.txt"
	arch_mp3 := "./ingresovoz.mp3"

	if !archExiste(arch_mp3) {
		log.Fatal("Error: archivo ingresovoz.mp3 NO localizado!")
	}

	if !verificarArchTxt(arch_txt) {
		log.Fatal("Error: archivo ingresotexto.txt NO localizado!")
	}

	err := convertirAaT(arch_mp3, arch_txt)
	if err != nil {
		log.Fatal(err)
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
	//dir_arch, _ := filepath.Abs(arch)
	//fmt.Println("La ruta absoluta del archivo es: " + dir_arch)
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
		fmt.Println("Se genera archivo: " + dir_arch + " - exitosamente!")
	}

	return nil
}

func convertirAaT(arch_audio string, arch_txt string) error {

	url := "https://api.openai.com/v1/audio/transcriptions"

	// Abrir el archivo de audio
	file, err := os.Open(arch_audio)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	// Crear un formulario multipart
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", arch_audio)
	if err != nil {
		fmt.Println(err)
		return err
	}

	io.Copy(part, file)

	// Agregar otros campos
	_ = writer.WriteField("model", "whisper-1")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	_ = writer.WriteField("response_format", "text")
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Crear la solicitud
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	//req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_API_KEY"))
	env, err := getEnv("OPENAI_API_KEY")
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Set("Authorization", "Bearer "+env)

	//fmt.Println("Petición a API:")
	//fmt.Println(req)

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	// Leer la respuesta
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(respBody))

	// Escribir la respuesta en un archivo
	err = os.WriteFile(arch_txt, respBody, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("La respuesta se ha guardado en ingreso.txt")

	return nil
}
