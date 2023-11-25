#!/bin/bash

#Este es el hilo principal de ejecución del programa.

#GRABACIÓN de voz. BASH.
sh ./grabar.sh

#   si se creó el archivo ingresovoz.mp3 continua, caso contrario se interrumpe.


#VOZATEXTO-IA. GO.
#Se convierte ingresovoz.mp3 en texto plano, usando una API AI.

#   si se creó el archivo ingresotexto.txt continua, caso contrario se interrumpe.

#CONSULTA-IA. GO.
#Se realiza la consulta del texto ingresotexto.txt a una API AI y se almacena el resultado en resultado.txt



#LECTURA. BASH.
sh ./leer.sh
#Se realiza la lectura del texto del archivo resultado.txt


#   siempre se elimina ingresovoz.mp3 antes de finalizar, si es que existe.
