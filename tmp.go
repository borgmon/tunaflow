package main

type User struct {
	id int
}

func (u *User) getID() int {
	return u.id
}

type Admin struct {
	User
	name string
}

func (a *Admin) getName() string {
	return a.name
}

type Guest struct {
	User
}

func (g *Guest) getName() string {
	return "Guest"
}

type ActiveUser interface {
	getName() string
}

func getNameFromUser(a ActiveUser) string {
	return a.getName()
}

func main() {
	admin := &Admin{User{1}, "Steve"}
	guest := &Guest{User{2}}

	println(getNameFromUser(admin))
	println(getNameFromUser(guest))
	println(admin.getID())
	println(guest.getID())

}
