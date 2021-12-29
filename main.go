package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var Lyrics = []string{
	"the horse and the hound and the horn that belonged to",
	"the farmer sowing his corn that kept",
	"the rooster that crowed in the morn that woke",
	"the priest all shaven and shorn that married",
	"the man all tattered and torn that kissed",
	"the maiden all forlorn that milked",
	"the cow with the crumpled horn that tossed",
	"the dog that worried",
	"the cat that killed",
	"the rat that ate",
	"the malt that lay in",
	"the house that Jack built",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// Phase 1
	fmt.Println("Phase 1 Answer")
	fmt.Printf("Recite(1) : %s \n", Recite(1))
	fmt.Printf("Recite(2) : %s \n", Recite(2))
	fmt.Printf("Recite(12) : %s \n\n", Recite(12))

	// Phase 2
	fmt.Println("Phase 2 Answer")
	fmt.Println("A answer ")
	fmt.Printf("ReciteRand() : %s\n\n", ReciteRand())

	fmt.Println("B answer")
	fmt.Printf("ReciteSubject(1) : %s \n", ReciteSubject(1))
	fmt.Printf("ReciteSubject(2) : %s \n", ReciteSubject(2))
	fmt.Printf("ReciteSubject(3) : %s \n\n", ReciteSubject(3))

	fmt.Println("C answer")
	fmt.Printf("ReciteRand() : %s\n", ReciteSubjectRand())
}

func Recite(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		result = Lyrics[len(Lyrics)-i] + " " + result
	}

	result = "This is " + result

	return result
}

func ReciteRand() string {
	num := rand.Intn(len(Lyrics))
	result := ""
	for i := 1; i <= num; i++ {
		result = Lyrics[len(Lyrics)-i] + " " + result
	}

	result = "This is " + result

	return result
}

func ReciteSubject(num int) string {
	result := ""
	for i := 1; i <= num; i++ {
		tmpLyric := strings.Split(Lyrics[len(Lyrics)-i], " that")

		if i == 2 {
			result = tmpLyric[0] + " and " + result
		} else if i > 2 {
			result = tmpLyric[0] + ", " + result
		} else {
			result = tmpLyric[0] + " " + result
		}

	}

	result = "This is " + result

	return result
}

func ReciteSubjectRand() string {
	num := rand.Intn(len(Lyrics))
	result := ""
	for i := 1; i <= num; i++ {
		tmpLyric := strings.Split(Lyrics[len(Lyrics)-i], " that")

		if i == 2 {
			result = tmpLyric[0] + " and " + result
		} else if i > 2 {
			result = tmpLyric[0] + ", " + result
		} else {
			result = tmpLyric[0] + " " + result
		}

	}

	result = "This is " + result

	return result
}
