## Clos

### 3Musketeer Pattern

Pattern using Make, Docker and Docker Compose for consistent development environments. All commands run in containers.

#### Why

- Zero local dependencies (except Make, Docker, Docker Compose)
- Same commands work locally and in CI
- Portable across all platforms

#### Usage

```golang
make run
make test
make build
make deploy
```

For Bruno Collection
<https://github.com/usebruno/bruno>
This is to replace Postman because it is being ass with its pricing tiers
.brunoCollection is the directory to import, will try to keep it up to date with the audiobook version of it

Use this for file hosting I think https://catbox.moe/