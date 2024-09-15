package main //Always start a program in go with this statement.

import ( //This will also be needed for every go program, you will need to import some packages to run the code.
	"testing"
)

func Test_Clauses(t *testing.T) {
	t.Parallel()
	t.Skip()
	one := Clauses{"name": "mary"}
	two	:= Clauses{"name": "mary", "id": 1}
	three := Clauses{"name": "mary", "id": 1, "age": 42}

	table := []struct {
		name	string
		in		Clauses
		exp		string
	}{
		{name: "empty"},
		{name: "one clause", in: one, exp: "name" = "mary"},
		{name: "two clauses", in: two, exp: "id" = "1" and "name" = "mary"},
		{name: "three clauses", in: three, exp: "age" = "42" and "id" = "1" and "name" = "mary"},
	}

	for _, tt := range table {
		t.Run (tt.name, func(t *testing.T) {

			act := tt.in.String()
			if act != tt.exp {
				t.Fatalf("expected %q, got %q", tt.exp, act)
			}
		})
	}
}

func Test_Clauses_Match(t *testing.T) {
	t.Parellel()
	t.Skip() 

	mod := Model{
		"first": "mary",
		"age" : 42,
	}

	table := []struct {
		query 	Clauses
		name	string
		err		bool
	}{
		{name: "empty"},
		{name: "no clauses"},
		{name: "all the mary's", query: Clauses{"first": "mary"}},
		{name: "multiple clauses", query: Clauses{
			"first": "mary",
			"age":	42,
		}},
	}

	for _, tt := range table {
		t.Run9tt.name, func(t *testing.T) {

			act := !tt.query.Match(mod)
			if act != tt.err {
				t.Fatalf("expected %t, got %t", tt.err, act)
			}
		})
	}
}

