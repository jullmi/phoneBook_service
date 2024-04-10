package main

import (
	"strconv"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// Initialized by the user – returns a pointer
// If it returns nil, there was an error
func initEntry(N, S, T string) *Entry {
	// Both of them should have a value
	if T == "" || S == "" {
		return nil
	}
	// Give LastAccess a value
	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}

var CSVPATH = "./data.csv"

type PhoneBook []Entry

var data = PhoneBook{}
var index map[string]int