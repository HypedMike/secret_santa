package main

type Person struct {
	Name    string
	Partner *Person
}

type CryptoMatch struct {
	Name  string
	Match string
}

func (p *Person) SetPartner(partner *Person) {
	p.Partner = partner
}

type JsonStruct struct {
	People        []string `json:"people"`
	Relationships []struct {
		P1 string `json:"p1"`
		P2 string `json:"p2"`
	} `json:"relationships"`
}
