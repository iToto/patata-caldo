package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ttacon/emoji"

	"github.com/gorilla/mux"
)

// Potato is a potato
type Potato struct {
	Text    string  `json:"text"`
	History []Entry `json:"history"`
}

// Entry is a historic entry
type Entry struct {
	Node string `json:"node"`
	Text string `json:"text"`
	Desc string `json:"desc"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT environment variable must be set")
	} else {
		fmt.Println("Running on port: " + port)
	}

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/process", Process)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Process is the main RPC endpoint for this service that will manipulate a passed in string
func Process(w http.ResponseWriter, r *http.Request) {
	// Parse string to wrap words with ":"
	decoder := json.NewDecoder(r.Body)
	var potato Potato
	err := decoder.Decode(&potato)

	if err != nil {
		errorString := fmt.Errorf("Error parsing JSON string: %s", err)
		w.Write([]byte(errorString.Error()))
	}

	emoji, err := buildEmojiString(potato)
	if err != nil {
		errorString := fmt.Errorf("Could not build Emoji string: %s", err)
		w.Write([]byte(errorString.Error()))
	}

	var node string

	node, err = os.Hostname()
	if err != nil {
		node = "iToto-Default-Host"
	}

	// Build Entry
	entry := Entry{
		Node: node,
		Text: emoji,
		Desc: "Convert words to Emojis",
	}
	potato.History = append(potato.History, entry)
	json, _ := json.Marshal(potato)

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func buildEmojiString(potato Potato) (string, error) {
	wordsArray := strings.Split(potato.Text, " ")

	for i, word := range wordsArray {
		// lower case
		word = strings.ToLower(word)
		// wrap with ':'
		emojifyCol := ":" + word + ":"
		// Check if we have an emoji for this word
		emojify := emoji.Emoji(emojifyCol)
		if strings.Compare(emojify, emojifyCol) != 0 {
			// Overwrite word with emoji version of word
			wordsArray[i] = emojify
		}
	}

	return strings.Join(wordsArray, "  "), nil
}
