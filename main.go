package main

import (
	"fmt"

	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
)
import "github.com/sethvargo/go-githubactions"

func main() {
	body := githubactions.GetInput("body")
	if body == "" {
		githubactions.Fatalf("missing input 'body'")
	}
	githubactions.AddMask(body)

	mytext := body
	parsedtext := sentitext.Parse(mytext, lexicon.DefaultLexicon)
	sentiment := sentitext.PolarityScore(parsedtext)
	fmt.Printf("::set-env name=POS::%f", sentiment.Positive)
	fmt.Printf("::set-env name=NEG::%f", sentiment.Negative)
	fmt.Printf("::set-env name=NEU::%f", sentiment.Neutral)
	fmt.Printf("::set-env name=CMP::%f", sentiment.Compound)
}