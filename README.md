# Issue Chill Action

## About

This action reads a GitHub Issue and outputs its Sentiment Analysis, so you can detect whether its content is redacted in a way that make it sound positive or negative.

## Example

Create a `.github/workflows/issue-sentiment.yml` file in your repository with the following content.
Make sure to replace <owner> and <repo> with your repository values.

```yaml
# .github/workflows/issue-sentiment.yml

name: Issue Sentiment Analysis
on:
  issues:
    types: [opened, edited]

jobs:  
  add-comment:
    runs-on: ubuntu-latest
    permissions:
      issues: write
    steps:
      - name: Sentiment Analysis
        id: sa
        uses: oscarnevarezleal/issue-chill@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
        with:
          owner: <org>
          repo: <repo>
          issue: ${{ github.event.issue.number }}

      - name: Add positive comment
        if: ${{ startsWith(steps.sa.outputs.IS_50_POS, 'true') }} # 60, 70, 80 and 90 also available
        uses: peter-evans/create-or-update-comment@a35cf36e5301d70b76f316e867e7788a55a31dae
        with:
          issue-number: ${{ github.event.issue.number }}
          body: |
            This issue is available for anyone to work on. **Make sure to reference this issue in your pull request.** :sparkles: Thank you for your contribution! :sparkles:

      - name: Add negative reaction
        if: ${{ startsWith(steps.sa.outputs.IS_70_NEG, 'true') }} # 50, 60, 70, 80 and 90 also available
        uses: peter-evans/create-or-update-comment@a35cf36e5301d70b76f316e867e7788a55a31dae
        with:
          issue-number: ${{ github.event.issue.number }}
          body: |
            Please refer to the Code Of Conduct.
            > "This Aggression Will Not Stand, Man."
```

## Tech

- **Go libraries**
  - [github.com/google/go-github/v40/github](github.com/google/go-github/v40/github)
  - [github.com/grassmudhorses/vader-go/lexicon](github.com/grassmudhorses/vader-go/lexicon)
  - [github.com/grassmudhorses/vader-go/sentitext](github.com/grassmudhorses/vader-go/sentitext)