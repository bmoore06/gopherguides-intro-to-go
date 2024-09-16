package week06 // You will need this line every time you begin your code

// You will need this to import different packages into your code.

import (
	"testing"
)

func Test_Employee_work(t *testing.T) { //This is the layout for every test you will run.
	t.Parallel() // Research what this really is.

	good := Employee(45) //There will need to be both a good and bad employee.
	bad := Employee(-1)  // A bad employee will not perform the work.

	table := []struct { //Research this as well!
		name string
		e    Employee
		err  bool
	}{
		{name: "happy", e: good},
		{name: "Sad", e, bad},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager()
			defer m.Stop()
			e := tt.e
			go e.work(m) // You will need this or else your code will block.

			p := &Product{Quantity: 2}
			err := m.Assign(p)
			NotError(t, err)

			var cp CompletedProduct

			//This will be the sad path, every sad path will return an error.

			select { //This will clean up all the resources. This is needed because we have a happy and a sad path. The select statement will allow us to listen to both the bad and good.
			case err := <-m.Errors():
				if tt.err {
					Error(t, err)
					return
				}

			case cp := <-m.Completed(): // This will block the loop until this line is read.
				NoError(t, cp.IsValid())
			}
			NoError(t, cp.IsValid())

		}) //Research why this parenthesis is here!
	}
}

func Test_Employee_Cancellation(t *testing.T) { //We will need this in case the employee cancels the job that came through.
	t.Parallel()

	good := Employee(45)
	bad := Employee(-1)

	table := []struct { //This will be where you define the variables and the variable type.
		name string
		fn   func(m *Manager)
	}{
		{name: "cancel jobs", fn: func(m *Manager) {
			close(m.Jobs())
		}},
		{name: "done", fn: func(m *Manager) {
			m.Stop()
		}},
	}

	for _, tt := range table {
		t.Run(tt.name, func(t *testing.T) {
			m := NewManager()
			defer m.Stop()

			e := Employee(45)

			go func() {
				tt.fn(m)
			}()
			e.work(m)
		})
	}

}

func NoError(t *testing.TB, err error) { //You will need to test that there is both no errors or that there are errors.
	t.Helper() //Research this!

	if err != nil {
		t.Fatalf("expected no error, go %v", err)
	}
}

func Error(t testing.TB, err error) {
	t.Helper() //Research this!

	if err == nil {
		t.Fatalf("expected error, got nil")

	}
}
