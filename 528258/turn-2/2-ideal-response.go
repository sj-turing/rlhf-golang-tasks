package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

// contacts custom type which represent slice of Contact
type contacts []Contact

type ContactManager struct {
	contacts contacts
}

func NewContactManager() *ContactManager {
	return &ContactManager{
		contacts: make([]Contact, 0),
	}
}

func (cm *ContactManager) AddContact(contact Contact) {
	cm.contacts = append(cm.contacts, contact)
}

func (cm *ContactManager) GetContact() contacts {
	return cm.contacts
}

func (cs contacts) WithNameFilter(name string) contacts {
	result := make([]Contact, 0)
	for _, c := range cs {
		if c.Name == name {
			result = append(result, c)
		}
	}

	return result
}

func (cs contacts) WithPhoneFilter(phone string) contacts {
	result := make([]Contact, 0)
	for _, c := range cs {
		if c.Phone == phone {
			result = append(result, c)
		}
	}

	return result
}

func (cs contacts) WithEmailFilter(email string) contacts {
	result := make([]Contact, 0)
	for _, c := range cs {
		if c.Email == email {
			result = append(result, c)
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

	// Add filters one by one, chainable
	filteredContacts := cm.GetContact().
		WithNameFilter("John Doe").
		WithEmailFilter("john@example.com") // you can add more filters here by dot notations

	// Process filteredContacts
	fmt.Println("Filtered Contacts:", filteredContacts)
}
