package main

import (
	"DB/DAO"
	"DB/VO"
)

func main() {
	jVO := VO.NewJugadorVO("Adrian", "1234", make([]byte, 1), "Hola esto es una prueba", "#Adrian")
	jDAO := DAO.JugadoresDAO{}
	jDAO.AddJugador(*jVO)
}
