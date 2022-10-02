package main

import (
	"datatypes/organization"
	"fmt"
)

func main() {
	// p := organization.NewPerson("James", "Wilson", organization.NewSocialSecurityNumber("123-45-6789"))
	p := organization.NewPerson("James", "Wilson", organization.NewEuIdentifier("123456789", "Hungary"))
	// p := organization.NewPerson("James", "Wilson", organization.NewEuIdentifier(123456789, "Hungary"))
	err := p.SetTwitterHandler("@jam_wils")
	if err != nil {
		fmt.Printf("An error occurred setting twitter handler: %s", err.Error())
	}

	// fmt.Printf("%T\n", organization.TwitterHandler("test"))
	println(p.ID())
	println(p.Country())
	// println(p.FullName())
	// println(p.TwitterHandler())
	// println(p.TwitterHandler().RedirectUrl())

	// name1 := Name{First: "James", Last: "Wilson", Middle: []string{"Jonathan"}}
	// otherName1 := Name{First: "James", Last: "Wilson", Middle: []string{"Jonathan"}}
	// name2 := OtherName{First: "James", Last: "Wilson"}

	// if name1 == (Name{}) {
	// 	println("We match!")
	// }

	// println(name1.Equals(otherName1))

	// ssn := organization.NewSocialSecurityNumber("123-45-6789") // organization.socialSecurityNumber
	// eu := organization.NewEuIdentifier("123456789", "Hungary") // organization.euIdentifier
	// portfolio := map[Name][]organization.Person{}
	// portfolio[name1] = []organization.Person{p}

	// if ssn == eu {
	// 	println("We match!")
	// }

	// fmt.Printf("%T\n", ssn)
	// fmt.Printf("%T\n", eu)
}

type Name struct {
	First  string
	Last   string
	Middle []string
}

func (n Name) Equals(otherName Name) bool {
	return n.First == otherName.First && n.Last == otherName.Last && len(n.Middle) == len(otherName.Middle)
}

type OtherName struct {
	First string
	Last  string
}
