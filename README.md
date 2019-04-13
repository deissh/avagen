
![Example](https://i.imgur.com/4tk6a5m.png)

# Avagen
[![CircleCI](https://circleci.com/gh/deissh/avagen.svg?style=svg)](https://circleci.com/gh/deissh/avagen)
![GitHub](https://img.shields.io/github/license/deissh/avagen.svg?style=flat-square)
![GitHub tag (latest by date)](https://img.shields.io/github/tag-date/deissh/avagen.svg?style=flat-square)

Generate avatars with initials from names.

## http benchmark

```bash
Bombarding http://127.0.0.1:8080?name=q&type=png for 10s using 1000 connection(s)
[==========================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     79390.31   18672.76  212498.18
  Latency       12.59ms    22.79ms      1.26s
  HTTP codes:
    1xx - 0, 2xx - 789992, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   109.92MB/s
```

## usage

### http example

```bash
$ chmod +x ./avagen
$ ./avagen
```

Path: `/`
Desc: Generate new avatars or return from cache

Params
 * `name` Nick or user name
 * `length` Max initials in avatar
 * `size` Avatar size (default 128px)
 * `type` Avatar type (png or jpeg)

### lib Example

```bash
go get github.com/deissh/avagen/...
```
```go
import  "github.com/deissh/avagen/pkg/avatar"

a := avatar.New("/path/to/fontfile")
b, _ := a.DrawToBytes("Deissh", 128, 2, "png")
```

## building

```bash
$ go get -insecure ./...
$ go build -v
```

## testing
Run all tests
```bash
go test ./...
```

## deploing

```bash
$ docker run --rm -p 8080:8080 deissh/avagen:latest
```

## contributing
- Fork the project
- Create a topic branch (preferably the in the gitflow style of feature/, hotfix/, etc)
- Make your changes and write complimentary tests to ensure coverage.
- Submit Pull Request once the full test suite is passing.
- Pull Requests will then be reviewed by the maintainer and the community and hopefully merged!
