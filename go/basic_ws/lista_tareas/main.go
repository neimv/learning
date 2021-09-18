package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	fmt.Println("Hello")
	t1 := &task{
		nombre:      "Completar mi curso de go",
		descripcion: "completar mi curso de go de platzi en esta semana",
	}
	t2 := &task{
		nombre:      "Completar mi curso de python",
		descripcion: "completar mi curso de python de platzi en esta semana",
	}
	t3 := &task{
		nombre:      "Completar mi curso de js",
		descripcion: "completar mi curso de js de platzi en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])

	lista.agregarALista(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i)
		fmt.Println("nombre", lista.tasks[i].nombre)
	}

	for index, t := range lista.tasks {
		fmt.Println("Index", index, "nombre", t.nombre)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("Imprimiendo lista")
	lista.imprimirLista()

	lista.tasks[0].marcarCompleta()
	fmt.Println("Tareas completadas")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["nombre"] = lista
	t4 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "completar mi curso de go de platzi en esta semana",
	}
	t5 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "completar mi curso de python de platzi en esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["otro"] = lista2

	fmt.Println("Tareas de nombre")
	mapaTareas["nombre"].imprimirLista()
	fmt.Println("Tareas de otro")
	mapaTareas["otro"].imprimirLista()
}
