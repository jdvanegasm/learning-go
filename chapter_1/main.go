package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func generatePass(length int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
	rand.Seed(time.Now().UnixNano()) // Init the seed for the random numbers

	password := make([]byte, length) // dynamic slice for the password concatenation

	for i := range password {
		password[i] = alphabet[rand.Intn(len(alphabet))] // characters concatenation
	}
	return string(password) // returning password in string type
}

func savePassword(password string) {
	file, err := os.OpenFile("passwords.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("An unexpected error has occurred while opening the password file: ", err)
		return
	}
	defer file.Close()

	// Writes the password in a new line
	_, err = file.WriteString(password + "\n")
	if err != nil {
		fmt.Println("An unexpected error has occured while writing to the password file: ", err)
		return
	}

	fmt.Println("The password has been saved sucessfully in 'passwords.txt'")
}

func main() {

	var length int
	var passNum int

	fmt.Println("--------------------------------------")
	fmt.Println("welcome to your password generator!")
	fmt.Println("--------------------------------------")

	fmt.Println("Type how many passwords you want to generate!")
	fmt.Scan(&passNum)

	if passNum <= 0 {
		fmt.Println("\nPlease type a valid number (greater than 0!)")
		return
	}

	for i := range passNum {
		fmt.Println("--------------------------------------")
		fmt.Printf("Password number: %d", i+1)
		fmt.Println("\nType the lenght for your password!")
		fmt.Scan(&length)
		if length <= 0 {
			fmt.Println("\nPlease, type a valid length (greater than 0)!")
			return
		}
		pass := generatePass(length)
		fmt.Printf("Your password is: %s \n", pass)
		savePassword(pass)
	}
}
