# district-21

District 21 Game

## CI/CD

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
    fmt --> build
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

## Colaborate

**Include Fix**
fix(#issueId): `include a comment`

**Include Feature**
feat(#issueId): `Include an commit change`

**Improbe documentation**
doc(#issueid): `Include an change`

For example
    doc(#1): [README.md] Include CI and CD Workflows
