package VO

type JugadoresVO struct {
	nombre string
	contra string
	perfil []byte
	descrp string
	codigo string
}

func NewJugadorVO(nombre string, contra string, perfil []byte, descrp string, codigo string) *JugadoresVO {
	j := JugadoresVO{nombre: nombre, contra: contra, perfil: perfil, descrp: descrp, codigo: codigo}
	return &j
}

func (j *JugadoresVO) GetNombre() string {
	return j.nombre
}

func (j *JugadoresVO) GetContra() string {
	return j.contra
}

func (j *JugadoresVO) GetPerfil() []byte {
	return j.perfil
}

func (j *JugadoresVO) GetDescrip() string {
	return j.descrp
}

func (j *JugadoresVO) GetCodigo() string {
	return j.codigo
}

type AmistadVO struct {
	estado string
	usr1   string
	usr2   string
}

func NewAmistadVO(estado string, usr1 string, usr2 string) *AmistadVO {
	a := AmistadVO{estado: estado, usr1: usr1, usr2: usr2}
	return &a
}

func (a *AmistadVO) GetEstado() string {
	return a.estado
}

func (a *AmistadVO) GetUsr1() string {
	return a.usr1
}

func (a *AmistadVO) GetUsr2() string {
	return a.usr2
}

type CartaVO struct {
	numero int
	palo   string
	foto   []byte
}

func NewCartaVO(numero int, palo string, foto []byte) *CartaVO {
	c := CartaVO{numero: numero, palo: palo, foto: foto}
	return &c
}

func (c *CartaVO) GetNumero() int {
	return c.numero
}

func (c *CartaVO) GetPalo() string {
	return c.palo
}

func (c *CartaVO) GetFoto() []byte {
	return c.foto
}
