package week07

import (
	"context" // We will need this package in order to cancel and interupt in the program.
	"fmt"
	"testing" // We will need this package in order to run the tests.
	"time"    // We will need this in order to add in the 5 second delay to the program.
)

// We will need to implement test cases for the Run function.
func Test_Run(t *testing.T) {
	t.Parallel()

	// Tests you will need to write:

	// 1. Timeout after 5 seconds if nothing happens. This is why we will need the time package.
	// 2. Interruption by a signal.
	// 3. Run returns when the products are completed.
	// 4. Test that the output is correct.

}

func Test_Manager(t *testing.T) { // This is the syntax needed to run test cases.
	t.Parallel()

	r := require.New(t)

	// This is where requirement #1 comes into play: timeout after 5 seconds if nothing happens.

	ctx := context.Background() // This is the syntax needed to start the context.
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	m := &Manager{}

	ctx, err := m.Start(ctx, 5)
	r.NoError(err)

	contextual.Print(ctx)

	for i := 0; i < 10; i++ {
		go m.Assign(ProductA) // This is the syntax we need to use to start a goroutine.
		go m.Assign(ProductB)
	}

	var completed []CompletedProduct

	go func() {
		fmt.Println("Waiting for a completed product")

		for cp := range m.Completed() {
			completed = append(completed, cp)

			if len(completed) >= 20 {
				m.Stop()
			}
		}
	}()

	// This will test that the output is correct.

	fmt.Println("validating output")
	r.len(completed, 20)
	fmt.Println("validated")
}
