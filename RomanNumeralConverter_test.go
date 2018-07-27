package romanNumerals

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput int
	}{

		{
			name:           "given I, returns 1",
			input:          "I",
			expectedOutput: 1,
		},
	}

	runConvertTestCases(tests, t)
}

func runConvertTestCases(tests []struct {
	name           string;
	input          string;
	expectedOutput int
}, t *testing.T) {
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			output := convert(testCase.input)

			if output != testCase.expectedOutput {
				t.Fatalf("%v - Expected < %v > but got < %v >", testCase.name, testCase.expectedOutput, output)
			}
		})
	}
}
