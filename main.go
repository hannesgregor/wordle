package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}

func GetRandomWord() {
	defer duration(track("getRandomWord"))
	resp, err := http.Get("https://api.frontendexpert.io/api/fe/wordle-words")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	bs := string(body)

	words := strings.Split(bs, ",")
	fmt.Println("Amount of random words:", len(words))

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(words))
	pick := words[randomIndex]

	fmt.Println(strings.Trim(pick, `""`))
	fmt.Println("The random word is: \n", pick)
}

func main() {
	start := time.Now()

	// Code to measure
	GetRandomWord()
	duration := time.Since(start)

	// Formatted string, such as "2h3m0.5s" or "4.503μs"
	fmt.Println(duration)

	// Nanoseconds as int64
	fmt.Println(duration.Nanoseconds())

}
