package main
import "fmt"
// import "reflect"

func main() {
	 nome := "Alice"
	 idade  := 22
	 versao  := 1.1
	fmt.Println("Ola senhora", nome," sua idade é",idade)
	fmt.Println("Este programa esta na versão", versao)

	// fmt.Println("O tipo da variavel nome ", reflect.TypeOf(nome))
	fmt.Println("1- iniciar Monitoramento")
	fmt.Println("2-Exibir os Logs")
	fmt.Println("3- Sair do programa")
    var comando int
	//É uma função que recebe dois argumrntos, um modificador (valor de entrada) e um ponteiro que receberá o valor de entrada
	fmt.Scanf("%d",&comando)
	// o simbolo & siginifica o endereço da váriavel
	fmt.Println("O endereço da variavel é",&comando)
	fmt.Println("O comando escolhido foi ",comando)
	
}