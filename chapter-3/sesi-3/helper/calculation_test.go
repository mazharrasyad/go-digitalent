package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFailSum(t *testing.T) {
	result := Sum(10, 10)

	// Contoh 3
	// if result != 40 {
	// t.Fail() // Contoh 1
	// t.FailNow() // Contoh 2
	// 	t.Error("Result has to be 40")
	// }

	require.Equal(t, 40, result, "Result has to be 40")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestSum(t *testing.T) {
	result := Sum(10, 10)

	// Contoh 3
	// if result != 20 {
	// panic("Result should be 20") // Contoh 1
	// 	t.FailNow() // Contoh 2
	// }

	assert.Equal(t, 20, result, "Result has to be 20")

	fmt.Println("TestFailSum Eksekusi Terhenti")
}

func TestTableSum(t *testing.T) {
	tests := []struct {
		request  int
		expected int
		errMsg   string
	}{
		{
			request:  Sum(10, 10),
			expected: 20,
			errMsg:   "Result has to be 20",
		},
		{
			request:  Sum(20, 20),
			expected: 40,
			errMsg:   "Result has to be 40",
		},
		{
			request:  Sum(25, 25),
			expected: 50,
			errMsg:   "Result has to be 50",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test-%d", i), func(t *testing.T) {
			require.Equal(t, test.expected, test.request, test.errMsg)
		})
	}
}
