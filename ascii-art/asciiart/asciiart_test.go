package asciiart

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	
	type parameters struct {
		words      []string
		arrWordArt []string
	}

	
	type teststruct struct {
		name string
		args parameters
	}

	File, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	WordArt := strings.Split(string(File), "\n")

	tests := []teststruct{
		{name: "empty string", args: parameters{words: []string{""}, arrWordArt: WordArt}},
		{name: "Hello world", args: parameters{words: []string{"Hello world"}, arrWordArt: WordArt}},
		{name: "1Hello 2There", args: parameters{words: []string{"1Hello 2There"}, arrWordArt: WordArt}},
		{name: "(Hello There)", args: parameters{words: []string{"(Hello There)"}, arrWordArt: WordArt}},
		{name: "\n", args: parameters{words: []string{"\n"}, arrWordArt: WordArt}},
		
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			
			CreateStringart(tc.args.words, tc.args.arrWordArt)
		})
	}
}
