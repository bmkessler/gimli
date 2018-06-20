package gimli

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var testVectors = []struct {
	input  string
	output string
}{

	{
		"There's plenty for the both of us, may the best Dwarf win.",
		"4afb3ff784c7ad6943d49cf5da79facfa7c4434e1ce44f5dd4b28f91a84d22c8",
	},
	{
		"If anyone was to ask for my opinion, which I note they're not, I'd say we were taking the long way around.",
		"ba82a16a7b224c15bed8e8bdc88903a4006bc7beda78297d96029203ef08e07c",
	},
	{
		"Speak words we can all understand!",
		"8dd4d132059b72f8e8493f9afb86c6d86263e7439fc64cbb361fcbccf8b01267",
	},
	{
		"It's true you don't see many Dwarf-women. And in fact, they are so alike in voice and appearance, that they are often mistaken for Dwarf-men.  And this in turn has given rise to the belief that there are no Dwarf-women, and that Dwarves just spring out of holes in the ground! Which is, of course, ridiculous.",
		"ebe9bfc05ce15c73336fc3c5b52b01f75cf619bb37f13bfc7f567f9d5603191a",
	},
	{
		"", // empty string
		"b0634b2c0b082aedc5c0a2fe4ee3adcfc989ec05de6f00addb04b3aaac271f67",
	},
}

func TestHash(t *testing.T) {
	for i, tt := range testVectors {
		want := tt.output
		outputLen := len(want) / 2
		output := make([]byte, outputLen)
		Hash(output, []byte(tt.input))
		got := hex.EncodeToString(output)
		if got != want {
			t.Errorf("testVector #%d\ninput: %q\ngot:   %s\nwant:  %s\n", i+1, tt.input, got, want)
		}
		// test variable output length
		for j := 0; j < outputLen; j++ {
			want := hex.EncodeToString(output[:j])
			varOutput := make([]byte, j)
			Hash(varOutput, []byte(tt.input))
			got := hex.EncodeToString(varOutput)
			if got != want {
				t.Errorf("testVector #%d\ninput: %q\noutput bytes: %d\ngot:   %s\nwant:  %s\n", i+1, tt.input, j, got, want)
			}
		}

	}
}

func BenchmarkHash(b *testing.B) {
	for _, tt := range testVectors {
		input := []byte(tt.input)
		for _, outputLen := range []int{16, 32, 64} {
			output := make([]byte, outputLen)
			b.Run(fmt.Sprintf("%d/%d", len(input), len(output)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					Hash(output, input)
				}
			})
		}
	}
}
