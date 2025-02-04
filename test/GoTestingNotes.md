File naming convention for test is orginial name + _test.go
With one test file for code file

"github.com/stretchr/testify/assert" for assert functions

functions must begin with the word Test with a capital T followed by the name of what you are testing

run tests with command go test

Sample function layout for a test

func TestFunction(t *testing.T) {

    type Testcase struct {
        //input value
        value string
        //expected values
        expected string
        err error
    }

    t.Run("function", func(t *testing.go)) { 

        tests := []testCase{
            {value:"test", expected: "test finished"},
            {value:"test", expected: "test finished"},
        }
    }

    for _, test := range tests {
        actual, err := functiontobetested(test.value)
        assert.NoError(t, err)
        assert.Equal(t, test.expected, actual)
    }
}

for mocking out apicalls and other external functions, separate the call from the function and pass it as a variable