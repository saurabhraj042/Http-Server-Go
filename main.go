package main

import (
    "log"
    "net/http"
)
func NewInMemoryPlayerStore() *InMemoryPlayerStore{
	return &InMemoryPlayerStore{
		map[string]int{},
	}
}
// InMemoryPlayerStore collects data about players in memory.
type InMemoryPlayerStore struct{
	store map[string]int
}

// GetPlayerScore retrieves scores for a given player.
func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
}

func main() {
	store := NewInMemoryPlayerStore()
	server := &PlayerServer{store}
	log.Fatal(http.ListenAndServe(":5050", server))
}