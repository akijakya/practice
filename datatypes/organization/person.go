package organization

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Citizen interface {
	Identifiable
	Country() string
}

type socialSecurityNumber string

func NewSocialSecurityNumber(value string) Citizen {
	return socialSecurityNumber(value) // type casting
}

func (ssn socialSecurityNumber) ID() string {
	return string(ssn)
}

func (ssn socialSecurityNumber) Country() string {
	return "United States of America"
}

type euIdentifier struct {
	id      string
	country []string
}

func NewEuIdentifier(id interface{}, country string) Citizen {
	switch v := id.(type) {
	case string:
		return euIdentifier{
			id:      v,
			country: []string{country},
		}
	case int:
		return euIdentifier{
			id:      strconv.Itoa(v),
			country: []string{country},
		}
	case euIdentifier:
		return v
	case Person:
		euID, _ := v.Citizen.(euIdentifier)
		return euID
	default:
		panic("Using an invalid típe to initialize EU Identifier")
	}
}

func (eui euIdentifier) ID() string {
	return fmt.Sprintf("%v", eui.id)
}

func (eui euIdentifier) Country() string {
	return fmt.Sprintf("EU: %v", eui.country)
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {
	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

type Name struct {
	first string
	last  string
}

func (n Name) FullName() string {
	return fmt.Sprintf("%s %s", n.first, n.last)
}

type Employee struct {
	Name
}

type Person struct {
	Name
	twitterHandler TwitterHandler
	Citizen
}

func NewPerson(firstName, lastName string, citizen Citizen) Person {
	return Person{
		Name: Name{
			first: firstName,
			last:  lastName,
		},
		Citizen: citizen,
	}
}

// func (p *Person) ID() string {
// 	return fmt.Sprintf("Person's identifier: %s", p.Citizen.ID())
// }

func (p *Person) SetTwitterHandler(handler TwitterHandler) error {
	if len(handler) == 0 {
		p.twitterHandler = handler
	} else if !strings.HasPrefix(string(handler), "@") {
		return errors.New("twitter handler must start with an @ symbol")
	}

	p.twitterHandler = handler
	return nil
}

func (p *Person) TwitterHandler() TwitterHandler {
	return p.twitterHandler
}
