package main

type personName struct {
	firstName  string
	lastName   string
	pushNumber string
}

type personNames []personName

func (p personNames) Len() int {
	return len(p)
}

func (p personNames) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p personNames) Less(i, j int) bool {
	return p[i].pushNumber < p[j].pushNumber
}
