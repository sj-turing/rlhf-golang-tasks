package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
type KeyValueStore struct {
	data map[int]User
	file string
}

func NewKeyValueStore(file string) *KeyValueStore {
	kv := &KeyValueStore{
		data: make(map[int]User),
		file: file,
	}
	kv.load()
	return kv
}
func (kv *KeyValueStore) load() error {
	data, err := ioutil.ReadFile(kv.file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	if err := json.Unmarshal(data, &kv.data); err != nil {
		return err
	}
	return nil
}
func (kv *KeyValueStore) save() error {
	data, err := json.MarshalIndent(kv.data, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(kv.file, data, 0644)
}
func (kv *KeyValueStore) Put(id int, user User) {
	kv.data[id] = user
}
func (kv *KeyValueStore) Get(id int) (User, bool) {
	user, ok := kv.data[id]
	return user, ok
}
func (kv *KeyValueStore) Delete(id int) {
	delete(kv.data, id)
}
func main() {
	kvStore := NewKeyValueStore("users.json")
	defer kvStore.save()
	user1 := User{Email: "user1@example.com", Name: "User One", Age: 25}
	kvStore.Put(1, user1)
	user2, ok := kvStore.Get(1)
	if ok {
		fmt.Println("User 1:", user2)
	} else {
		fmt.Println("User 1 not found.")
	}
	kvStore.Delete(1)
	user2, ok = kvStore.Get(1)
	if ok {
		fmt.Println("User 1:", user2)
	} else {
		fmt.Println("User 1 not found.")
	}
}
