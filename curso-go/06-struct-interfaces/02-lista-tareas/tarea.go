package main

import "fmt"

// Lista de tareas
type TaskList struct {
	tasks []*Task
}

func (tl *TaskList) showlist() {
	for i, task := range tl.tasks {
		fmt.Println(i, task.name)
	}
}

func (tl *TaskList) appendTask(t *Task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *TaskList) deleteTask(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

// tareas
type Task struct {
	name      string
	desc      string
	completed bool
}

func (t *Task) toPrint() {
	fmt.Printf("Nombre: %s\nDescripción: %s\nCompletado: %t\n", t.name, t.desc, t.completed)
}

func (t *Task) markComplete() {
	t.completed = true
}

func main() {
	tarea1 := Task{name: "Curso de Go",
		desc:      "completar curso de go",
		completed: false,
	}
	tarea2 := Task{name: "curso Python",
		desc:      "Completar curso de Python",
		completed: true,
	}
	lista := TaskList{}
	lista.appendTask(&tarea1)
	lista.appendTask(&tarea2)

	fmt.Println(lista)
	for i, task := range lista.tasks {
		if task.completed == true {
			fmt.Printf("Task %s is completed, removing from the list!!\n", task.name)
			lista.deleteTask(i)
			fmt.Println("Updated List:")
			lista.showlist()

		}
	}

}
