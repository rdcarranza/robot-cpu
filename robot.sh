#!/bin/bash

#Este es el hilo principal de ejecución del programa.

#GRABACIÓN de voz. BASH.
sh ./grabar.sh

#   si se creó el archivo ingresovoz.mp3 continua, caso contrario se interrumpe.
arch1=./ingresovoz.mp3

if [-f "$arch1"]
then
echo "ingresovoz.mp3 se generó correctamente!"
else
exit ;
echo "Error en la captura de audio."
fi

#VOZATEXTO-IA. GO.
#Se convierte ingresovoz.mp3 en texto plano, usando una API AI.

arch2=./vozatexto-ia/ingresotexto.txt

#   si se creó el archivo ingresotexto.txt continua, caso contrario se interrumpe.

if [-f "$arch2"]
then
echo "ingresotexto.txt se generó correctamente!"
else
exit ;
echo "Error en la generación de texto."
fi

#CONSULTA-IA. GO.
#Se realiza la consulta del texto ingresotexto.txt a una API AI y se almacena el resultado en resultado.txt



#LECTURA. BASH.
sh ./leer.sh
#Se realiza la lectura del texto del archivo resultado.txt


#   siempre se elimina ingresovoz.mp3 antes de finalizar, si es que existe.
rm -f arch1
