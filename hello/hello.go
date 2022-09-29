package main

import "fmt"

func main() {
	greetings := hello("Victor", "en")
	fmt.Println(greetings)
}
// hello , je vien de mars 
func hello(name string, language string) string {
	switch language {
	case "en":
		return "Hello " + name
	case "es":
		return "Hola " + name
	case "fr":
		return "Bonjour " + name
	}
	return "Wrong Language !"
}
