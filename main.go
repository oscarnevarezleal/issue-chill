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
		githubactions.Fatalf("Not such issue")
	}

	issueTitle := *issue.Title
	issueBody := *issue.Body

	fmt.Printf("%s", issueTitle)
	fmt.Printf("%s", issueBody)

	parsedtext := sentitext.Parse(issueBody, lexicon.DefaultLexicon)
	sentiment := sentitext.PolarityScore(parsedtext)
	fmt.Printf("::set-output name=POS::%f\n", sentiment.Positive)
	fmt.Printf("::set-output name=NEG::%f\n", sentiment.Negative)
	fmt.Printf("::set-output name=NEU::%f\n", sentiment.Neutral)
	fmt.Printf("::set-output name=CMP::%f\n", sentiment.Compound)
}
