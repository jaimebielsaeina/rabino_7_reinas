/*
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	idCliente := 0
	// Escuchar en un puerto específico
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Servidor escuchando en puerto 8080...")

	for {
		// Aceptar conexiones entrantes
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn, idCliente)
		idCliente++;
	}
}

func handleConnection(conn net.Conn, idCliente int) {
	fin := false
	for !fin {
		// Leer los mensajes del cliente
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			fin = true
			return
		}

		// Comprobar el mensaje
		m := strings.TrimSpace(message)
		if m == "nuevaPartida" {

			// Crear la partida
			//idPartida := crearPartida()

			// Añadir jugador
			//anyadirJugador(idPartida, idCliente)

			// Enviar una respuesta al cliente
			conn.Write([]byte("1"))

		} else if m == "unirsePartida" {




		} else if m == "empezarPartida" {




		} else if m == "jugada" {




		} else if m == "chat" {



		} else if m == "finPartida" {
			fin = true


		}
	}
	// Cerrar la conexión
	conn.Close()
}

*/
