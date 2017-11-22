package arithmetic

import "testing"
import "fmt"

type mockStorage struct {
	items []float64
}

func TestCalculate(t *testing.T) {
	cases := []struct {
		in      string
		out     float64
		err     bool
		storage storage
	}{
		{
			// simple
			"1 + 2",
			3,
			false,
			nil,
		},
		{
			// no spaces
			"1+2",
			3,
			false,
			nil,
		},
		{
			// parenthesis
			"1 + (2 * 3) - 5",
			2,
			false,
			nil,
		},
		{
			// error; division by zero
			"10 / (1 - 1)",
			0,
			true,
			nil,
		},
		{
			// error; missing number at the end
			"1 +",
			0,
			true,
			nil,
		},
		{
			// error; parenthesis not closed
			"(1 + 2",
			0,
			true,
			nil,
		},
		{
			// error; double parenthesis not closed
			"((1 + 2)",
			0,
			true,
			nil,
		},
		{
			// error; double operation
			"2 + * 3",
			0,
			true,
			nil,
		},
		{
			// starts with an operator (ok, assumes leading 0)
			"- 2 + 3",
			1,
			false,
			nil,
		},
		{
			// variables
			"$1 * $2",
			8,
			false,
			newMockStorage(2, 4),
		},
		{
			// error; variable doesn't exist
			"$1 * $2",
			0,
			true,
			newMockStorage(2),
		},
		{
			// error; unknown operator
			"$1 ^ $2",
			0,
			true,
			nil,
		},
	}

	for _, testCase := range cases {
		result, err := Calculate(testCase.in, testCase.storage)

		if err == nil && testCase.err {
			t.Errorf("Expected error; didn't get one")
		}

		if err != nil && !testCase.err {
			t.Errorf("Didn't expect an error; got \"%v\"", err)
		}

		if !testCase.err && result != testCase.out {
			t.Errorf("Expected result to equal %f, got %f", testCase.out, result)
		}
	}
}

func (m *mockStorage) Get(pos int) (float64, error) {
	if len(m.items) < pos {
		return 0, fmt.Errorf("Index %d out of reach in storage", pos)
	}

	return m.items[pos-1], nil
}

func (m *mockStorage) Append(_ float64) {
}

func newMockStorage(items ...float64) *mockStorage {
	return &mockStorage{items: items}
}
