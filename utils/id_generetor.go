package utils

import "math/rand" // Permite sortear números aleatórios

func GerarIdAleatorio() int {

	const caracteres = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Aqui estão todos os símbolos possíveis para o ID: letras minúsculas, maiúsculas e números

	n := 5                // O ID terá 5 caracteres
	id := make([]byte, n) // Cria uma “caixa” para guardar 8 símbolos, cada um é um byte

	for i := 0; i < n; i++ { // Para cada posição (de 0 até 4) do ID...
		id[i] = caracteres[rand.Intn(len(caracteres))]
		// - rand.Intn(len(letras)) sorteia um número de 0 até a quantidade de letras e números disponíveis.
		// - letras[...] pega o caractere nessa posição sorteada da string 'letras'.
		// - id[i] = ... coloca esse caractere sorteado na posição i do slice 'id', montando o ID final passo a passo.
	}

	return int() // Junta todos os bytes gerados e transforma em texto antes de retornar
}
