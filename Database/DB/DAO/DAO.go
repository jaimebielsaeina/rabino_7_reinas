package DAO

import (
	"DB/VO"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "frances"
	password = "1234"
	dbname   = "Pro_Soft"
)

type JugadoresDAO struct{}

func (jDAO *JugadoresDAO) AddJugador(jVO VO.JugadoresVO) {

	//String para la conexión
	psqlcon := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	//abrir base de datos
	db, err := sql.Open("postgres", psqlcon)
	CheckError(err)

	//cerrar base de datos
	defer db.Close()

	//Añadir jugador j
	addj := "INSERT INTO JUGADORES VALUES ($1, $2, $3, $4)"
	_, e := db.Exec(addj, jVO.GetNombre(), jVO.GetContra(), jVO.GetDescrip(), jVO.GetCodigo())
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
