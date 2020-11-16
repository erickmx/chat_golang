# chat_golang

## Examen Parcial 01

## Objetivo

Desarrollar una sala de chat donde los usuarios _(clientes)_ que se conecten al servidor puedan enviar/recibir mensajes de texto así como archivos de los otros usuarios.

## Descripción

Implementar un cliente-servidor con las siguientes características:

- El servidor será el encargado (intermediario) de coordinar el envío de mensajes o archivos.
- El servidor tendrá las siguientes opciones:
  1. Mostrar los mensajes/nombre de los archivos enviados.
  2. Opción para respaldar en un archivo de texto los mensajes/nombre de los archivos enviados.
  3. Terminar servidor.

Los clientes tendrán las siguientes opciones:

- Enviar un mensaje de texto al servidor.
- Enviar un archivo (con opción para escribir la ubicación del archivo a enviar).
- Mostrar los mensajes recibidos del servidor (como si fuera una ventana de chat).

Hay que tener en cuenta lo siguiente:

- El servidor deberá de mantener una _lista_ de clientes, los cuales en cualquier momento pueden enviar un mensaje de texto o un archivo. Con lo anterior, el servidor siempre estará _escuchando_ por mensajes de entrada de los clientes.
- Cuando un cliente envíe una señal para desconectarse, este deberá ser liberado de la _lista._

## Rúbrica

- Envío de mensajes de texto al servidor (20%).
- Envío de mensajes de texto desde un cliente, pasando por el servidor y el servidor envía el mensaje a los clientes (20%).
- Envío de archivos al servidor y el servidor a los clientes (30%).
- Respaldo de los mensajes/nombres de archivos del lado del servidor (10%).
- Mostrar mensajes/nombres de archivos en servidor (10%).
- Mostrar mensajes/nombres de archivos en cliente (10%).
