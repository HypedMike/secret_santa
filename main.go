package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const MODULE = 10

func main() {
	if os.Args[1] == "run" {
		startSanta()
	} else if os.Args[1] == "decrypt" {
		decryptMyName(os.Args[2])
	}
}

func decryptMyName(name string) {

	fmt.Printf("Decrypting %s\n", name)
	fmt.Printf("Result: %s\n", caesarDecrypt(name, 10))
}

func startSanta() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// read file
	rawContent, err := os.ReadFile(os.Getenv("FILE_PATH"))
	if err != nil {
		panic(err)
	}

	// parse json
	jsonStruct := JsonStruct{}
	err = json.Unmarshal(rawContent, &jsonStruct)
	if err != nil {
		panic(err)
	}

	// create people
	var people []Person
	for _, name := range jsonStruct.People {
		people = append(people, Person{Name: name})
	}

	// set relationships
	for _, relationship := range jsonStruct.Relationships {
		p1 := findPerson(people, relationship.P1)
		p2 := findPerson(people, relationship.P2)

		// find p1 and p2
		for i := 0; i < len(people); i++ {
			if people[i].Name == p1.Name {
				people[i].SetPartner(&p2)
			}
			if people[i].Name == p2.Name {
				people[i].SetPartner(&p1)
			}
		}
	}

	// print results
	for _, person := range people {
		if person.Partner == nil {
			fmt.Printf("%20s has no partner\n", person.Name)
			continue
		}
		fmt.Printf("%20s -> %20s\n", person.Name, person.Partner.Name)
	}

	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")
	fmt.Println("------------------------------------------------------------")

	// randomize matches
	randomizedPeople := randomizeMatches(people)

	// print results
	for _, person := range randomizedPeople {
		fmt.Printf("%20s -> %20s\n", person.Name, person.Match)
	}

}

func randomizeMatches(people []Person) []CryptoMatch {
	matches := make([]Person, len(people))

	for i := 0; i < len(matches); i++ {
		randomNumber := rand.Intn(100)
		randomPerson := people[randomNumber%len(people)]

		matches[i] = randomPerson

		// check that this person is not a partner of the previous one
		if i > 0 && matches[i-1].Partner != nil && matches[i-1].Partner.Name == randomPerson.Name {
			i--
			continue
		}

		// check that this person is not already matched
		for j := 0; j < i; j++ {
			if matches[j].Name == randomPerson.Name {
				i--
				continue
			}
		}
	}

	var matchesWithNames []CryptoMatch

	// print results
	for i := 1; i < len(matches); i++ {
		matchesWithNames = append(matchesWithNames, CryptoMatch{
			Name:  matches[i-1].Name,
			Match: caesarEncrypt(matches[i].Name, MODULE),
		})
	}

	matchesWithNames = append(matchesWithNames, CryptoMatch{
		Name:  matches[len(matches)-1].Name,
		Match: caesarEncrypt(matches[0].Name, MODULE),
	})

	return randomizeArray(matchesWithNames)

	// fmt.Printf("%20s -> %20s\n", matches[len(matches)-1].Name, caesarEncrypt(matches[0].Name, 10))
}

func randomizeArray(input []CryptoMatch) []CryptoMatch {
	rand.Shuffle(len(input), func(i, j int) {
		input[i], input[j] = input[j], input[i]
	})
	return input
}

func findPerson(people []Person, name string) Person {
	for _, person := range people {
		if person.Name == name {
			return person
		}
	}
	panic("person not found")
}

func caesarEncrypt(name string, module int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	// normalize name
	if len(name) < 15 {
		i := 15 - len(name)
		for i >= 0 {
			name += string(alphabet[rand.Intn(len(alphabet))])
			i--
		}
	}

	result := ""

	for _, char := range name {
		charIndex := strings.Index(alphabet, string(char))
		charIndex = (charIndex + module) % len(alphabet)
		result += string(alphabet[charIndex])
	}

	return result
}

func caesarDecrypt(name string, module int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	result := ""

	for _, char := range name {
		charIndex := strings.Index(alphabet, string(char))
		charIndex = (charIndex - module) % len(alphabet)
		if charIndex < 0 {
			charIndex = len(alphabet) + charIndex
		}
		result += string(alphabet[charIndex])
	}

	return result
}
