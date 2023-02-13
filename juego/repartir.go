package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/lists/doublylinkedlist"
)

type Carta struct { //Struct utilizado para definir la estructura de datos que representa las cartas
	Valor int
	Palo  int
	Color int
}

func compararCartasN(a Carta, b Carta) int { //Función parte del sort encargada de filtrar las cartas por valor y color
	if a.Valor < b.Valor {
		return 1
	} else if a.Valor > b.Valor {
		return -1
	} else {
		if a.Color < b.Color {
			return 1
		} else if a.Color > b.Color {
			return -1
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
			for k := 1; k <= 12; k++ {
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

func calcularEscaleras(mano *doublylinkedlist.List) int { //Función encargada de calcular las diferentes escaleras posibles en la mano (y así obtener los puntos)
	puntos := 0
	mano = SortStart(mano, 1)
	mostrarMano(mano)
	nuevoPalo := true
	hay_as := false
	for i := 0; i < mano.Size()-1; i++ {
		num_c := 1
		puntos_t := 0
		v1, _ := mano.Get(i)
		carta1, _ := v1.(Carta)
		if nuevoPalo {
			hay_as = carta1.Valor == 1
		}
		if carta1.Valor >= 10 {
			puntos_t = puntos_t + 10
		} else {
			puntos_t = puntos_t + carta1.Valor
		}
		hay_esc := true
		for hay_esc {
			v2, _ := mano.Get(i + 1)
			carta2, _ := v2.(Carta)
			if carta1.Palo == carta2.Palo {
				nuevoPalo = false
			} else {
				nuevoPalo = true
			}
			if carta1.Valor+1 == carta2.Valor && carta1.Palo == carta2.Palo {
				fmt.Println("carta1: ", carta1)
				fmt.Println("carta2: ", carta2)
				if carta2.Valor >= 10 {
					puntos_t = puntos_t + 10
				} else if carta1.Valor == 12 && carta2.Valor == 1 {
					puntos_t += 11 //contains
				} else {
					puntos_t = puntos_t + carta2.Valor
				}
				num_c += 1
				i++
			} else if num_c >= 2 && carta1.Valor == 12 && hay_as {
				fmt.Println("carta1: ", carta1)
				fmt.Println("carta2: ", carta2)
				puntos_t = puntos_t + 11
				num_c += 1
				hay_esc = false
			} else if carta1.Valor == carta2.Valor {
				i++
			} else {
				hay_esc = false
			}
			carta1 = carta2
		}
		if num_c >= 3 {
			puntos += puntos_t
		}
	}
	return puntos
}

func calcularTrios(mano *doublylinkedlist.List) int { //Función encargada de calcular los puntos de los posibles trios de las barajas
	puntos := 0
	mano = SortStart(mano, 0)
	mostrarMano(mano)
	for i := 0; i < mano.Size()-2; i++ {
		palo := 0
		v1, _ := mano.Get(i)
		carta1, _ := v1.(Carta)
		v2, _ := mano.Get(i + 1)
		carta2, _ := v2.(Carta)
		v3, _ := mano.Get(i + 2)
		carta3, _ := v3.(Carta)
		if carta1.Valor == carta2.Valor && carta2.Valor == carta3.Valor {
			fmt.Println("carta: ", carta1)
			fmt.Println("carta2: ", carta2)
			fmt.Println("carta3: ", carta3)
			if carta1.Palo != carta2.Palo && carta2.Palo != carta3.Palo && carta1.Palo != carta3.Palo {
				fmt.Println("carta: ", carta1, " ok")
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
				if carta1.Valor == carta4.Valor && palo == 10 {
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
			}
		}
	}
	return puntos
}

func calcularPuntosPosibles(mano *doublylinkedlist.List) int { //Función encargada de revisar los puntos posibles de una mano
	puntos := 0
	puntos += calcularTrios(mano)
	puntos += calcularEscaleras(mano) //Revisar calcular puntos con cartas ya utilizadas

	return puntos
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

func main() {
	fmt.Println("Hola1")
	rand.Seed(time.Now().UnixNano())
	mazo := doublylinkedlist.New()
	descarte := doublylinkedlist.New()
	i := 4
	fmt.Println("Hola2")

	creacionBaraja(mazo)
	fmt.Println("Hola3")

	mano := repartirMano(mazo)
	fmt.Println("Hola4")

	mostrarMano(mano)
	fmt.Println("Hola5")

	robarCarta(mazo, mano)
	fmt.Println("MANO CARTA ROBADA")
	mostrarMano(mano)

	fmt.Println("Puntos ", calcularPuntosPosibles(mano)) //Revisar as
	fmt.Println(descarte)

	finTurno(mazo, mano, descarte, i)
	fmt.Println("MANO DESCARTE HECHO")
	fmt.Println(descarte)
	mostrarMano(mano)

	robarCarta(mazo, mano)
	fmt.Println("MANO CARTA ROBADA")
	mostrarMano(mano)

	fmt.Println("Puntos ", calcularPuntosPosibles(mano)) //Revisar as
	fmt.Println(descarte)

	finTurno(mazo, mano, descarte, i)
	fmt.Println("MANO DESCARTE HECHO")
	fmt.Println(descarte)
	mostrarMano(mano)

}
