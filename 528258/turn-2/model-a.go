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

type FilterBuilder struct {
	filters []FilterFunc
}

func NewFilterBuilder() *FilterBuilder {
	return &FilterBuilder{
		filters: make([]FilterFunc, 0),
	}
}

func (fb *FilterBuilder) AddFilter(filter FilterFunc) *FilterBuilder {
	fb.filters = append(fb.filters, filter)
	return fb
}

func (fb *FilterBuilder) Apply(contacts []Contact) []Contact {
	var result []Contact

	for _, contact := range contacts {
		match := true
		for _, filter := range fb.filters {
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

func main() {
	cm := NewContactManager()
	cm.AddContact(Contact{Name: "John Doe", Phone: "123-456-7890", Email: "john@example.com"})
	cm.AddContact(Contact{Name: "John Doe", Phone: "882-456-7890", Email: "jdoe@example.com"})
	cm.AddContact(Contact{Name: "Jane Smith", Phone: "987-654-3210", Email: "jane@example.com"})
	cm.AddContact(Contact{Name: "Jane Smith", Phone: "123-654-3210", Email: "jsmith@example.com"})

	// Create a filter builder
	filterBuilder := NewFilterBuilder()

	// Add filters one by one, chainable
	filterBuilder.AddFilter(NameFilter("John Doe")).AddFilter(EmailFilter("john@example.com"))

	// Apply filters to get the filtered contacts
	filteredContacts := filterBuilder.Apply(cm.contacts)

	// Process filteredContacts
	fmt.Println("Filtered Contacts:", filteredContacts)
}
