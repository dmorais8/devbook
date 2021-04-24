package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario reprensenta um usuario da rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

//Preparar vai chama os metodos para validar os valores da struct
func (usuario *Usuario) Preparar() error {

	if erro := usuario.validar(); erro != nil {
		return erro
	}

	usuario.formatar()
	return nil
}

func (usuario *Usuario) validar() error {

	if usuario.Nome == "" {
		return errors.New("o nome é obrigatorio e nao pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("o nick é obrigatorio e nao pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o email é obrigatorio e nao pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("a senha é obrigatorio e nao pode estar em branco")
	}

	return nil

}

func (usuario *Usuario) formatar() {

	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

}