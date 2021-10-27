package main

import (
	"fmt"
	"io"
)

type Venue struct { //We will need to define a venue type
	Audience int // We will need to take the number of audience members within this program.
	Log      io.Writer
}

func (a *Venue) Entertain(audience int, acts ...Entertainer) error { // *Venue is used as the pointer. This takes the Venue type and interfaces implement an Entertain method on *Venue.
	if len(acts) == 0 {
		return fmt.Errorf("there are no entertainers to perform")

	}
	a.Audience = audience
	for _, act := range acts {
		if err := a.play(act); err != nil {
			return err
		}
	}
	return nil
}

func (a Venue) play(act Entertainer) error {
	name := act.Name()

	if b, ok := act.(Setuper); ok { // The venue should check each entertainer to see if it implements the Setuper interface.
		if err := b.Setup(a); err != nil {
			return fmt.Errorf("%s: %w", name, err)

		}
		fmt.Fprint(a.Log, "%s has completed setup.\n", name) // Use the format, ("%s has completed setup. \n")

	}
	if err := act.Perform(a); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	fmt.Fprint(a.Log, "%s has performed for %d people.\n", name, a.Audience) // Use the format, ("%s has performed for %d people. \n")

	if c, ok := act.(Teardowner); ok { // The venue should check each entertainer to see if it implements the Teardowner interface.
		if err := c.Teardown(a); err != nil {
			return fmt.Errorf("%s: %w", name, err)

		}
		fmt.Fprintf(a.Log, "%s has completed teardown.\n", name) // Use the format, ("%s has completed teardown. \n")
	}

	return nil

}
