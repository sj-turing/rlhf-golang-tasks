package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

type ContactManager struct {
	contacts []Contact
}

func NewContactManager() *ContactManager {
	return &ContactManager{
		contacts: make([]Contact, 0),
	}
}

func (cm *ContactManager) AddContact(contact Contact) {
	cm.contacts = append(cm.contacts, contact)
}

type FilterFunc func(Contact) bool

func NameFilter(name string) FilterFunc {
	return func(c Contact) bool {
		return c.Name == name
	}
}

func PhoneFilter(phone string) FilterFunc {
	return func(c Contact) bool {
		return c.Phone == phone
	}
}

func EmailFilter(email string) FilterFunc {
	return func(c Contact) bool {
		return c.Email == email
	}
}

func (cm *ContactManager) FilterContacts(filters ...FilterFunc) []Contact {
	var result []Contact

	for _, contact := range cm.contacts {
		match := true
		for _, filter := range filters {
			if !filter(contact) {
				match = false
				break
			}
		}
		if match {
			result = append(result, contact)
		}
	}

	return result
}

func main() {
	cm := NewContactManager()
	cm.AddContact(Contact{Name: "John Doe", Phone: "123-456-7890", Email: "john@example.com"})
	cm.AddContact(Contact{Name: "John Doe", Phone: "882-456-7890", Email: "jdoe@example.com"})
	cm.AddContact(Contact{Name: "Jane Smith", Phone: "987-654-3210", Email: "jane@example.com"})
	cm.AddContact(Contact{Name: "Jane Smith", Phone: "123-654-3210", Email: "jsmith@example.com"})

	// Example: Filter by name
	filteredContacts := cm.FilterContacts(NameFilter("John Doe"))

	fmt.Println("filteredContacts after Name filter:", filteredContacts)

	// Example: Filter by name and email
	filteredContacts = cm.FilterContacts(NameFilter("Jane Smith"), EmailFilter("jane@example.com"))
	fmt.Println("filteredContacts after Name and Email filter:", filteredContacts)

	// Process filteredContacts
}
