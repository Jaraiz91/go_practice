package main

import "fmt"

type Course struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Course
}

type User struct {
	nombre string
	email  string
	activo bool
}

type Student struct {
	user   User
	codigo string
}

func main() {
	//Relacion uno a uno
	javier := User{
		nombre: "Javier",
		email:  "jaraiz37@gmail.com",
		activo: true,
	}

	isabel := User{
		nombre: "Isabel",
		email:  "iromerojordano@gmail.com",
		activo: true,
	}
	javierstudent := Student{
		user:   javier,
		codigo: "001asd",
	}
	fmt.Println(isabel)
	fmt.Println(javier)
	fmt.Println(javierstudent)

	//Relacion de uno a muchos
	video1 := Video{
		titulo: "01-Introducción",
	}
	video2 := Video{
		titulo: "02-Instalación",
	}
	curso1 := Course{
		titulo: "curso de Go",
		videos: []Video{video1, video2},
	}
	video1.curso = curso1
	video2.curso = curso1
	fmt.Println(video1.curso.titulo)

	for _, video := range curso1.videos {
		fmt.Println(video.titulo)
	}

}
