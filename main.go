// https://blog.gopheracademy.com/advent-2016/enigma-emulator-in-go/
// Was confused on the technical aspect of creating this machine, however Edward Medvedev gave a wonderful breakdown and I feel as if this article helped my understanding of the actual enigma machine as well as GoLang which I've been specifically struggling with
package main

import (
	"fmt"
	"github.com/emedvedev/enigma"
)

func main() {
	// Encoding function is defined
	encode := enigma.NewEnigma(
		// Rotors are defined with established starting points and go through rotation of alphabet
		[]enigma.RotorConfig{
			enigma.RotorConfig{ID: "III", Start: 'I', Ring: 1},
			enigma.RotorConfig{ID: "II", Start: 'S', Ring: 1},
			enigma.RotorConfig{ID: "IV", Start: 'T', Ring: 1},
		},
		// Plugboard establishes the specific change in letters
		"B",
		[]string{"AB", "CD", "EF"},
	)
	// Encoding function is used
	plaintext := "LMSIX"

	ciphertext := encode.EncodeString(plaintext)

	fmt.Printf("%s is encoded as %s\n", plaintext, ciphertext)

	decode := enigma.NewEnigma(
		// Rotors are defined same as encode and go through their rotation
		[]enigma.RotorConfig{
			enigma.RotorConfig{ID: "III", Start: 'I', Ring: 1},
			enigma.RotorConfig{ID: "II", Start: 'S', Ring: 1},
			enigma.RotorConfig{ID: "IV", Start: 'T', Ring: 1},
		},
		// Plugboard changes letters back
		"B",
		[]string{"AB", "CD", "EF"},
	)
	decoded := decode.EncodeString(ciphertext)

	fmt.Printf("%s is encoded as %s\n", ciphertext, decoded)
}
