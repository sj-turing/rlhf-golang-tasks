package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}
type KeyValueStore struct {
	mu   *sync.RWMutex
	data map[int64]User
	file string
}

func NewKeyValueStore(file string) *KeyValueStore {
	kv := &KeyValueStore{
		data: make(map[int64]User),
		file: file,
		mu:   &sync.RWMutex{},
	}
	kv.load()
	return kv
}
func (kv *KeyValueStore) load() error {
	data, err := os.ReadFile(kv.file)
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
	kv.mu.Lock()
	defer kv.mu.Unlock()
	data, err := json.MarshalIndent(kv.data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(kv.file, data, 0644)
}
func (kv *KeyValueStore) Put(id int64, user User) {
	kv.mu.Lock()
	kv.data[id] = user
	kv.mu.Unlock()
}
func (kv *KeyValueStore) Get(id int64) (User, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	user, ok := kv.data[id]
	return user, ok
}
func (kv *KeyValueStore) Delete(id int64) {
	kv.mu.Lock()
	delete(kv.data, id)
	kv.mu.Unlock()
}

const alphabets = "abcdefghijklmnopqrstuvwxyz"

func randomString(size int) string {
	if size == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()

	for size > 0 {
		buf.WriteByte(alphabets[rand.Intn(len(alphabets))])
		size--
	}
	return buf.String()
}

var emailHosts = [3]string{"example.com", "learning.tech", "test.io"}

func randomEmail(size int) string {
	if size == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	defer buf.Reset()

	for size > 0 {
		buf.WriteByte(alphabets[rand.Intn(len(alphabets))])
		size--
	}

	buf.WriteByte('@')
	buf.WriteString(emailHosts[rand.Intn(len(emailHosts))])

	return buf.String()
}

func generateRandomUsers() <-chan User {
	userInfoChan := make(chan User)
	// generating random records for 2 seconds
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		defer close(userInfoChan)

		for {
			select {
			case <-ticker.C:
				return
			default:
				userInfoChan <- User{ID: time.Now().UnixMilli(), Email: randomEmail(rand.Intn(30)), Name: randomString(rand.Intn(30)), Age: rand.Intn(80)}
			}
		}
	}()

	return userInfoChan
}

func main() {
	kvStore := NewKeyValueStore("users.json")
	defer kvStore.save()

	userInfoChan := generateRandomUsers()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for userInfo := range userInfoChan {
			kvStore.Put(userInfo.ID, userInfo)
		}
	}()

	size := 100
	for size > 0 {
		key := time.Now().Add(-100 * time.Millisecond).UnixMilli()
		kvStore.Get(key)
		size--
	}

	wg.Wait()

	fmt.Println("records created: ", len(kvStore.data))
}
