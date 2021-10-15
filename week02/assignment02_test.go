package assignment02

import (
	"log"
	"testing"
)

func Test_Array(t *testing.T) { // Always use this line to start a test file
	t.Parallel()

	exp := [4]string{"Benjamin", "Samantha", "Becca", "Ollie"}
	var act [4]string

	for i, x := range exp {
		act[i] = x
	}

	for i, a := range act {
		x := exp[i]
		if a != x {
			log.Fatalf("expected %s, got %s", x, a)
		}
	}
}

func Test_Slice(t *testing.T) { // Always use this line for a test file
	t.Parallel()

	exp := []string{"Benjamin", "Samantha", "Becca", "Ollie"}
	act := make([]string, 0, len(exp))

	for _, x := range exp {
		act = append(act, x)
	}

	if len(act) != len(exp) {
		log.Fatalf("expected slice length to be %d, got %d", len(exp), len(act))
	}

	for i, a := range act {
		x := exp[i]
		if a != x {
			log.Fatalf("expected %s, got %s", x, a)
		}
	}
}

func Test_Map(t *testing.T) { // Always use this line for a test file
	t.Parallel()

	exp := map[string]string{
		"Benjamin":   "Moore",
		"Samantha":   "Moore",
		"Becca": "Richards",
		"Ollie":  "Moore",
	}

	act := map[string]string{}

	for k, v := range exp {
		act[k] = v
	}

	if len(act) != len(exp) {
		log.Fatalf("expected map length to be %d, got %d", len(exp), len(act))
	}

	for k, a := range act {
		x, ok := exp[k]
		if !ok {
			log.Fatalf("could not find %s", k) // These two lines are VERY useful to double check everything is correct in your code.
		}

		if a != x {
			log.Fatalf("expected %s, %s", x, a)
		}
	}

}
