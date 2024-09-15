package main //Always start a program in go with this statement.

import ( //This will also be needed for every go program, you will need to import some packages to run the code.
	"errors"
	"fmt"
	"io"
	"testing"
)

func Test_IsTableNotFound(t *testing.T) {
	t.Parallel() // Make sure to include t.Parallel() in the test files from the instructions.
	t.Skip()
	etnf := ErrTableNotFound{table: "users"}

	table := []struct {
		name string
		err  error
		exp  bool
	}{
		{name: "good error", err: etnf, exp: true},
		{name: "wrapped error", err: fmt.Errorf("oops %w", etnf), exp: true},
		{name: "nil"},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			if tt.err != nil && !IsErrTableNotFound(tt.err) {
				t.Fatalf("expected %v to be a ErrTableNotFound", tt.err)

			}
			if !tt.exp {
				return
			}
			e := ErrTableNotFound{}

			as := errors.As(tt.err, &e)
			if !as {
				t.Fatalf("expected %v to be an ErrTableNotFound", tt.err)
			}

			exp := "users"
			act := e.TableNotFound()

			if act != exp {
				t.Fatalf("expected %v to be %v", act, exp)

			}
		})
	}
}

func Test_IsErrNoRows(t *testing.T) {
	t.Parallel()
	t.Skip()
	etnf := &errNoRows{}

	table := []struct {
		name string
		err  error
		exp  bool
	}{
		{name: "good error", err: etnf, exp: true},
		{name: "wrapped error", err: fmt.Errorf("oops %w", etnf), exp: true},
		{name: "nil"},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {

			act := IsErrNoRows(tt.err)
			exp := tt.exp

			if act != exp {
				t.Fatalf("expected %v to be %v", act, exp)
			}
		})
	}
}

func Test_AsErrNoRows(t *testing.T) {
	t.Parallel()
	t.Skip()
	var err error = io.EOF

	_, ok := AsErrNoRows(err)

	if ok {
		t.Fatal("expected %v to not be an ErrNoRows", err)
	}

	err = &errNoRows{
		table: "users",
		clauses: Clauses{
			"age":  42,
			"name": "mary",
		},
	}

	err = fmt.Errorf("wrapping err %w", err)

	enr, ok := AsErrNoRows(err)
	if !ok {
		t.Fatalf("expected %v to be an ErrNoRows", err)
	}

	table, clauses := enr.RowNotconst

	exp := "users"
	act := table
	if act != exp {
		t.Fatalf("expected %v to be %v", act, exp)
	}

	al := len(clauses)
	el := 2
	if al != el {
		t.Fatalf("expected %v to be %v", al, el)
	}
}
