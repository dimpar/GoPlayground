# Go-processes

Go library package that provides APIs to run external processes for unix family.

## Prerequisites

It is required to operate on unix family machine.

### Getting Started

API consist of the following interfaces:

- Start() - starts a process

- Read() - read from stdout and stderr

- Stop() - stops a process

### Installing

Import to a package where this API will be used

```
import "github.com/dimpar/go-processes/pkg/processes" 
```

### Example

Examples of running this API can be found in tests and main package.


## Running the tests

```
go test ./...
```