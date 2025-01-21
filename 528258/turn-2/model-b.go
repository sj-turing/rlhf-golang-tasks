package main

/*

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}
type FilterOption func(*FilterConfig)
type FilterConfig struct {
	NameFilter  func(Contact) bool
	PhoneFilter func(Contact) bool
	EmailFilter func(Contact) bool
}

func (f *FilterConfig) Apply(options ...FilterOption) {
	for _, option := range options {
		option(f)
	}
}
func NameFilter(name string) FilterOption {
	return func(f *FilterConfig) {
		f.NameFilter = func(c Contact) bool {
			return c.Name == name
		}
	}
}
func PhoneFilter(phone string) FilterOption {
	return func(f *FilterConfig) {
		f.PhoneFilter = func(c Contact) bool {
			return c.Phone == phone
		}
	}
}
func EmailFilter(email string) FilterOption {
	return func(f *FilterConfig) {
		f.EmailFilter = func(c Contact) bool {
			return c.Email == email
		}
	}
}
func (cm *ContactManager) FilterContacts(options ...FilterOption) []Contact {
	var config FilterConfig
	config.Apply(options...)
	var result []Contact
	for _, contact := range cm.contacts {
		match := true
		if config.NameFilter != nil && !config.NameFilter(contact) {
			match = false
		}
		if config.PhoneFilter != nil && !config.PhoneFilter(contact) {
			match = false
		}
		if config.EmailFilter != nil && !config.EmailFilter(contact) {
			match = false
		}
		if match {
			result = append(result, contact)
		}
	}
	return result
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
func main() {
	cm := NewContactManager()
	cm.AddContact(Contact{Name: "John Doe", Phone: "123-456-7890", Email: "john@example.com"})
	cm.AddContact(Contact{Name: "John Doe", Phone: "882-456-7890", Email: "jdoe@example.com"})
	cm.AddContact(Contact{Name: "Jane Smith", Phone: "987-654-3210", Email: "jane@example.com"})
	cm.AddContact(Contact{Name: "Jane Smith", Phone: "123-654-3210", Email: "jsmith@example.com"})
	// Example: Chainable filtering using options
	filteredContacts := cm.FilterContacts(
		NameFilter("John Doe"),
		EmailFilter("jdoe@example.com"),
	)
	fmt.Println("filteredContacts after Name and Email filter:", filteredContacts)
}
*/
