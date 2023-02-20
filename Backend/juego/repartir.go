package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/lists/doublylinkedlist"
)

type tablero struct {
	Mazo          *doublylinkedlist.List
	Descartes     *doublylinkedlist.List
	Combinaciones *list.List //Es una lista de doublylinkedlist donde se guardan las cartas jugadas(trios y escaleras en cada lista)
}

type Carta struct { //Struct utilizado para definir la estructura de datos que representa las cartas
	Valor int
	Palo  int
	Color int
}

func compararCartasN(a Carta, b Carta) int { //Función parte del sort encargada de filtrar las cartas por valor y color
	if a.Valor < b.Valor {
		return -1
	} else if a.Valor > b.Valor {
		return 1
	} else {
		if a.Color < b.Color {
			return -1
		} else if a.Color > b.Color {
			return 1
		} else {
			return 0
		}
	}
}

func compararCartasE(a Carta, b Carta) int { //Función parte del sort encargada de filtrar las cartas por palo y valor
	if a.Palo < b.Palo {
		return 1
	} else if a.Palo > b.Palo {
		return -1
	} else {
		if a.Valor < b.Valor {
			return 1
		} else if a.Valor > b.Valor {
			return -1
		} else {
			return 0
		}
	}

}

func creacionBaraja(list *doublylinkedlist.List) { //Función que inicializa la baraja de cartas del sistema
	carta := Carta{0, 0, 0}
	for i := 1; i <= 2; i++ {
		carta.Color = i
		for j := 1; j <= 4; j++ {
			carta.Palo = j
			for k := 1; k <= 13; k++ {
				carta.Valor = k
				list.Add(carta)
			}
		}
	}
	carta.Valor = 0 // Se añaden 2 comodines
	carta.Color = 1
	list.Add(carta)
	carta.Color = 2
	list.Add(carta)
}

func repartirMano(list *doublylinkedlist.List) *doublylinkedlist.List { //Función encargada de, a partir de la creación de la baraja de cartas, repartir 14 de ellas
	i := 96
	listR := doublylinkedlist.New()
	for j := 0; j < 14; j++ {
		r := rand.Intn(list.Size()) + 1 //Crea aleatorio
		value, ok := list.Get(r)        //Obtiene el valor a repartir
		for !ok {
			fmt.Println("Lista no contiene el valor", r)
			r = rand.Intn(i) + 1
			value, ok = list.Get(r)
		}
		listR.Add(value) //Lo añade a la mano
		list.Remove(r)   //Lo borra

	}

	return listR
}

func mostrarMano(mano *doublylinkedlist.List) { //Función que muestra los valores de la mano repartida
	mano.Each(func(index int, value interface{}) {
		fmt.Printf("%d: %v\n", index, value)
	})
}

// Función encargada de encontrar una escalera en la mano, devuelve los puntos del trio, las
// cartas que lo forman y si se ha encontrado trio
func calcularEscalerasJoker(mano *doublylinkedlist.List, joker *doublylinkedlist.List) (int,
	*doublylinkedlist.List, *doublylinkedlist.List, bool) {
	puntos := 0
	comb := doublylinkedlist.New()
	// ordenar la mano por palos de menor a mayor
	mano = SortStart(mano, 1)
	nuevoPalo := true
	hay_as := false
	ind_as := 0
	esc := false
	num_j := joker.Size()

	no_elim := -1
	if num_j > 0 {
		// bucle hasta que recorre toda la mano o encuentra una escalera
		for i := 0; i < mano.Size() && !esc; i++ {
			num_j_anyadidos := 0
			num_c := 1
			puntos_t := 0
			v1, _ := mano.Get(i)
			carta1, _ := v1.(Carta)
			// comprobar si hay as en el palo
			if nuevoPalo {
				hay_as = carta1.Valor == 1
				ind_as = i
			}
			if carta1.Valor >= 10 {
				puntos_t = puntos_t + 10
			} else {
				puntos_t = puntos_t + carta1.Valor
			}
			// lista temporal donde añadir las cartas que se van encontrando de la escalera
			l := *doublylinkedlist.New()
			l.Add(carta1)
			i_inf := i
			hay_esc := true
			mirar_j := false
			for hay_esc {
				v2, _ := mano.Get(i + 1)
				carta2, _ := v2.(Carta)
				if carta1.Palo == carta2.Palo {
					nuevoPalo = false
				} else {
					nuevoPalo = true
				}
				// comprobar si las dos cartas son escalera
				if carta1.Valor+1 == carta2.Valor && carta1.Palo == carta2.Palo {
					//añadir la nueva carta a l
					l.Add(carta2)
					if carta2.Valor >= 10 {
						puntos_t = puntos_t + 10
					} else {
						puntos_t = puntos_t + carta2.Valor
					}
					num_c += 1
					i++
					carta1 = carta2
				} else if carta1.Valor == 13 && hay_as && !mirar_j {
					// hay escalera valida de la forma 11 12 AS
					// recupero la carta del as
					as, _ := mano.Get(ind_as)
					as_c, _ := as.(Carta)
					l.Add(as_c)
					puntos_t = puntos_t + 11
					num_c += 1
					mirar_j = true
				} else if carta1.Valor == carta2.Valor && carta1.Palo == carta2.Palo {
					// dos cartas con el mismo numero seguidas, avanzo indice
					i++
					if no_elim == -1 {
						no_elim = i
					} else {
						no_elim = no_elim*100 + i
					}

				} else if num_j > 0 { // mirar si puedo añadir el joker para hacer escalera
					v_joker, _ := joker.Get(num_j - 1)
					c_joker, _ := v_joker.(Carta)
					l.Add(c_joker)
					num_j_anyadidos++
					if carta1.Valor == 13 && !hay_as { //joker como as
						puntos_t = puntos_t + 11
					} else if carta1.Valor >= 10 {
						puntos_t = puntos_t + 10
						if mirar_j {
							l.Swap(l.Size()-1, l.Size()-3)
							l.Swap(l.Size()-1, l.Size()-2)
						}
					} else {
						puntos_t = puntos_t + carta1.Valor + 1
					}
					num_j--
					num_c++
					carta1 = Carta{carta1.Valor + 1, carta1.Palo, carta1.Color}
					if hay_as {
						hay_esc = false
					}
				} else {
					hay_esc = false
				}

			}
			if num_c >= 3 && num_c-num_j_anyadidos >= 2 {
				// si el numero de cartas seguidas ha sido >=3, escalera valida
				puntos += puntos_t
				// añado l a la combinación a devolver
				comb.Add(l)
				if !mirar_j {
					// si no hay AS, borro de la mano las cartas de los indices seguidos que correspondan
					k := no_elim % 100
					no_elim = no_elim / 100
					for j := i; j >= i_inf; j-- {
						if j != k {
							mano.Remove(j)
							if j < k {
								k = no_elim % 100
								no_elim = no_elim / 100
								if k == 0 {
									k = -1
								}
							}
						}

					}
				} else {
					// si hay AS, borro las cartas de la mano de los indices seguidos, ADEMAS del indice del AS
					k := no_elim % 100
					no_elim = no_elim / 100
					for j := i; j >= i_inf; j-- {
						if j != k {
							mano.Remove(j)
							if j < k {
								k = no_elim % 100
								no_elim = no_elim / 100
								if k == 0 {
									k = -1
								}
							}
						}
					}
					mano.Remove(ind_as)
				}
				for j := 0; j < num_j_anyadidos; j++ {
					joker.Remove(0)
				}
				esc = true
			} else {
				num_j = num_j + num_j_anyadidos
				mirar_j = false
			}
		}
	}

	return puntos, comb, joker, esc
}

// Función encargada de encontrar una escalera en la mano, devuelve los puntos del trio, las
// cartas que lo forman y si se ha encontrado trio
func calcularEscaleras(mano *doublylinkedlist.List) (int, *doublylinkedlist.List, bool) {
	puntos := 0
	comb := doublylinkedlist.New()
	// ordenar la mano por palos de menor a mayor
	mano = SortStart(mano, 1)
	nuevoPalo := true
	hay_as := false
	ind_as := 0
	esc := false
	no_elim := -1
	// bucle hasta que recorre toda la mano o encuentra una escalera
	for i := 0; i < mano.Size() && !esc; i++ {
		num_c := 1
		puntos_t := 0
		v1, _ := mano.Get(i)
		carta1, _ := v1.(Carta)
		// comprobar si hay as en el palo
		if nuevoPalo {
			hay_as = carta1.Valor == 1
			ind_as = i
		}
		if carta1.Valor >= 10 {
			puntos_t = puntos_t + 10
		} else {
			puntos_t = puntos_t + carta1.Valor
		}
		// lista temporal donde añadir las cartas que se van encontrando de la escalera
		l := *doublylinkedlist.New()
		l.Add(carta1)
		i_inf := i
		hay_esc := true
		borrar_as := false
		for hay_esc {
			v2, _ := mano.Get(i + 1)
			carta2, _ := v2.(Carta)
			if carta1.Palo == carta2.Palo {
				nuevoPalo = false
			} else {
				nuevoPalo = true
			}
			// comprobar si las dos cartas son escalera
			if carta1.Valor+1 == carta2.Valor && carta1.Palo == carta2.Palo {
				//añadir la nueva carta a l
				l.Add(carta2)
				if carta2.Valor >= 10 {
					puntos_t = puntos_t + 10
				} else {
					puntos_t = puntos_t + carta2.Valor
				}
				num_c += 1
				i++
			} else if num_c >= 2 && carta1.Valor == 13 && hay_as {
				// hay escalera valida de la forma 11 12 AS
				// recupero la carta del as
				as, _ := mano.Get(ind_as)
				as_c, _ := as.(Carta)
				l.Add(as_c)
				puntos_t = puntos_t + 11
				num_c += 1
				hay_esc = false
				borrar_as = true
			} else if carta1.Valor == carta2.Valor && carta1.Palo == carta2.Palo {
				// dos cartas con el mismo numero seguidas, avanzo indice
				i++
				if no_elim == -1 {
					no_elim = i
				} else {
					no_elim = no_elim*100 + i
				}
			} else {
				// no hay escalera
				hay_esc = false
			}
			carta1 = carta2
		}
		if num_c >= 3 {
			// si el numero de cartas seguidas ha sido >=3, escalera valida
			puntos += puntos_t
			// añado l a la combinación a devolver
			comb.Add(l)
			if !borrar_as {
				// si no hay AS, borro de la mano las cartas de los indices seguidos que correspondan
				k := no_elim % 100
				no_elim = no_elim / 100
				for j := i; j >= i_inf; j-- {
					if j != k {
						mano.Remove(j)
						if j < k {
							k = no_elim % 100
							no_elim = no_elim / 100
							if k == 0 {
								k = -1
							}
						}
					}
				}
			} else {
				// si hay AS, borro las cartas de la mano de los indices seguidos, ADEMAS del indice del AS
				k := no_elim % 100
				no_elim = no_elim / 100
				for j := i; j >= i_inf; j-- {
					if j != k {
						mano.Remove(j)
						if j < k {
							k = no_elim % 100
							no_elim = no_elim / 100
							if k == 0 {
								k = -1
							}
						}
					}
				}
				mano.Remove(ind_as)
			}
			esc = true
		}
	}
	return puntos, comb, esc
}

// Función encargada de encontrar un trío con joker en la mano, devuelve los puntos del trio, las
// cartas que lo forman, si se ha encontrado trio y los jokers que quedan
func calcularTriosJoker(mano *doublylinkedlist.List, joker *doublylinkedlist.List) (int,
	*doublylinkedlist.List, *doublylinkedlist.List, bool) {
	puntos := 0
	mano = SortStart(mano, 0)
	comb := doublylinkedlist.New()
	trio := false
	if !joker.Empty() {
		// bucle hasta que recorre toda la mano o encuentra un trio
		for i := 0; i < mano.Size()-2 && !trio; i++ {
			i_inf := i
			v1, _ := mano.Get(i)
			carta1, _ := v1.(Carta)
			v2, _ := mano.Get(i + 1)
			carta2, _ := v2.(Carta)
			if carta1.Valor == carta2.Valor {
				// las tres cartas tienen el mismo numero
				if carta1.Palo != carta2.Palo {
					// las tres cartas son de distinto palo
					trio = true
					// lista donde añadir las cartas del trio
					l := *doublylinkedlist.New()
					l.Add(carta1)
					l.Add(carta2)
					v_joker, _ := joker.Get(0)
					c_joker, _ := v_joker.(Carta)
					l.Add(c_joker)
					if carta1.Valor == 1 {
						puntos = puntos + 11*3
					} else if carta1.Valor >= 10 {
						puntos = puntos + 10*3
					} else {
						puntos = puntos + carta1.Valor*3
					}
					i += 1

					for j := i; j >= i_inf; j-- {
						// se eliminan de la mano las cartas que hemos cojido
						mano.Remove(j)
					}
					joker.Remove(0) // borro joker
					comb.Add(l)
				}
			}
		}
	}
	return puntos, comb, joker, trio
}

// Función encargada de encontrar un trío en la mano, devuelve los puntos del trio, las
// cartas que lo forman y si se ha encontrado trio
func calcularTrios(mano *doublylinkedlist.List) (int, *doublylinkedlist.List, bool) {
	puntos := 0
	mano = SortStart(mano, 0)
	comb := doublylinkedlist.New()
	trio := false
	// bucle hasta que recorre toda la mano o encuentra un trio
	for i := 0; i < mano.Size()-2 && !trio; i++ {
		i_inf := i
		palo := 0
		v1, _ := mano.Get(i)
		carta1, _ := v1.(Carta)
		v2, _ := mano.Get(i + 1)
		carta2, _ := v2.(Carta)
		v3, _ := mano.Get(i + 2)
		carta3, _ := v3.(Carta)
		if carta1.Valor == carta2.Valor && carta2.Valor == carta3.Valor {
			// las tres cartas tienen el mismo numero
			if carta1.Palo != carta2.Palo && carta2.Palo != carta3.Palo && carta1.Palo != carta3.Palo {
				// las tres cartas son de distinto palo
				trio = true
				// lista donde añadir las cartas del trio
				l := *doublylinkedlist.New()
				l.Add(carta1)
				l.Add(carta2)
				l.Add(carta3)
				// sumo los palos de las cartas, luego se explica porqué
				palo = palo + carta1.Palo + carta2.Palo + carta3.Palo
				if carta1.Valor == 1 {
					puntos = puntos + 11*3
				} else if carta1.Valor >= 10 {
					puntos = puntos + 10*3
				} else {
					puntos = puntos + carta1.Valor*3
				}
				i += 2
				v4, _ := mano.Get(i + 1)
				carta4, _ := v4.(Carta)
				palo += carta4.Palo
				// la suma de los cuatro palos 1+2+3+4 = 10
				// si al añadir la cuarta carta el valor que teniamos en palo + el palo de la nueva carta
				// es == 10, entonces significa que las 4 cartas tienen palo diferente, por eso puede
				// formar el cuarteto
				if carta1.Valor == carta4.Valor && palo == 10 {
					l.Add(carta4)
					if carta1.Valor == 1 {
						puntos = puntos + 11
					} else if carta1.Valor >= 10 {
						puntos = puntos + 10
					} else {
						puntos = puntos + carta1.Valor
					}
					i += 1
				}
				for j := i; j >= i_inf; j-- {
					// se eliminan de la mano las cartas que hemos cojido
					mano.Remove(j)
				}
				comb.Add(l)
			}
		}
	}
	return puntos, comb, trio
}

func separarJokers(mano *doublylinkedlist.List) (*doublylinkedlist.List, *doublylinkedlist.List) {
	mano = SortStart(mano, 0)
	joker := doublylinkedlist.New()
	hay_j := true
	for hay_j {
		v, _ := mano.Get(mano.Size() - 1)
		carta, _ := v.(Carta)
		if carta.Valor == 0 {
			joker.Add(carta)
			mano.Remove(mano.Size() - 1)
		} else {
			hay_j = false
		}
	}
	return mano, joker
}

func descarteBot(mazo *doublylinkedlist.List, mano *doublylinkedlist.List, descarte *doublylinkedlist.List) {
	mano = SortStart(mano, 0)
	finTurno(mazo, mano, descarte, mano.Size()-1)
}

func calcularPuntosPosibles(mano *doublylinkedlist.List) (int, *doublylinkedlist.List) { //Función encargada de revisar los puntos posibles de una mano
	puntos := 0
	esc := true
	comb := doublylinkedlist.New()
	mano, joker := separarJokers(mano)
	trio := true
	for esc {
		// bucle para encontrar todas las escaleras
		puntos_m, combE, escR := calcularEscaleras(mano)
		puntos += puntos_m
		if escR {
			//añade a comb la nueva escalera encontrada
			comb.Add(combE)
		}
		esc = escR
	}
	esc_j := true
	for esc_j {
		puntos_m, combE, jokerR, escR := calcularEscalerasJoker(mano, joker)
		puntos += puntos_m
		if escR {
			//añade a comb el nuevo trio encontrado
			comb.Add(combE)
		}
		esc_j = escR
		joker = jokerR
	}
	for trio {
		// bucle para encontrar todos los trios
		puntos_m, combT, trioR := calcularTrios(mano)
		puntos += puntos_m
		if trioR {
			//añade a comb el nuevo trio encontrado
			comb.Add(combT)
		}
		trio = trioR
	}
	trio_j := true
	for trio_j {
		puntos_m, combT, jokerR, trioR := calcularTriosJoker(mano, joker)
		puntos += puntos_m
		if trioR {
			//añade a comb el nuevo trio encontrado
			comb.Add(combT)
		}
		trio_j = trioR
		joker = jokerR
	}

	return puntos, comb
}

func partition(mano *doublylinkedlist.List, low, high int, tipo int) (*doublylinkedlist.List, int) { //Función del sort encargada de particionar los datos
	v1, _ := mano.Get(high)
	carta1, _ := v1.(Carta)
	i := low
	for j := low; j < high; j++ {
		v2, _ := mano.Get(j)
		carta2, _ := v2.(Carta)
		if tipo == 0 {
			if compararCartasN(carta1, carta2) == -1 {
				mano.Swap(i, j)
				i++
			}
		} else if tipo == 1 {
			if compararCartasE(carta1, carta2) == -1 {
				mano.Swap(i, j)
				i++
			}
		}
	}
	mano.Swap(i, high)
	return mano, i
}

func Sort(mano *doublylinkedlist.List, low, high int, tipo int) *doublylinkedlist.List { //Función inicial del sort
	if low < high {
		var p int
		mano, p = partition(mano, low, high, tipo)
		mano = Sort(mano, low, p-1, tipo)
		mano = Sort(mano, p+1, high, tipo)
	}
	return mano
}

func SortStart(mano *doublylinkedlist.List, tipo int) *doublylinkedlist.List { //Función inicial del sort
	return Sort(mano, 0, mano.Size()-1, tipo)
}

func robarCarta(list *doublylinkedlist.List, mano *doublylinkedlist.List) { //Función encargada de robar una carta del mazo
	r := rand.Intn(list.Size()) + 1 //Obtiene un número aleatorio de la lista
	value, ok := list.Get(r)        //Obtiene el valor de la carta de la lista
	if ok {
		mano.Add(value) //Añade el valor a la mano
		list.Remove(r)  //Elimina el valor del mazo
	}

}

func finTurno(mazo *doublylinkedlist.List, mano *doublylinkedlist.List, descarte *doublylinkedlist.List, i int) {
	value, _ := mano.Get(i) //Obtiene el valor de la mano a descartar
	mano.Remove(i)          //Elimina el valor de la mano
	descarte.Add(value)     //Añade el valor a descartes
	if descarte.Size() > 1 {
		fmt.Println(descarte, "DESCARTE METE A MAZO") //Si hay más de un valor en descartes lo añade a la lista de mazo
		value, _ = descarte.Get(0)
		mazo.Add(value)
		descarte.Remove(0)
	}

}

func suma51(jugada *doublylinkedlist.List) bool { // cuenta los puntos de la primera jugada que se hace y devuelve true si llega a 51
	total := 0
	for i := 0; i <= jugada.Size(); i++ {
		v1, _ := jugada.Get(i)
		carta, _ := v1.(Carta)
		if carta.Valor == 1 {
			total += 11
		} else if carta.Valor >= 10 {
			total += 10
		} else {
			total += carta.Valor
		}
	}
	if total >= 51 {
		return true
	} else {
		fmt.Println(total)
		return false
	}
}

/*
Devuelve true cuabdo sea ha podido realizar la juaga con exito y
false en caso contrario
*/
func abrir(jugada *doublylinkedlist.List, mano *doublylinkedlist.List, t *tablero) bool { //falta comprobar trios y escaleras
	jugada = SortStart(jugada, 0)
	if !EscaleraValida(jugada) && !TrioValido(jugada) {
		return false
	}
	listaC := doublylinkedlist.New()
	for i := 0; i <= jugada.Size()-1; i++ {
		v1, _ := jugada.Get(i)
		carta, _ := v1.(Carta)

		ind := mano.IndexOf(carta)
		mano.Remove(ind)
		fmt.Println("carta eliminada", carta)
		listaC.Add(carta)
	}
	//listaC = SortStart(listaC, 0)
	t.Combinaciones.PushBack(listaC)
	return true
}

/*

// función para añadir una carta a una combinación
//Devuelve -1 si es una juagada invalida, 0 si es valida, 1 si es valida y devuelve un comodin
func anyadirCarta(jugada *doublylinkedlist.List, mano *doublylinkedlist.List, t *tablero, idCombinacion int) int {
	index := 0
	if !jugada.Empty() {
		v1, _ := jugada.Get(0)
		carta, _ := v1.(Carta)

		id_comb := 0
		for e := t.Combinaciones.Front(); e != nil; e = e.Next() {
			if id_comb == idCombinacion {
				listaC := e.Value.(*doublylinkedlist.List)
				if NumComodines(listaC) > 0 {
					/*for i := 0; i < listaC.Size(); i++ {
						cart, _ := listaC.Get(i)//Cogemos la primera carta
						cartaRef, _ := cart.(Carta)
						CartaValor = cartaRef.Valor //Sacamos este valor por si es un comodin

						if EsComodin(CartaValor){
							if i > 0 && i < listaC.Size() {
								cart_ant, _ := listaC.Get(i - 1)//Cogemos la primera carta
								carta_ant, _ := cart_ant.(Carta)
								CartaValorAnt = carta_ant.Valor //Sacamos este valor por si es un comodin
								index = i
								for u := index; u < listaC.Size() && EsComodin(CartaValorAnt) ; u++ { //Tomamos de referencia una carta que no sea comodin
									cart_ant, _ := jugada.Get(index)//Cogemos la primera carta
									carta_ant, _ := cart.(Carta)
									CartaValorAnt = carta.Valor //Sacamos este valor por si es un comodin
									index++
								}

								if EsComodin(CartaValorAnt) { //En caso de no encontrar el valor de iterrar a la derecha
									index = i
									for d := index; d >= 0 && EsComodin(CartaValorAnt); d-- {
										cart_ant, _ := jugada.Get(index)//Cogemos la primera carta
										carta_ant, _ := cart.(Carta)
										CartaValorAnt = carta.Valor //Sacamos este valor por si es un comodin
										index--
									}

									if carta.Valor == ( CartaValorAnt + Abs((index + 1) - i).(float64) ) ) {
										listaC.Remove(i) //Aqui quitamos el coodin a intercambiar
										listaC.Add(carta)
										listaC = SortStart(listaC, 0)
										if !EscaleraValida(listaC) && !TrioValido(listaC) {
											return -1
										}
										t.Combinaciones.Remove(e)
										t.Combinaciones.PushBack(listaC)
										ind := mano.IndexOf(carta)
										mano.Remove(ind)

										return 1

									}

								} else {
									if carta.Valor == ( CartaValorAnt + Abs((index - 1) - i).(float64) ) ) {
										listaC.Remove(i) //Aqui quitamos el coodin a intercambiar
										listaC.Add(carta)
										listaC = SortStart(listaC, 0)
										if !EscaleraValida(listaC) && !TrioValido(listaC) {
											return -1
										}
										t.Combinaciones.Remove(e)
										t.Combinaciones.PushBack(listaC)
										ind := mano.IndexOf(carta)
										mano.Remove(ind)

										return 1

									}
								}

							} else if index == 0 {
								cart_ant, _ := listaC.Get(i + 1)//Cogemos la primera carta
								carta_ant, _ := cart_ant.(Carta)
								CartaValorAnt = carta_ant.Valor //Sacamos este valor por si es un comodin
								index = i //Indice del comodin que miramos
								for u := index; u < listaC.Size() && EsComodin(CartaValorAnt) ; u++ { //Tomamos de referencia una carta que no sea comodin
									cart_ant, _ := jugada.Get(index)//Cogemos la primera carta
									carta_ant, _ := cart.(Carta)
									CartaValorAnt = carta.Valor //Sacamos este valor por si es un comodin
									index++
								}

								if carta.Valor == ( CartaValorAnt - Abs(((index - 1) - i).(float64) ) ) {
									listaC.Remove(i) //Aqui quitamos el coodin a intercambiar
									listaC.Add(carta)
									listaC = SortStart(listaC, 0)
									if !EscaleraValida(listaC) && !TrioValido(listaC) {
										return -1
									}
									t.Combinaciones.Remove(e)
									t.Combinaciones.PushBack(listaC)
									ind := mano.IndexOf(carta)
									mano.Remove(ind)

									return 1

								}
							}
						} else {

						}
					}*/
/*
					for i := 0; i < listaC.Size(); i++ {
						cart, _ := listaC.Get(i) //Cogemos la primera carta
						cartaRef, _ := cart.(Carta)
						CartaValor = cartaRef.Valor //Sacamos este valor por si es un comodin

						if i == 0 { //Agniadir al primero o al final
							if carta.Valor < CartaValor {
								listaC.Add(carta)
								listaC = SortStart(listaC, 0)
								if !EscaleraValida(listaC) && !TrioValido(listaC) {
									indice := listaC.IndexOf(carta)
									cart, _ := listaC.Get(indice + 1) //Cogemos la primera carta
									cartaRef, _ := cart.(Carta)
									CartaValor = cartaRef.Valor //Sacamos este valor por si es un comodin

									if EsComodin(CartaValor) {
										listaC.Remove(indice + 1)
										return 1
									}
									return -1
								}

								t.Combinaciones.Remove(e)
								t.Combinaciones.PushBack(listaC)
								ind := mano.IndexOf(carta)
								mano.Remove(ind)
							}
						} else if i == (listaC.Size() - 1) {

							if carta.Valor < CartaValor {
								listaC.Add(carta)
								listaC = SortStart(listaC, 0)
								if !EscaleraValida(listaC) && !TrioValido(listaC) {
									indice := listaC.IndexOf(carta)
									cart, _ := listaC.Get(indice - 1) //Cogemos la primera carta
									cartaRef, _ := cart.(Carta)
									CartaValor = cartaRef.Valor //Sacamos este valor por si es un comodin

									if EsComodin(CartaValor) {
										listaC.Remove(indice - 1)
										return 1
									}
									return -1
								}

								t.Combinaciones.Remove(e)
								t.Combinaciones.PushBack(listaC)
								ind := mano.IndexOf(carta)
								mano.Remove(ind)
							}

						} else {
							if carta.Valor < CartaValor {
								listaC.Add(carta)
								listaC = SortStart(listaC, 0)

								if !TrioValido(listaC) {

									if !EscaleraValida(listaC) {
										indice := listaC.IndexOf(carta)
										cart, _ := listaC.Get(indice - 1) //Cogemos la primera carta
										cartaRef, _ := cart.(Carta)
										CartaValor = cartaRef.Valor //Sacamos este valor por si es un comodin

										if EsComodin(CartaValor) {
											listaC.Remove(indice - 1)
											if EscaleraValida(listaC) {
												return 1
											} else {
												return -1
											}
										} else {
											indice = listaC.IndexOf(carta)
											cart, _ = listaC.Get(indice + 1) //Cogemos la primera carta
											cartaRef, _ = cart.(Carta)
											CartaValor = cartaRef.Valor //Sacamos este valor por si es un comodin

											if EsComodin(CartaValor) {
												listaC.Remove(indice + 1)
												if EscaleraValida(listaC) {
													return 1
												} else {
													return -1
												}
											}
										}
									}

								}
								t.Combinaciones.Remove(e)
								t.Combinaciones.PushBack(listaC)
								ind := mano.IndexOf(carta)
								mano.Remove(ind)

								return 0
							}
						}
					}

				} else {
					listaC.Add(carta)
					listaC = SortStart(listaC, 0)
					if !EscaleraValida(listaC) && !TrioValido(listaC) {
						return -1
					}
					t.Combinaciones.Remove(e)
					t.Combinaciones.PushBack(listaC)
					ind := mano.IndexOf(carta)
					mano.Remove(ind)

					return 0
				}
			}
			id_comb++
		}
	}
}

*/

func mostrarTablero(t tablero) {
	fmt.Println("MAZO: ", t.Mazo)

	fmt.Println("DESCARTES: ", t.Descartes)

	fmt.Println("COMBINACIONES: ", t.Combinaciones)

	l := t.Combinaciones
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("---------------------------------------\n")

}

func iniciarTablero() (tablero, *doublylinkedlist.List) {

	rand.Seed(time.Now().UnixNano())
	mazo := doublylinkedlist.New()
	descarte := doublylinkedlist.New()

	creacionBaraja(mazo)

	mano := repartirMano(mazo)

	t := tablero{mazo, descarte, list.New()}

	return t, mano
}

func realizarJugada(t *tablero, mano *doublylinkedlist.List, jugada int, i int, cartasAjugar *doublylinkedlist.List) {
	switch jugada {
	case 0: //Descarte
		finTurno(t.Mazo, mano, t.Descartes, i)
		return
	case 1: //Robar
		robarCarta(t.Mazo, mano)
		return
	case 2: //Abrir
		if suma51(cartasAjugar) {
			abrir(cartasAjugar, mano, t)
		}
		cartasAjugar.Clear()
		return
	case 3: //Añadir 1 carta a una combinación existente
		//anyadirCarta(cartasAjugar, mano, t, 0)
		cartasAjugar.Clear()
		return
	default:
	}
}

/*
Pre: TRUE
Post: return true si es un comodin, es decir vale 0

	false en caso contrario
*/
func EsComodin(valor int) bool {
	return valor == 0
}

/*
Pre: TRUE
Post: devuelve el numero de comodines en la lista
*/
func NumComodines(jugada *doublylinkedlist.List) int {
	num_comodines := 0
	for j := 0; j < jugada.Size(); j++ {
		cart, _ := jugada.Get(j)
		carta, _ := cart.(Carta)
		ValorCarta := carta.Valor

		if EsComodin(ValorCarta) {
			num_comodines++
		}
	}
	return num_comodines
}

/*
Pre: lista ordenada en orden de jugada (Consideramos que el comodin es el 0)
Post: return true si es una escalera válida en el juego del Rabino, y

	false en caso contrario
*/
func EscaleraValida(jugada *doublylinkedlist.List) bool {

	if jugada.Empty() { //Si la lista de la jugada es vacia
		return false
	}

	num_cartas := jugada.Size()

	//COMPROBACION: NUMERO DE CARTAS VALIDO
	//Escalera maxima: 1,2,3,4,5,6,7,8,9,10,Sota(11),Caballo(12),Rey(13),As(1 o 0)
	if num_cartas > 14 { //Tamagno maximo de escalera 14
		return false
	}

	num_comodines := NumComodines(jugada)

	//COMPROBACION: NUMERO DE COMODINES VALIDO
	if num_comodines > (num_cartas - 2) { //Numero de comodines es como mucho num_cartas - 2
		return false
	}

	//COMPROBACION: TIENEN EL MISMO PALO
	index := 0                   //Indice inicial
	cart, _ := jugada.Get(index) //Cogemos la primera carta
	carta, _ := cart.(Carta)
	CartaValorRef := carta.Valor //Sacamos este valor por si es un comodin
	PaloCartaRef := carta.Palo

	for EsComodin(CartaValorRef) { //Tomamos de referencia una carta que no sea comodin
		index++                      //Miramos la siguiente carta
		cart, _ := jugada.Get(index) //Cogemos la primera carta
		carta, _ := cart.(Carta)
		CartaValorRef = carta.Valor //Sacamos este valor por si es un comodin
		PaloCartaRef = carta.Palo
	}

	for u := index + 1; u < jugada.Size(); u++ { //Miramos que tenga todas las cartas el mismo palo
		cart1, _ := jugada.Get(u)
		carta1, _ := cart1.(Carta)
		CartaValorMirar := carta1.Valor //Sacamos este valor por si es un comodin
		PaloCartaMirar := carta1.Palo

		if PaloCartaRef != PaloCartaMirar && !EsComodin(CartaValorRef) && !EsComodin(CartaValorMirar) { //Si tiene distinto palo, no valido
			return false
		}
	}

	//COMPROBACION: VALOR DE CARTAS CRECIENTE
	index = 0 //Indice inicial
	cart, _ = jugada.Get(index)
	carta, _ = cart.(Carta)
	CartaValorRef = carta.Valor //Sacamos este valor de la primera carta

	for EsComodin(CartaValorRef) { //Tomamos de referencia una carta que no sea comodin
		index++ //Miramos la siguiente carta
		cart, _ = jugada.Get(index)
		carta, _ = cart.(Carta)
		CartaValorRef = carta.Valor
	}

	//El numero de comodines que hay delante de 1 debe ser cero, el numero de comodines que pueden ir delante
	// de 2 es 1, y asi sucesibamente. Index en este caso tambien tomaria el numero de comodines que hay delante
	if (CartaValorRef - index) <= 0 {
		return false
	}

	for j := index + 1; j < jugada.Size(); j++ { //Miramos si el valor de las cartas es creciente
		//Empezamos a comparar con cartas posteriores a la de referencia
		CartaValorRef++
		cart1, _ := jugada.Get(j)
		carta1, _ := cart1.(Carta)
		CartaValorMirar := carta1.Valor //Sacamos este valor de la carta

		if !EsComodin(CartaValorMirar) { //Si es un comodin seguro que valdra para la escalera
			if CartaValorMirar != 1 || CartaValorRef != 14 { //Esta condicion no se cumple cuando despues del Rey(13), se pone un As(1)
				if CartaValorMirar != CartaValorRef { //Si concuerda el valor con lo que deberia dar(p.ejem: 2 != 15(valor despues de As))
					return false
				}
			}
		} else if CartaValorRef > 14 { //Si la jugada continua despues de ..., Rey, As,...sera erronea
			return false
		}
	}

	return true //Si cumple todas las condiciones
}

/*
Pre: TRUE
Post: return true si es una trio o cuarteto válida en el juego del Rabino, y

	false en caso contrario
*/
func TrioValido(jugada *doublylinkedlist.List) bool {

	if jugada.Empty() { //Si la lista de la jugada es vacia
		return false
	}

	//COMPROBACION: ES UN TRIO O UN CUARTETO
	if jugada.Size() < 3 || jugada.Size() > 4 { //El tamaño de la jugada puede ser 3 o 4
		return false
	}

	//COMPROBACION: NUMERO DE COMODINES VALIDO
	num_comodines := NumComodines(jugada)
	if num_comodines > 1 { //Numero de comodines es como mucho 1
		return false
	}

	//COMPROBACION: TIENEN EL MISMO VALOR
	index := 0
	cart, _ := jugada.Get(index)
	carta, _ := cart.(Carta)
	ValorCartaRef := carta.Valor

	for EsComodin(ValorCartaRef) { //Tomamos de referencia una carta que no sea comodin
		index++ //Miramos la siguiente carta
		cart, _ = jugada.Get(index)
		carta, _ = cart.(Carta)
		ValorCartaRef = carta.Valor
	}

	for i := index + 1; i < jugada.Size(); i++ { //Comprobamos que todas las cartas tengan el mismo valor

		cart1, _ := jugada.Get(i)
		carta1, _ := cart1.(Carta)
		ValorCarta := carta1.Valor

		if ValorCartaRef != ValorCarta && !EsComodin(ValorCarta) { //Si la carta a comparar es un comodin sera valida
			return false
		}
	}

	//COMPROBACION: TIENEN DISTINTO PALO
	for j := 0; j < jugada.Size(); j++ {

		cart, _ := jugada.Get(j)
		carta, _ := cart.(Carta)
		CartaValorRef := carta.Valor //Sacamos este valor por si es un comodin
		PaloCartaRef := carta.Palo

		if !EsComodin(CartaValorRef) { //Si es un comodin seguro que sera valido
			for u := j + 1; u < jugada.Size(); u++ { //Miramos que tenga todas las cartas el mismo palo
				cart1, _ := jugada.Get(u)
				carta1, _ := cart1.(Carta)
				CartaValorMirar := carta1.Valor //Sacamos este valor por si es un comodin
				PaloCartaMirar := carta1.Palo

				if PaloCartaRef == PaloCartaMirar && !EsComodin(CartaValorMirar) { //Si tiene distinto palo, no valido
					return false
				}
			}
		}
	}

	return true //Si cumple todas las condiciones

}

func main() {
	fmt.Println("Hola1")
	rand.Seed(time.Now().UnixNano())
	mazo := doublylinkedlist.New()
	descarte := doublylinkedlist.New()
	//i := 4
	fmt.Println("Hola2")

	creacionBaraja(mazo)
	fmt.Println("Hola3")

	mano := repartirMano(mazo)
	fmt.Println("Hola4")

	mostrarMano(mano)
	fmt.Println("Hola5")

	emp := false
	for !emp {
		robarCarta(mazo, mano)
		fmt.Println("MANO CARTA ROBADA")
		mostrarMano(mano)

		puntos, comb := calcularPuntosPosibles(mano)
		fmt.Println("Puntos ", puntos)

		iterator := comb.Iterator()
		i := 0
		for iterator.Next() {
			i++
			fmt.Println("Combinación", i)
			l := iterator.Value()
			lista := l.(*doublylinkedlist.List)
			iterator2 := lista.Iterator()
			for iterator2.Next() {
				c := iterator2.Value()
				cartas := c.(doublylinkedlist.List)
				iterator_c := cartas.Iterator()
				for iterator_c.Next() {
					v := iterator_c.Value()
					valor := v.(Carta)
					fmt.Println(valor)
				}
			}

		}
		descarteBot(mazo, mano, descarte)
		if puntos >= 51 {
			emp = true
		} else {
			iterator := comb.Iterator()
			i := 0
			for iterator.Next() {
				i++
				l := iterator.Value()
				lista := l.(*doublylinkedlist.List)
				iterator2 := lista.Iterator()
				for iterator2.Next() {
					c := iterator2.Value()
					cartas := c.(doublylinkedlist.List)
					iterator_c := cartas.Iterator()
					for iterator_c.Next() {
						v := iterator_c.Value()
						valor := v.(Carta)
						mano.Add(valor)
					}
				}

			}
		}
		fmt.Println("Mano final: ")
		mostrarMano(mano)
	}

	// for i := 0; i < comb.Size(); i++ {
	// 	fmt.Printf("combinacion %d: \n", i)
	// 	l, _ := comb.Get(i)
	// 	fmt.Printf("lista : %v\n", l)
	// 	lista, ok := l.(*doublylinkedlist.List)
	// 	if !ok {
	// 		fmt.Print("fail")
	// 	}
	// 	fmt.Printf("tam %d: ", lista.Size())
	// 	for j := 0; j < lista.Size(); j++ {
	// 		c, _ := lista.Get(j)
	// 		carta, _ := c.(Carta)
	// 		fmt.Printf("tamc %d: ", carta.Valor)
	// 		fmt.Printf("carta %d: ", i)
	// 		fmt.Println(carta)
	// 	}
	// }

	// comb.Each((func(index int, value interface{}) {
	// 	l, _ := value.(*doublylinkedlist.List)
	// 	fmt.Printf("combinacion %d: %v\n", index, value)
	// 	l.Each((func(index int, value interface{}) {
	// 		lista, _ := value.(*doublylinkedlist.List)
	// 		fmt.Printf("combinacion lis %d: %v\n", index, value)
	// 	}))
	// }))
	// fmt.Println(descarte)

	// finTurno(mazo, mano, descarte, i)
	// fmt.Println("MANO DESCARTE HECHO")
	// fmt.Println(descarte)
	// mostrarMano(mano)

	// robarCarta(mazo, mano)
	// fmt.Println("MANO CARTA ROBADA")
	// mostrarMano(mano)

	// fmt.Println("Puntos ", calcularPuntosPosibles(mano)) //Revisar as
	// fmt.Println(descarte)

	// finTurno(mazo, mano, descarte, i)
	// fmt.Println("MANO DESCARTE HECHO")
	// fmt.Println(descarte)
	// mostrarMano(mano)

	// rand.Seed(time.Now().UnixNano())
	// mazo := doublylinkedlist.New()
	// descarte := doublylinkedlist.New()

	// fmt.Println("Hola2")

	// creacionBaraja(mazo)
	// fmt.Println("Hola3")

	// mano := repartirMano(mazo)
	// fmt.Println("Hola4")

	// mostrarMano(mano)
	// t := tablero{mazo, descarte, list.New()}
	// t.Combinaciones.PushBack(4555555)

	// i := 8
	// t, mano := iniciarTablero()

	// mostrarMano(mano)
	// fmt.Println("--------------------------")
	// t, mano := iniciarTablero()
	// mostrarTablero(t)
	// mostrarMano(mano)

	// jugada := doublylinkedlist.New()
	// v1, _ := mano.Get(0)
	// carta, _ := v1.(Carta)
	// jugada.Add(carta)
	// v1, _ = mano.Get(1)
	// carta, _ = v1.(Carta)
	// jugada.Add(carta)
	// v1, _ = mano.Get(2)
	// carta, _ = v1.(Carta)
	// jugada.Add(carta)
	// v1, _ = mano.Get(3)
	// carta, _ = v1.(Carta)
	// jugada.Add(carta)
	// v1, _ = mano.Get(4)
	// carta, _ = v1.(Carta)
	// jugada.Add(carta)
	// v1, _ = mano.Get(5)
	// carta, _ = v1.(Carta)
	// jugada.Add(carta)
	// v1, _ = mano.Get(6)
	// carta, _ = v1.(Carta)
	// jugada.Add(carta)
	// c := doublylinkedlist.New()
	// v1, _ = t.Mazo.Get(0)
	// carta, _ = v1.(Carta)
	// v1, _ = t.Mazo.Get(1)
	// carta2, _ := v1.(Carta)
	// v1, _ = t.Mazo.Get(2)
	// carta3, _ := v1.(Carta)
	// c.Add(carta)
	// c.Add(carta2)
	// c.Add(carta3)
	// t.Combinaciones.PushBack(c)
	// c1 := doublylinkedlist.New()
	// c1.Add(carta2)
	// c1.Add(carta3)
	// c1.Add(carta)
	// t.Combinaciones.PushBack(c1)
	// mostrarTablero(t)

	// realizarJugada(&t, mano, 3, 0, jugada)
	// mostrarTablero(t)
	// mostrarMano(mano)

	// realizarJugada(&t, mano, 1, i, jugada)
	// mostrarTablero(t)
	// mostrarMano(mano)

	// realizarJugada(&t, mano, 2, i, jugada)
	// mostrarTablero(t)
	// mostrarMano(mano)

	/*
		//Codigo para probar la funcion TrioValido()
		jugada := doublylinkedlist.New()

		jugada.Add(Carta{0, 1, 1})

		jugada.Add(Carta{1, 2, 1})

		jugada.Add(Carta{1, 3, 1})

		jugada.Add(Carta{1, 4, 1})

		fmt.Println("LA JUGADA: ", jugada)

		fmt.Println("¿LA JUAGADA ES VALIDA?: ", TrioValido(jugada))
	*/
	/*
		//Codigo para probar la funcion EscaleraValida()
		jugada := doublylinkedlist.New()

		jugada.Add(Carta{0, 1, 1})

		jugada.Add(Carta{4, 1, 1})

		jugada.Add(Carta{5, 1, 1})

		jugada.Add(Carta{0, 1, 1})

		fmt.Println("LA JUGADA: ", jugada)

		fmt.Println("¿ES UNA ESCALERA VALIDA?: ", EscaleraValida(jugada))
	*/
}
