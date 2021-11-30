package week08

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Manager_Demonstration(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	m := &Manager{}

	ctx, err := m.Start(ctx, 5)
	r.NoError(err)

	for i := 0; i < 10; i++ {
		go m.Assign(ProductA)
		go m.Assign(ProductB)
	}

	var completed []CompletedProduct

	go func() {
		fmt.Println("waiting for a completed product")

		for cp := range m.Completed() {
			completed = append(completed, cp)

			if len(completed) >= 20 {
				m.Stop()
			}
		}
	}()

	fmt.Println("waiting for the ctx to be cancelled")
	<-ctx.Done()

	fmt.Println("validating output")
	r.Len(completed, 20)
	fmt.Println("validated")
}
