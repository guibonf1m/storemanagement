package entity

// O pacote entity geralmente representa o núcleo dos dados do domínio da aplicação
// Ela define a estrutura principal com os campos essenciais e é usada em todas as
// outras camadas (repository, service e handlers) para garantir a consistência e
//
// tipagem forte.
type Produto struct {
	ID         int
	Nome       string
	Quantidade int
	Preco      float64
}
