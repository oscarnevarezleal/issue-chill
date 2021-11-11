package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/google/go-github/v40/github"
	"github.com/grassmudhorses/vader-go/lexicon"
	"github.com/grassmudhorses/vader-go/sentitext"
	"golang.org/x/oauth2"
)
import "github.com/sethvargo/go-githubactions"

func main() {
	// content := githubactions.GetInput("content")
	// if content == "" {
	// 	githubactions.Fatalf("missing input 'content'")
	// }
	// githubactions.AddMask(content)

	owner := githubactions.GetInput("owner")
	if owner == "" {
		githubactions.Fatalf("missing input 'owner'")
	}

	repo := githubactions.GetInput("repo")
	if repo == "" {
		githubactions.Fatalf("missing input 'repo'")
	}

	issueStr := githubactions.GetInput("issue")
	if issueStr == "" {
		githubactions.Fatalf("missing input 'issue'")
	}

	issueId, err := strconv.Atoi(issueStr)
	if err != nil {
		return
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	issue, _, err := client.Issues.Get(ctx, owner, repo, issueId)

	if err != nil {
		githubactions.Fatalf("An error occurred: %s", err.Error())
	}

	issueTitle := *issue.Title
	issueBody := *issue.Body

	fmt.Printf("::set-output name=ISSUE_TITLE::%s\n", issueTitle)

	parsedtext := sentitext.Parse(issueBody, lexicon.DefaultLexicon)
	sentiment := sentitext.PolarityScore(parsedtext)

	fmt.Printf("::set-output name=POS::%f\n", sentiment.Positive)
	fmt.Printf("::set-output name=IS_50_POS::%s\n", strconv.FormatBool(sentiment.Positive > 0.5))
	fmt.Printf("::set-output name=IS_60_POS::%s\n", strconv.FormatBool(sentiment.Positive > 0.6))
	fmt.Printf("::set-output name=IS_70_POS::%s\n", strconv.FormatBool(sentiment.Positive > 0.7))
	fmt.Printf("::set-output name=IS_80_POS::%s\n", strconv.FormatBool(sentiment.Positive > 0.8))
	fmt.Printf("::set-output name=IS_90_POS::%s\n", strconv.FormatBool(sentiment.Positive > 0.9))

	fmt.Printf("::set-output name=NEG::%f\n", sentiment.Negative)
	fmt.Printf("::set-output name=IS_50_NEG::%s\n", strconv.FormatBool(sentiment.Negative > 0.5))
	fmt.Printf("::set-output name=IS_60_NEG::%s\n", strconv.FormatBool(sentiment.Negative > 0.6))
	fmt.Printf("::set-output name=IS_70_NEG::%s\n", strconv.FormatBool(sentiment.Negative > 0.7))
	fmt.Printf("::set-output name=IS_80_NEG::%s\n", strconv.FormatBool(sentiment.Negative > 0.8))
	fmt.Printf("::set-output name=IS_90_NEG::%s\n", strconv.FormatBool(sentiment.Negative > 0.9))

	fmt.Printf("::set-output name=NEU::%f\n", sentiment.Neutral)
	fmt.Printf("::set-output name=CMP::%f\n", sentiment.Compound)
}
