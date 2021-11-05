package main //Always start a program in go with this statement.

import ( //This will also be needed for every go program, you will need to import some packages to run the code.
	"fmt"
	"testing"
)

func assertModel(t testing.TB, act.Model, exp Model) {
	t.Helper()
	for k, v := range exp {
		if act[k] != v {
			t.Fatalf("expected %v, got %v", v, act[k])
		}
	}
}

func Test_Score_Select(t *testing.T) { //In order to run a test in go we will need to have the following format: Test_Name(t *testing.T).
	t.Parallel() // We need to declare t.Parallel() for this assignment.
	mary := Model{
		"name": "mary",
		"id": 1,
	}

	john := Model {
		"name" : "john",
		"id" : 2,
	}
	
	const tn = "users"
	fullStore := &Store{}
	fullStore.Insert(tn, mary, john)

	unknown := Clauses{
		"id" : -1,
	}

	table := []struct {
		name  string
		query Clauses
		table string
		err	  error
		exp   Models 

	}{
		{name: "happy path", table: tn, exp: []Model{mary, john}},
		{name: "table doesn't exist", table: "unknown", err: ErrTableNotFound{table:"unknown"},
		{
			name: "row not found.",
			table: tn,
			query: unknown,
			err: &errNoRows{
				table:	tn,
				clauses: unknown,

			},
		},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			act, err := fullStore.Select(tt.table, tt.query)

			if tt.err != nil {
				b := errors.Is(err, tt.err)
				if !b {
					t.Fatalf("expected %v, got %v", tt.err, err)

				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			el := len(tt.exp)
			al := len(act)
			if el != al {
				t.Fatalf("expected %v, got %v", el, al)
			}
			if act.String() != tt.exp.String() {
				t.Fatalf("expected %v, got %v", tt.exp, act)
			}

			for i, m := range tt.exp {
				assertModel(t, act[i], m)
			}
		})
	}
}

func Test_Store_Length(t *testing.T) {
	t.Parallel()
	t.Skip()
	s := &Store {

	}
}

