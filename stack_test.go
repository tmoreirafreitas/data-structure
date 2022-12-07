package main

import (
	"testing"

	. "github.com/tmoreirafreitas/data-structure/generics"
	. "github.com/tmoreirafreitas/data-structure/model"
)

var stack = new(Stack[Produto])

func TestStackEmpty(t *testing.T) {
	resultado := stack.IsEmpty()
	esperado := true

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}
func TestPush(t *testing.T) {
	macarrao := new(Produto)
	macarrao.Categoria = "Massas"
	macarrao.Nome = "Macarrão"
	macarrao.Preco = 3.77

	err := stack.Push(macarrao)

	if err != nil {
		t.Errorf(err.Error())
	}

	resultado := int(stack.Count())
	esperado := 1

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestPop(t *testing.T) {
	_, err := stack.Pop()
	if err != nil {
		t.Errorf(err.Error())
	}
}
