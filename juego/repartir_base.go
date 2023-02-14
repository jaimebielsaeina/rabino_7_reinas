package main

import "github.com/emirpasic/gods/lists/doublylinkedlist"

import(
	"fmt"
	"math/rand"
	"time"
	
)

type Carta struct{
	Valor int
	Palo int
	Color int
}

func creacionBaraja(list *doublylinkedlist.List){
	carta := Carta{0,0,0}
	for  i := 1; i <= 2; i++ {
		carta.Color = i;
		for j := 1; j <= 4; j++ {
			carta.Palo = j;
			for k := 1; k <= 12; k++{
				carta.Valor = k;
				list.Add(carta);
			}
		}
	}
	carta.Valor = 0;
	carta.Color = 1;
	list.Add(carta);
	carta.Color = 2;
	list.Add(carta);
}

func repartirMano(list *doublylinkedlist.List)*doublylinkedlist.List{
	i := 98
	listR :=  doublylinkedlist.New()
	for j := 0; j < 14; j++{
		r := rand.Intn(i) + 1
		value, ok := list.Get(r)
		if !ok {
			fmt.Println("Lista no contiene el valor" ,r)
		}else{
			listR.Add(value)
		}
		
	}

	return listR
}

func mostrarMano(mano *doublylinkedlist.List){
	mano.Each(func(index int, value interface{}) {
		fmt.Printf("%d: %v\n", index, value)
	})
}

func main2(){
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

}