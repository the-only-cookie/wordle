package main

import (
 "fmt"
 "math/rand"
 "time"
)

func main() {
 rand.Seed(time.Now().UnixNano())

 var words []string
 var wordToGuess string
 var guess string
 maxAttempts := 5
 attemptsLeft := maxAttempts

 // Запрос пользовательского ввода для списка слов
 fmt.Println("Welcome to WORDLY!")
 fmt.Println("Enter a list of words, separated by spaces:")
 fmt.Print("> ")
 input := ""
 fmt.Scanln(&input)
 words = splitInput(input)

 // Выбор случайного слова из пользовательского списка
 if len(words) > 0 {
  wordToGuess = words[rand.Intn(len(words))]
  fmt.Println("Try to guess the word.")
 } else {
  fmt.Println("No words provided. Exiting the game.")
  return
 }

 for attemptsLeft > 0 {
  fmt.Printf("Attempts left: %d\n", attemptsLeft)
  fmt.Print("Enter your guess: ")
  fmt.Scanln(&guess)

  if guess == wordToGuess {
   fmt.Println("Congratulations! You guessed the word correctly.")
   return
  } else {
   attemptsLeft--
   fmt.Println("Incorrect guess. Try again.")
  }
 }

 fmt.Printf("Sorry, you've run out of attempts. The word was: %s\n", wordToGuess)
}

// Функция для разделения введенной строки на список слов
func splitInput(input string) []string {
 words := []string{}
 word := ""
 for _, char := range input {
  if char == ' ' {
   if word != "" {
    words = append(words, word)
    word = ""
   }
  } else {
   word += string(char)
  }
 }
 if word != "" {
  words = append(words, word)
 }
 return words
}
