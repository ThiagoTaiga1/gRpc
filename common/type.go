package common

import "fmt"

//struct de estudante
type Student struct {
	ID                  int
	FirstName, LastName string
}

//retornando nome completo do estudante
func (s Student) FullName() string {

	return s.FirstName + " " + s.LastName
}

// struct representando um college
type College struct {
	database map[int]Student
}

// Adicionar métodos adiciona um aluno à faculdade (procedimento).
func (c *College) Add(payload Student, reply *Student) error {

	// verifica se o aluno já existe no banco de dados
	if _, ok := c.database[payload.ID]; ok {
		return fmt.Errorf("Student with id '%d'already exists", payload.ID)
	}
	// adiciona o studante `p`no banco de dados
	c.database[payload.ID] = payload
	// seta o valor replay
	*reply = payload
	// retorna null se der algum erro
	return nil

}

// Get métodos retorna um aluno com id específico (procedimento).
func (c *College) Get(payload int, replay *Student) error {
	// busca um student e o id na base de dados.
	result, ok := c.database[payload]
	// / verifica se o aluno já existe no banco de dados
	if !ok {
		return fmt.Errorf("student with id '%d' does not exist", payload)
	}
	*replay = result

	return nil
}

// A função NewCollege retorna uma nova instância de College (ponteiro).
func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
