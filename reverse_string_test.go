package reverse_string

import (
	"testing"
)

type test struct {
	testString string
	expString  string
}

func TestReverse(t *testing.T) {
	tests := map[string]test{
		"simple string":          {testString: "Hello", expString: "olleH"},
		"empty string":           {testString: "", expString: ""},
		"one letter string":      {testString: "a", expString: "a"},
		"two letter string":      {testString: "no", expString: "on"},
		"string with whitespace": {testString: "\t whitespace   ", expString: "   ecapsetihw \t"},
		"cyrillic string":        {testString: "Привет", expString: "тевирП"},
		"string with emojis":     {testString: "I like tennis 🤚🎾😀", expString: "😀🎾🤚 sinnet ekil I"},
		"multiple strings": {
			testString: `This is the first line
This one is the second line
Finally, a third line`,
			expString: `enil driht a ,yllaniF
enil dnoces eht si eno sihT
enil tsrif eht si sihT`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			res := ReverseString(tt.testString)
			if tt.expString != res {
				t.Errorf("error: got %s, want %s", res, tt.expString)
			}
		})
	}
}
