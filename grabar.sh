#!/bin/bash

#Este script solo graba la interacción generando un archivo, hasta que se presiona una tecla.
#La validación sobre la ejecución de este script se hace sobre si se generó o no el archivo ingresovoz.mp3.
#dependencias: arecord, lame.

source ./robot.conf


arecord --device=$disp_grab -d 10 -f S16_LE -c 2 -r 44100 -t wav ./ingresovoz.wav &
#arecord --device=$dispositivo -d 10 -f S16_LE -c 2 -r 44100 -t raw ./ingresovoz &
pid1=${!}

#arecord --device="hw:0,0" -d 10 -f S16_LE -c 2 -r 44100 -t wav  ./prueba.wav

#arecord -r 48000 -f S16_LE -t raw | lame -r -s 48 --signed --little-endian -m m --vbr-new -V 9 - ~/Desktop/student-`date +%Y-%m-%d`.mp3

echo "Presionar una tecla para finalizar grabación!"

contador=0
while [ true ] ; do
read -t 3 -n 1
if [ $? = 0 -o "$contador" -gt 2 ] ; then
kill -TERM $pid1

#convertir a mp3 ./ingresovoz - CORREGIR
lame -r -s 48 --signed --little-endian -m m --vbr-new -V 9 ./ingresovoz.wav ./ingresovoz.mp3
exit ;
echo "Fin de grabación."
else
let contador++
echo "GRABANDO... "
fi
done
