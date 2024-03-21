# ray-tracing

## Test Frameworks
### [Ginkgo](https://onsi.github.io/ginkgo/#getting-started)

### [Bootstrapping a Suite](https://onsi.github.io/ginkgo/#bootstrapping-a-suite)
Generate a new test suite in the current directory:
```bash
cd path/to/books
ginkgo bootstrap
```

### Running Tests
- Run the tests in the current directory:
```bash
ginkgo
```
- Run all tests for the project:
```bash
go run github.com/onsi/ginkgo/v2/ginkgo -r --randomize-all --randomize-suites --fail-on-pending --keep-going --cover --coverprofile=cover.profile --race --trace --json-report=report.jso
```

### [Adding Specs to a Suite](https://onsi.github.io/ginkgo/#adding-specs-to-a-suite)
This will generate a test file named book_test.go in the books directory:
```bash
```shell
ginkgo generate book
```
