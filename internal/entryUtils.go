package internal

import (
	"fmt"
	"strconv"
	"time"
)

func insert(pS *Entry) error {
	// If it already exists, do not add it
	_, ok := Index[(*pS).Tel]
	if ok {
		return fmt.Errorf("%s already exists", pS.Tel)
	}

	*&pS.LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	Data = append(Data, *pS)
	// Update the index
	_ = CreateIndex()

	err := SaveCSVFile(CSVPATH)
	if err != nil {
		return err
	}
	return nil
}

func deleteEntry(key string) error {
	i, ok := Index[key]
	if !ok {
		return fmt.Errorf("%s cannot be found!", key)
	}
	Data = append(Data[:i], Data[i+1:]...)
	// Update the index - key does not exist any more
	delete(Index, key)

	err := SaveCSVFile(CSVPATH)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) *Entry {
	i, ok := Index[key]
	if !ok {
		return nil
	}
	Data[i].LastAccess = strconv.FormatInt(time.Now().Unix(), 10)
	return &Data[i]
}



func list() string {
	var all string
	for _, k := range Data {
		all = all + k.Name + " " + k.Surname + " " + k.Tel + "\n"
	}
	return all
}