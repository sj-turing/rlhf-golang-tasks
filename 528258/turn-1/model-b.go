package main

import (
	"fmt"
	"strings"
)

type Contact struct {
	Name  string
	Phone string
	Email string
}

type contactOption func(*contactFilter)

type contactFilter struct {
	name    []string
	phone   []string
	email   []string
	options []contactOption
}

func (f *contactFilter) apply() func(Contact) bool {
	return func(contact Contact) bool {
		if len(f.name) > 0 {
			match := false
			for _, name := range f.name {
				if strings.Contains(contact.Name, name) {
					match = true
					break
				}
			}
			if !match {
				return false
			}
		}

		if len(f.phone) > 0 {
			match := false
			for _, phone := range f.phone {
				if strings.Contains(contact.Phone, phone) {
					match = true
					break
				}
			}
			if !match {
				return false
			}
		}

		if len(f.email) > 0 {
			match := false
			for _, email := range f.email {
				if strings.Contains(contact.Email, email) {
					match = true
					break
				}
			}
			if !match {
				return false
			}
		}

		for _, option := range f.options {
			if !option(f)(contact) {
				return false
			}
		}

		return true
	}
}

func WithName(name string) contactOption {
	return func(f *contactFilter) func(Contact) bool {
		f.name = append(f.name, name)
		return func(Contact) bool { return true }
	}
}

func WithPhone(phone string) contactOption {
	return func(f *contactFilter) func(Contact) bool {
		f.phone = append(f.phone, phone)
		return func(Contact) bool { return true }
	}
}

func WithEmail(email string) contactOption {
	return func(f *contactFilter) func(Contact) bool {
		f.email = append(f.email, email)
		return func(Contact) bool { return true }
	}
}

func (cm *contactManager) FindContacts(options ...contactOption) []Contact {
	filter := contactFilter{
		options: options,
	}

	var filteredContacts []Contact
	for _, contact := range cm.contacts {
		if filter.apply()(contact) {
			filteredContacts = append(filteredContacts, contact)
		}
	}

	return filteredContacts
}

func main() {
	cm := NewContactManager()
	cm.AddContact(Contact{Name: "Alice", Phone: "+1234567890", Email: "alice@example.com"})
	cm.AddContact(Contact{Name: "Bob", Phone: "+0987654321", Email: "bob@example.com"})
	cm.AddContact(Contact{Name: "Charlie", Phone: "+1122334455", Email: "charlie@example.com"})

	result := cm.FindContacts(WithName("Alice"), WithPhone("123"))
	fmt.Println("Found:", result)

	result = cm.FindContacts(WithEmail("example.com"))
	fmt.Println("Found:", result)
}
