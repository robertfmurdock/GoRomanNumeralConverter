package romanNumerals

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		name           string
		input          int
		expectedOutput string
	}{

		{
			name:           "will return I for 1",
			input:          1,
			expectedOutput: "I",
		},



	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			output := convert(testCase.input)

			if output != testCase.expectedOutput {
				t.Fatalf("Expected < %s > but got < %s >", testCase.expectedOutput, output)
			}
		})
	}
}
