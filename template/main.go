package main

import (
	"html/template"
	"os"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) IsOld() bool {
	return u.Age > 30
}

func main() {

	user := User{Name: "saechim", Email: "saechim@daeki.com", Age: 29}
	user2 := User{Name: "aaa", Email: "email@email.com", Age: 40}
	users := []User{user, user2}
	tmpl1, err := template.New("Tmpl1").ParseFiles("template/tmpl/tmpl.tmpl", "template/tmpl/tmpl2.tmpl")

	if err != nil {
		panic(err)
	}

	tmpl1.ExecuteTemplate(os.Stdout, "tmpl.tmpl", user)
	tmpl1.ExecuteTemplate(os.Stdout, "tmpl2.tmpl", users)
}
