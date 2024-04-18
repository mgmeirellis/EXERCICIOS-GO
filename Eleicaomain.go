func main() {
	defer fmt.Println("Encerrando votação. Obrigado por participar!")

	for {
		fmt.Println("\nMENU:")
		fmt.Println("1. Votar em um candidato")
		fmt.Println("2. Apuração parcial")
		fmt.Println("3. Conferência de votos")
		fmt.Println("0. Encerrar votação")
		fmt.Print("Escolha uma opção: ")

		var opcao int
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			votar()
		case 2:
			apuracaoParcial()
		case 3:
			conferirVotos()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Opção inválida. Por favor, escolha uma opção válida.")
		}
	}
}
