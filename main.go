package main

import (
	"fmt"
	"strings"
)

// Function to create the Playfair matrix from a keyword
func createPlayfairMatrix(key string) [5][5]string {
	matrix := [5][5]string{}
	letters := "ABCDEFGHIKLMNOPQRSTUVWXYZ"
	key = strings.ToUpper(key + letters)

	// Initialize matrix with empty strings
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			matrix[i][j] = ""
		}
	}

	// Fill the matrix with unique characters from the key
	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for matrix[i][j] != "" {
				k++
			}
			matrix[i][j] = string(key[k])
		}
	}

	return matrix
}

// Function to find the coordinates of a letter in the Playfair matrix
func findLetterCoordinates(matrix [5][5]string, letter string) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if matrix[i][j] == letter {
				return i, j
			}
		}
	}
	return -1, -1
}

// Function to encrypt plaintext using the Playfair Cipher
func encryptPlayfair(plaintext, key string) string {
	plaintext = strings.ToUpper(strings.ReplaceAll(plaintext, "J", "I"))
	matrix := createPlayfairMatrix(key)
	ciphertext := ""
	i := 0

	for i < len(plaintext) {
		letter1 := string(plaintext[i])
		letter2 := ""
		if i+1 < len(plaintext) {
			letter2 = string(plaintext[i+1])
		}

		if letter1 == letter2 {
			letter2 = "X"
			i++
		} else {
			i += 2
		}

		row1, col1 := findLetterCoordinates(matrix, letter1)
		row2, col2 := findLetterCoordinates(matrix, letter2)

		if row1 == row2 {
			col1 = (col1 + 1) % 5
			col2 = (col2 + 1) % 5
		} else if col1 == col2 {
			row1 = (row1 + 1) % 5
			row2 = (row2 + 1) % 5
		} else {
			col1, col2 = col2, col1
		}

		ciphertext += matrix[row1][col1] + matrix[row2][col2]
	}

	return ciphertext
}

// Function to decrypt Playfair ciphertext
func decryptPlayfair(ciphertext, key string) string {
	matrix := createPlayfairMatrix(key)
	plaintext := ""

	for i := 0; i < len(ciphertext); i += 2 {
		letter1 := string(ciphertext[i])
		letter2 := string(ciphertext[i+1])

		row1, col1 := findLetterCoordinates(matrix, letter1)
		row2, col2 := findLetterCoordinates(matrix, letter2)

		if row1 == row2 {
			col1 = (col1 - 1 + 5) % 5
			col2 = (col2 - 1 + 5) % 5
		} else if col1 == col2 {
			row1 = (row1 - 1 + 5) % 5
			row2 = (row2 - 1 + 5) % 5
		} else {
			col1, col2 = col2, col1
		}

		plaintext += matrix[row1][col1] + matrix[row2][col2]
	}

	return plaintext
}

func main() {
	key := "KEYWORD"
	plaintext := "HELLO"
	ciphertext := encryptPlayfair(plaintext, key)
	decrypted := decryptPlayfair(ciphertext, key)

	fmt.Println("Plaintext:", plaintext)
	fmt.Println("Ciphertext:", ciphertext)
	fmt.Println("Decrypted:", decrypted)
}
