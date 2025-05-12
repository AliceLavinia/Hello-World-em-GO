package main
import "fmt"
import "os"
import "net/http"
// import "reflect"

func main() {
	
	exibeIntroducao()
	exibeMenu()

	// fmt.Println("O tipo da variavel nome ", reflect.TypeOf(nome))
	// fmt.Println("1- iniciar Monitoramento")
	// fmt.Println("2-Exibir os Logs")
	// fmt.Println("3- Sair do programa")
    // var comando int
	// //É uma função que recebe dois argumrntos, um modificador (valor de entrada) e um ponteiro que receberá o valor de entrada
	// fmt.Scan(&comando)
	// // o simbolo & siginifica o endereço da váriavel
	// fmt.Println("O endereço da variavel é",&comando)
	// fmt.Println("O comando escolhido foi ",comando)

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
    //     fmt.Println("Exibindo logs")
	// } else if comando== 3 {
    //     fmt.Println("Saindo do Programa")
	// } else {
	// 	fmt.Println("Não conheço esse comando")
	// }
	
    comando:=leComando()

	switch comando {
	   case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo logs")
	case 3:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Douglas"
    versao := 1.1
    fmt.Println("Olá, sr.", nome)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir os Logs")
	fmt.Println("3- Saindo do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi",comandoLido)
	
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
    site := "http://www.alura.com.br"
	// resp,err := http.Get(site)
}