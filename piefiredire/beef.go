package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// beef summary gets beef article and takes out beefss
func beefSummary(w http.ResponseWriter, r *http.Request) {
	// get beefs from beef's url
	resp, err := http.Get(beefUrl)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	// get beef map return and response in json
	beef := countBeef(string(body))
	if err := WriteJson(w, http.StatusOK, map[string]map[string]int{
		"beef": beef,
	}); err != nil {
		log.Println(err)
	}
}

func countBeef(body string) map[string]int {
	// get rid of . and ,
	beefText := strings.ReplaceAll(strings.ReplaceAll(body, ".", ""), ",", "")
	beefs := strings.Fields(beefText)
	// loop and map beefs
	beef := make(map[string]int)
	for _, b := range beefs {
		if _, ok := beef[b]; !ok {
			beef[b] = 1
			continue
		}

		beef[b]++
	}
	return beef
}
