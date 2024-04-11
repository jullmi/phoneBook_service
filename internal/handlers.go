package internal

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const PORT = ":1234"

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get telephone
	paramStr := strings.Split(r.URL.Path, "/")

	fmt.Println("Path:", paramStr)

	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not found: "+r.URL.Path)
		return
	}

	log.Println("Serving:", r.URL.Path, r.Host)

	tel := paramStr[2]
	err := deleteEntry(tel)

	if err != nil {
		fmt.Println(err)
		Body := err.Error() + "\n"
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s", Body)
		return
	}

	Body := tel + "deleted!\n"
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "$s", Body)

}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := list()
	fmt.Fprintf(w, "%s", Body)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving: ", r.URL.Path, r.Host)
	w.WriteHeader(http.StatusOK)
	Body := fmt.Sprintf("Total entries: %d\n", len(Data))
	fmt.Fprintf(w, "%s", Body)
}

func InsertHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 5 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Not enough arguments"+r.URL.Path)
		return
	}

	name := paramStr[2]
	surname := paramStr[3]
	tel := paramStr[4]

	t := strings.ReplaceAll(tel, "-", "")
	if !MatchTel(t) {
		fmt.Println("not a valid phone number:", t)
		return
	}

	temp := Entry{Name: name, Surname: surname, Tel: t}
	err := insert(&temp)

	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		Body := "Failed to add record\n"
		fmt.Fprintf(w, "%s", Body)
	} else {
		log.Println("Serving:", r.URL.Path, "from", r.Host)
		w.WriteHeader(http.StatusOK)
		Body := "New record added successfully\n"
		fmt.Fprintf(w, "%s", Body)
	}
	log.Println("Serving: ", r.URL.Path, "from", r.Host)

}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	paramStr := strings.Split(r.URL.Path, "/")
	fmt.Println("Path:", paramStr)

	if len(paramStr) < 3 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Not found: "+r.URL.Path)
		return
	}

	var Body string
	tel := paramStr[2]

	t := search(tel)

	if t == nil {
		w.WriteHeader(http.StatusNotFound)
		Body = "Could not be found:" + tel + "\n"
	} else {
		w.WriteHeader(http.StatusOK)
		Body = fmt.Sprintf("%s, %s, %s", t.Name, t.Surname, t.Tel)
	}

	fmt.Println("Serving:", r.URL.Path, "from", r.Host)
	fmt.Fprintf(w, "%s", Body)

}
