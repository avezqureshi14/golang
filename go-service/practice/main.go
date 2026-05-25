package main

import "fmt"

type Task struct {
	Id   int
	Name string
}

type Todos struct {
	tasks  []Task
	nextId int
}

func (t *Todos) addTodo(task string) {
	t.tasks = append(t.tasks, Task{Id: t.nextId, Name: task})
	t.nextId++
}

func (t *Todos) getTodos() {
	fmt.Println(t.tasks)
}

func (t *Todos) deleteTodo(id int) {
	for i, task := range t.tasks {
		if task.Id == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
		}
	}
}

func (t *Todos) updateTodo(id int, name string) {
	for _, task := range t.tasks {
		if task.Id == id {
			task.Name = name
		}
	}
}

func main() {
	myt := Todos{}
	myt.addTodo("Avez")
	myt.addTodo("Bhai")
	myt.getTodos()
	myt.deleteTodo(1)
	myt.getTodos()
}
