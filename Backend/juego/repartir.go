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
	fmt.Println("Escalesras joker")
	puntos := 0
	comb := doublylinkedlist.New()
	// ordenar la mano por palos de menor a mayor
	mano = SortStart(mano, 1)
	mostrarMano(mano)
	nuevoPalo := true
	hay_as := false
	ind_as := 0
	esc := false
	num_j := joker.Size()

	no_elim := -1
	if num_j > 0 {
		fmt.Println("Hay joker")
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
					fmt.Println("dos seguidas")
					fmt.Println(carta1)
					fmt.Println(carta2)
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
					fmt.Println("rey-as")
					fmt.Println(carta1)
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
					fmt.Println("añadir joker a ", carta1)
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
				fmt.Println("numer joker añ", num_j_anyadidos)
				fmt.Println("num joker quedan", num_j)
				// si el numero de cartas seguidas ha sido >=3, escalera valida
				puntos += puntos_t
				// añado l a la combinación a devolver
				comb.Add(l)
				comb.Each((func(index int, value interface{}) {
					fmt.Printf("cl %d: %v\n", index, value)
				}))
				fmt.Printf("BORRAR %d  %d\n", i, i_inf)
				if !mirar_j {
					// si no hay AS, borro de la mano las cartas de los indices seguidos que correspondan
					k := no_elim % 100
					no_elim = no_elim / 100
					for j := i; j >= i_inf; j-- {
						fmt.Printf("A BORRAR %d\n", j)
						fmt.Printf("NO BORRAR %d\n", k)
						if j != k {
							mano.Remove(j)
							fmt.Printf("BORRARDO %d\n", j)
							mostrarMano(mano)
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
						fmt.Printf("A BORRAR %d\n", j)
						fmt.Printf("NO BORRAR %d\n", k)
						if j != k {
							mano.Remove(j)
							fmt.Printf("BORRARDO %d\n", j)
							mostrarMano(mano)
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
					fmt.Printf("BORRARDO %d\n", ind_as)
					mostrarMano(mano)
				}
				fmt.Printf("JOKERS QUEDAN")
				mostrarMano(joker)
				fmt.Printf("BORRAR JOKER 0 %d\n", num_j_anyadidos)
				for j := 0; j < num_j_anyadidos; j++ {
					joker.Remove(0)
					mostrarMano(joker)
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
	mostrarMano(mano)
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
			comb.Each((func(index int, value interface{}) {
				fmt.Printf("cl %d: %v\n", index, value)
			}))
			fmt.Printf("BORRAR %d  %d\n", i, i_inf)
			if !borrar_as {
				// si no hay AS, borro de la mano las cartas de los indices seguidos que correspondan
				k := no_elim % 100
				no_elim = no_elim / 100
				for j := i; j >= i_inf; j-- {
					fmt.Printf("A BORRAR %d\n", j)
					fmt.Printf("NO BORRAR %d\n", k)
					if j != k {
						mano.Remove(j)
						fmt.Printf("BORRARDO %d\n", j)
						mostrarMano(mano)
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
					fmt.Printf("A BORRAR %d\n", j)
					fmt.Printf("NO BORRAR %d\n", k)
					if j != k {
						mano.Remove(j)
						fmt.Printf("BORRARDO %d\n", j)
						mostrarMano(mano)
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
				fmt.Printf("BORRARDO %d\n", ind_as)
				mostrarMano(mano)
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
	fmt.Println("Calcular trios con joker")
	fmt.Println("Lista de joker")
	mostrarMano(joker)
	puntos := 0
	mano = SortStart(mano, 0)
	comb := doublylinkedlist.New()
	fmt.Println("Mano")
	mostrarMano(mano)
	fmt.Println("Empezar")
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
				fmt.Println("carta: ", carta1)
				fmt.Println("carta2: ", carta2)
				if carta1.Palo != carta2.Palo {
					// las tres cartas son de distinto palo
					trio = true
					fmt.Println("carta: ", carta1, " ok")
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
						fmt.Printf("BORRARDO %d\n", j)
						mostrarMano(mano)
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
	mostrarMano(mano)
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
			fmt.Println("carta: ", carta1)
			fmt.Println("carta2: ", carta2)
			fmt.Println("carta3: ", carta3)
			if carta1.Palo != carta2.Palo && carta2.Palo != carta3.Palo && carta1.Palo != carta3.Palo {
				// las tres cartas son de distinto palo
				trio = true
				fmt.Println("carta: ", carta1, " ok")
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
					fmt.Println("carta4: ", carta4)
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
					fmt.Printf("BORRARDO %d\n", j)
					mostrarMano(mano)
				}
				comb.Add(l)
			}
		}
	}
	return puntos, comb, trio
}

func separarJokers(mano *doublylinkedlist.List) (*doublylinkedlist.List, *doublylinkedlist.List) {
	fmt.Println("Separar jokers")
	mano = SortStart(mano, 0)
	joker := doublylinkedlist.New()
	mostrarMano(mano)
	hay_j := true
	for hay_j {
		v, _ := mano.Get(mano.Size() - 1)
		carta, _ := v.(Carta)
		fmt.Println("mirar joker ", carta)
		if carta.Valor == 0 {
			joker.Add(carta)
			mano.Remove(mano.Size() - 1)
		} else {
			hay_j = false
		}
	}
	return mano, joker
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

func abrir(jugada *doublylinkedlist.List, mano *doublylinkedlist.List, t *tablero) { //falta comprobar trios y escaleras
	for i := 0; i <= jugada.Size(); i++ {
		v1, _ := jugada.Get(i)
		carta, _ := v1.(Carta)

		ind := mano.IndexOf(carta)
		mano.Remove(ind)
		fmt.Println("carta eliminada", carta)

	}

}

// función para añadir una carta a una combinación
func anyadirCarta(jugada *doublylinkedlist.List, mano *doublylinkedlist.List, t *tablero, idCombinacion int) {
	if !jugada.Empty() {
		v1, _ := jugada.Get(0)
		carta, _ := v1.(Carta)

		i := 0
		for e := t.Combinaciones.Front(); e != nil; e = e.Next() {
			if i == idCombinacion {
				listaC := e.Value.(*doublylinkedlist.List)
				listaC.Add(carta)
				//falta comprobar trio y escalera
				t.Combinaciones.Remove(e)
				t.Combinaciones.PushBack(listaC)
				ind := mano.IndexOf(carta)
				mano.Remove(ind)
				return
			}
			i++
		}
	}
}

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
		return
	case 3: //Añadir 1 carta a una combinación existente
		anyadirCarta(cartasAjugar, mano, t, 0)
		return
	default:
	}
}

func main() {
	fmt.Println("Hola1")
	rand.Seed(time.Now().UnixNano())
	mazo := doublylinkedlist.New()
	//descarte := doublylinkedlist.New()
	//i := 4
	fmt.Println("Hola2")

	creacionBaraja(mazo)
	fmt.Println("Hola3")

	mano := repartirMano(mazo)
	fmt.Println("Hola4")

	mostrarMano(mano)
	fmt.Println("Hola5")

	// robarCarta(mazo, mano)
	// fmt.Println("MANO CARTA ROBADA")
	// mostrarMano(mano)

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

}
