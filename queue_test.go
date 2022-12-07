package main

import (
	"testing"

	. "github.com/tmoreirafreitas/data-structure/generics"
	. "github.com/tmoreirafreitas/data-structure/model"
)

var queue = new(Queue[Produto])

func TestQueueEmpty(t *testing.T) {
	resultado := queue.IsEmpty()
	esperado := true

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestEnqueue(t *testing.T) {
	macarrao := new(Produto)
	macarrao.Categoria = "Massas"
	macarrao.Nome = "Macarrão"
	macarrao.Preco = 3.77

	queue.Queue(macarrao)

	resultado := int(queue.Count())
	esperado := 1

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestDequeue(t *testing.T) {
	_, err := queue.Dequeue()
	if err != nil {
		t.Errorf(err.Error())
	}
}
