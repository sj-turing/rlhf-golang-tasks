package main

import (
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type User struct {
	Name string
	Age  int
}

type UserStore struct {
	data sync.Map
	file string
	mux  sync.Mutex
}

func NewUserStore(file string) *UserStore {
	us := &UserStore{file: file}
	us.loadFromFile()
	return us
}
func (us *UserStore) loadFromFile() {
	us.mux.Lock()
	defer us.mux.Unlock()
	data, err := ioutil.ReadFile(us.file)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Error reading file: %v", err)
		}
		return
	}
	var m map[string]User
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&m); err != nil {
		log.Printf("Error decoding data: %v", err)
		return
	}
	for k, v := range m {
		us.data.Store(k, v)
	}
}
func (us *UserStore) saveToFile() {
	us.mux.Lock()
	defer us.mux.Unlock()
	var m map[string]User
	us.data.Range(func(k, v interface{}) bool {
		m[k.(string)] = v.(User)
		return true
	})
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(m); err != nil {
		log.Printf("Error encoding data: %v", err)
		return
	}
	if err := ioutil.WriteFile(us.file, buf.Bytes(), 0644); err != nil {
		log.Printf("Error writing file: %v", err)
	}
}
func (us *UserStore) Create(email string, user User) {
	us.data.Store(email, user)
	us.saveToFile()
}
func (us *UserStore) Read(email string) (User, bool) {
	v, ok := us.data.Load(email)
	if !ok {
		return User{}, false
	}
	return v.(User), true
}
func (us *UserStore) Update(email string, user User) {
	us.data.Store(email, user)
	us.saveToFile()
}
func (us *UserStore) Delete(email string) {
	us.data.Delete(email)
	us.saveToFile()
}
func main() {
	us := NewUserStore("users.gob")
	go func() {
		for {
			us.Create("alice@example.com", User{Name: "Alice", Age: 30})
		}
	}()
	go func() {
		for {
			us.Update("bob@example.com", User{Name: "Bob", Age: 25})
		}
	}()
	go func() {
		for {
			_, ok := us.Read("alice@example.com")
			_, ok := us.Read("bob@example.com")
		}
	}()
	go func() {
		for {
			us.Delete("alice@example.com")
			us.Delete("bob@example.com")
		}
	}()
	select {}
}
