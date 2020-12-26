package main

import "fmt"

const INITILIAZE = 0

var database = make([]User, INITILIAZE)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UsersRepository interface {
	create(id int, name, email, pwd string) User
	destroy(id int)
}

type UsersRepositoryImpl struct{}

func (r *UsersRepositoryImpl) create(id int, name, email, pwd string) User {
	u := User{
		Id:       id,
		Name:     name,
		Email:    email,
		Password: pwd,
	}

	database = append(database, u)
	return u
}

func (r *UsersRepositoryImpl) destroy(id int) {
	var copyDatabase = make([]User, 0)
	for _, user := range database {
		if user.Id != id {
			copyDatabase = append(copyDatabase, user)
		}
	}

	database = copyDatabase
}

func Run(r UsersRepository) {
	u1 := r.create(1, "Thalyson", "tha.motog@gmail.com", "12345")
	u2 := r.create(2, "Maria", "maria@gmail.com", "12345")
	u3 := r.create(3, "João", "joao@gmail.com", "12345")

	fmt.Printf("User added: %s", u1.Name)
	fmt.Printf("User added: %s", u2.Name)
	fmt.Printf("User added: %s", u3.Name)
	fmt.Println()
	fmt.Println(database)

	r.destroy(3)
	fmt.Printf("User removed: %s", u3.Name)
	fmt.Println()
	fmt.Println(database)
}

func main() {
	impl := UsersRepositoryImpl{}
	Run(&impl)
}
