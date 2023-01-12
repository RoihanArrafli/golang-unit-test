package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// benchmark

func BenchmarkSub(b *testing.B)  {
	b.Run("Roihan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Roihan")
		}
	})
	b.Run("Budi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Budi")
		}
	})
}

func BenchmarkRoihan(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Roihan")
	}
}

func BenchmarkHelloWord(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		HelloWorld("Hello Word")
	}
}

func BenchmarkTable(b *testing.B)  {
	benchmarks := []struct {
		name string
		request string
	}{
		{
			name: "Roihan",
			request: "Roihan",
		},
		{
			name: "Eko",
			request: "Eko",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// Test table
func TestTableHelloWorld(t *testing.T)  {
	test := []struct {
		name string
		request string
		expected string
	}{
		{
			name: "Roihan",
			request: "Roihan",
			expected: "Hello Roihan",
		},
		{
			name: "budi",
			request: "budi",
			expected: "Hello budi",
		},
	}
	for _, test := range test {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}


func TestSubTest(t *testing.T)  {
	t.Run("Roi", func(t *testing.T) {
		result := HelloWorld("Roihan")
		require.Equal(t, "Hello Roihan", result, "Result must be 'Hello Roihan'")
	})
	t.Run("rafli", func(t *testing.T) {
		result := HelloWorld("rafli")
		require.Equal(t, "Hello rafli", result, "Result must be 'Hello rafli'")
	})
}

func TestMain(m *testing.M)  {
	// Before
	fmt.Println("Before Unit Test")

	m.Run()

	// after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T)  {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not runing on MAC OS")
	}
	result := HelloWorld("Roihan")
	require.Equal(t, "Hello Roihan", result, "Result must be 'Hello Roihan'")
}

func TestHelloWorldRequire(t *testing.T)  {
	result := HelloWorld("Roihan")
	require.Equal(t, "Hello Roihan", result, "Result must be 'Hello Roihan'")
	fmt.Println("TestHelloWorldRequire")
}

func TestHelloWorldAssert(t *testing.T)  {
	result := HelloWorld("Roihan")
	assert.Equal(t, "Hello Roihan", result, "Result must be 'Hello Roihan'")
	fmt.Println("TestHelloWorldAssert")
}

func TestHelloWorld(t *testing.T)  {
	result := HelloWorld("Roihan")

	if result != "Hello Roihan" {
		// error
		// panic("Result is not 'Hello Roihan'")
		// t.Fail()
		t.Error("Result must be 'Hello Roihan'")
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldRoi(t *testing.T)  {
	result := HelloWorld("Roi")

	if result != "Hello Roi" {
		// error
		// panic("Result is not 'Hello Roi'")
		// t.FailNow()
		t.Fatal("result must be 'Hello Roi'")
	}
	fmt.Println("TestHelloWorldRoi Done")
}