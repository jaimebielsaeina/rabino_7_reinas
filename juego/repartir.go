package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/emirpasic/gods/lists/doublylinkedlist"
)

type Carta struct {
	Valor int
	Palo  int
	Color int
}

func compararCartas(a Carta, b Carta) int {
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

func creacionBaraja(list *doublylinkedlist.List) {
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
	// carta.Valor = 0
	// carta.Color = 1
	// list.Add(carta)
	// carta.Color = 2
	// list.Add(carta)
}

func repartirMano(list *doublylinkedlist.List) *doublylinkedlist.List {
	i := 98
	listR := doublylinkedlist.New()
	for j := 0; j < 14; j++ {
		r := rand.Intn(i) + 1
		value, ok := list.Get(r)
		for !ok {
			fmt.Println("Lista no contiene el valor", r)
			r = rand.Intn(i) + 1
			value, ok = list.Get(r)
		}
		listR.Add(value)
		list.Remove(r)

	}

	return listR
}

func mostrarMano(mano *doublylinkedlist.List) {
	mano.Each(func(index int, value interface{}) {
		fmt.Printf("%d: %v\n", index, value)
	})
}

func calcularTrios(mano *doublylinkedlist.List) int {
	puntos := 0
	mano = SortStartNum(mano)
	mostrarMano(mano)
	for i := 0; i < 14-2; i++ {
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

func calcularPuntosPosibles(mano *doublylinkedlist.List) int {
	puntos := 0
	puntos += calcularTrios(mano)

	// mano.Each(func(index int, value interface{}) {

	return puntos
}

func partition(mano *doublylinkedlist.List, low, high int) (*doublylinkedlist.List, int) {
	v1, _ := mano.Get(high)
	carta1, _ := v1.(Carta)
	i := low
	for j := low; j < high; j++ {
		v2, _ := mano.Get(j)
		carta2, _ := v2.(Carta)
		if compararCartas(carta1, carta2) == -1 {
			mano.Swap(i, j)
			i++
		}
	}
	mano.Swap(i, high)
	return mano, i
}

func SortNum(mano *doublylinkedlist.List, low, high int) *doublylinkedlist.List {
	if low < high {
		var p int
		mano, p = partition(mano, low, high)
		mano = SortNum(mano, low, p-1)
		mano = SortNum(mano, p+1, high)
	}
	return mano
}

func SortStartNum(mano *doublylinkedlist.List) *doublylinkedlist.List {
	return SortNum(mano, 0, mano.Size()-1)
}

func main() {
	fmt.Println("Hola1")
	rand.Seed(time.Now().UnixNano())
	list := doublylinkedlist.New()
	fmt.Println("Hola2")

	creacionBaraja(list)
	fmt.Println("Hola3")

	mano := repartirMano(list)
	fmt.Println("Hola4")

	mostrarMano(mano)
	fmt.Println("Hola5")
	fmt.Println("Puntos ", calcularPuntosPosibles(mano))

}
