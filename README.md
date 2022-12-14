# district-21

![Coverage](https://img.shields.io/badge/Coverage-90.0%25-brightgreen)

District 21 Game

## Testing

> go test ./... -v

If we want test with coverage

> go test ./... -test.coverprofile coverage
> go test ./... -cover

## CI/CD

CI Tool Chain For Go

For my personal code I use GitHub as a repository, so assuming that the tools
I’ll employ are:

- GitHub Actions: Automate the workflow
- Codecov.io: A coverage dashboard
- Readme badges: Easy visibility dashboard on CI status

- Templates CI/CD for golang projects
  
### CI-WORKFLOW [ ./gihub/workspaces/ci-go.yml ]

Start with any events on

- Push:
  - On branches `main` or `feature`
- Any pull_request

```mermaid
stateDiagram
    direction LR
    [*] --> checkout@v3
    checkout@v3 --> setup_go@v3
    setup_go@v3 --> vet
    vet --> fmt
    fmt --> test
    test --> go_coverage_badge
    go_coverage_badge --> verify_changed_files
    verify_changed_files --> push_changes
    push_changes --> build
    build --> [*]
```

### LINT Golang [ ./github/workspaces/golangci-lint]

Start with

- Push:
  - On Tag with name as `v*`
  - On branches `main` or `feature`
- Any pull_request

```mermaid
stateDiagram
    direction LR
    [*] --> checkout@v3
    checkout@v3 --> golangci_lint
    golangci_lint --> [*]
```

### CD Workflow [ ./github/workspaces/release.yml ]

- Push:
  - On Tag with name as `v*`
  - On only branches `main`
- Any pull_request

```mermaid
stateDiagram
    direction LR
    [*] --> checkout@v3
    checkout@v3 --> setup_go@v3
    setup_go@v3 --> goreleaser
    goreleaser --> [*]
```

### WellCome workflow

```yml
jobs:
  greeting:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/first-interaction@v1.1.0
      with:
        repo-token: ${{ secrets.GITHUB_TOKEN }}
        issue-message: "Thanks for reporting this issue, don't forget to star this project if you haven't already to help us reach a wider audience."
        pr-message: "Thanks for implementing a fix, could you ensure that the test covers your changes if applicable."
```

### TODOs

- [coverage-badge-go](https://github.com/tj-actions/coverage-badge-go/tree/main/.github/workflows)
- [how-to-set-up-a-test-coverage-threshold](https://medium.com/synechron/how-to-set-up-a-test-coverage-threshold-in-go-and-github-167f69b940dc)

## Colaborate

**Include Fix**
fix(#issueId): `include a comment`

**Include Feature**
feat(#issueId): `Include an commit change`

**Improbe documentation**
doc(#issueid): `Include an change`

For example
    doc(#1): [README.md] Include CI and CD Workflows
