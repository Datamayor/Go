package main

import (
	"errors"
	"fmt"
)

type Todo struct {
	Title  string
	Status string
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:  title,
		Status: "not_done",
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {

	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) markCompleted(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Status = "completed"
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	for i, t := range *todos {
		fmt.Printf("[%d] %s — %s\n", i, t.Title, t.Status)
	}
}
