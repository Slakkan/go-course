package main

import "log"


func main() {
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User {"John", "Smith", "john@smith.com", 30})
	users = append(users, User {"Mary", "Johnes", "mary@johnes.com", 20})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}


// func main() {
// 	animals := make(map[string]string)
// 	animals["dog"] = "Fido"
// 	animals["cat"] = "grumpy"

// 	for key, animal := range animals {
// 		log.Println(key, animal)
// 	}
// }

// func main() {
// 	animals := []string {"dog", "fish", "horse", "cow", "cat"}

// 	for _, animal := range animals {
// 		log.Println(animal)
// 	}
// }

// func main() {
// 	for i := 0; i <= 5; i++ {
// 		log.Println(i)
// 	}
// }