package main

import (
	"fmt"

	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)
import "github.com/sethvargo/go-githubactions"

func main() {
	content := githubactions.GetInput("content")
	if content == "" {
		githubactions.Fatalf("missing input 'content'")
	}
	githubactions.AddMask(content)

	parsedtext := sentitext.Parse(content, lexicon.DefaultLexicon)
	sentiment := sentitext.PolarityScore(parsedtext)
	fmt.Printf("::set-output name=POS::%f", sentiment.Positive)
	fmt.Printf("::set-output name=NEG::%f", sentiment.Negative)
	fmt.Printf("::set-output name=NEU::%f", sentiment.Neutral)
	fmt.Printf("::set-output name=CMP::%f", sentiment.Compound)
}
