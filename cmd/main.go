package main

import (
	"fmt"
	"net/http"
	"phonebook/internal"
	"time"
)

func main() {
	err := internal.SaveCSVFile(internal.CSVPATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.ReadCSVFile(internal.CSVPATH)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.CreateIndex()
	if err != nil {
		fmt.Println("Cannot create index.")
		return
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         internal.PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		WriteTimeout: time.Second,
		ReadTimeout:  time.Second,
	}

	mux.Handle("/list", http.HandlerFunc(internal.ListHandler))
	mux.Handle("/insert/", http.HandlerFunc(internal.InsertHandler))
	mux.Handle("/insert", http.HandlerFunc(internal.InsertHandler))
	mux.Handle("/search/", http.HandlerFunc(internal.SearchHandler))
	mux.Handle("/search", http.HandlerFunc(internal.SearchHandler))
	mux.Handle("/delete/", http.HandlerFunc(internal.DeleteHandler))
	mux.Handle("/status", http.HandlerFunc(internal.StatusHandler))
	mux.Handle("/", http.HandlerFunc(internal.DefaultHandler))

	fmt.Println("Ready to serve at", internal.PORT)
	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
