# Testing

Makefile targets provided

## Command

Run tests, optionally for a specific package or matching a specific pattern
```sh
make test
```

```sh
make test PKG=./sorting/
```

```sh
make test RUN=MergeSort
```

## Optional Arguments

| Argument | Description | Example |
|---------|-------------|---------|
| `PKG` | Run tests for a specific package instead of all packages | `PKG=./sorting` |
| `RUN` | Run a specific test (supports Go test regex patterns) | `RUN=MergeSort` |

## Examples

### Run all tests

```sh
make test
```

### Run tests for a specific package

```sh
make test PKG=./sorting
```

### Run a specific test

```sh
make test RUN=MergeSort
```

### Run a specific test in a specific package

```sh
make test PKG=./sorting RUN=MergeSort
```
