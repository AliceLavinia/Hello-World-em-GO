package main
import "fmt"
import "reflect"

func main() {
	 nome := "Alice"
	 idade  := 22
	 versao  := 1.1
	fmt.Println("Ola senhora", nome," sua idade é",idade)
	fmt.Println("Este programa esta na versão", versao)

	fmt.Println("O tipo da variavel nome ", reflect.TypeOf(nome))
}