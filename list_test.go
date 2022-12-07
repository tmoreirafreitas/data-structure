package main

import (
	"fmt"
	"testing"

	. "github.com/tmoreirafreitas/data-structure/generics"
	. "github.com/tmoreirafreitas/data-structure/model"
)

var lista = new(List[Produto])

func carregarListaDeProdutos() {
	if lista.Front() == nil {
		macarrao := new(Produto)
		macarrao.Categoria = "Massas"
		macarrao.Nome = "Macarrão"
		macarrao.Preco = 3.77

		pizza := Produto{
			Categoria: "Massas",
			Nome:      "Pizza",
			Preco:     15.78,
		}

		oleo := Produto{
			Categoria: "Olivas",
			Nome:      "Olea Liza",
			Preco:     7.73,
		}

		lista.Insert(macarrao)
		lista.Insert(&pizza)
		lista.InsertStart(&oleo)

		fmt.Println(lista.Count())
	}
}

func TestIsEmptyList(t *testing.T) {
	resultado := lista.IsEmpty()
	esperado := true

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}
func TestAddElementInList(t *testing.T) {
	carregarListaDeProdutos()
	resultado := int(lista.Count())
	esperado := 3

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestContainsElementInList(t *testing.T) {
	carregarListaDeProdutos()
	resultado := lista.Contains("Nome", "Macarrão")
	esperado := true

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestGetByIndexFromElementInList(t *testing.T) {
	pizza, err := lista.GetByIndex(2)
	if err != nil {
		t.Errorf(err.Error())
	}

	resultado := pizza.Nome
	esperado := "Pizza"

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}

func TestRemoveElementInList(t *testing.T) {
	oleo := Produto{
		Categoria: "Olivas",
		Nome:      "Olea Liza",
		Preco:     7.73,
	}
	lista.Remove(&oleo)

	resultado := lista.Contains("Nome", "Olea Liza")
	esperado := false

	if resultado != esperado {
		t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
	}
}
