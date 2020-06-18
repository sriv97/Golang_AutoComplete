package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strings"
)

func process(filepath string) map[string]int {

	// Map to hold word --> frequency
	m := make(map[string]int)

	// Read file
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// Text processing:
	// Make everything lowercase
	st := strings.ToLower(string(data))

	// Remove all punctuation
	reg, err := regexp.Compile("[^a-zA-Z\\s\\-]+")
	if err != nil {
		log.Fatal(err)
	}

	st = reg.ReplaceAllString(st, "")

	//Split words by whitespace
	words := strings.Fields(st)

	for _, word := range words {
		_, ok := m[word]
		if ok {
			m[word] += 1
		} else {
			m[word] = 1
		}
	}

	// fmt.Println(m) ()
	// fmt.Println(len(m))

	return m
}

func autocomplete(source map[string]int, term string, top int) []string {
	ret := make([]string, 0)

	possible := make(map[string]int)
	for k, v := range source {
		if strings.Contains(k, term) {
			possible[k] = v
		}
	}

	// Sort possible
	tmp := map[int][]string{}
	var a []int
	for k, v := range possible {
		tmp[v] = append(tmp[v], k)
	}
	for k := range tmp {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range tmp[k] {
			ret = append(ret, s)
		}
	}

	if len(ret) < top {
		return ret[:len(ret)]
	} else {
		return ret[:top]
	}
}

var source map[string]int

func apiResponse(w http.ResponseWriter, r *http.Request) {
	//return type is set to JSON
	w.Header().Set("Content-Type", "application/json")

	//Response methods
	switch r.Method {
	case "GET":
		keys, ok := r.URL.Query()["term"]

		if !ok {
			log.Println("Url parameter `term` is missing")
			return
		}

		term := keys[0]
		suggests := autocomplete(source, term, 25)

		log.Println(len(suggests))
		j, _ := json.Marshal(suggests)
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Can't find method requested"}`))
	}
}

func main() {
	source = process("shakespeare-complete.txt")
	http.HandleFunc("/autocomplete", apiResponse)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))

}
