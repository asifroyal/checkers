# CheckEnv
CheckEnv is a command-line tool written in Go that scans specified directories for files with specified extensions and checks if they contain references to environment variables. If any environment variables are referenced but not defined, the tool outputs a message indicating which variables are missing. This tool can be useful for ensuring that all required environment variables are defined before running an application.

## Installation
To install CheckEnv, you can use `go get`:

```go
go get github.com/username/checkenv
```

## Usage
To use CheckEnv, you must specify the directories and file extensions to scan using the `-dirs` and `-exts` flags. You can also specify a comma-separated list of directories to ignore using the `-ignore` flag.

Here's an example usage:

```go
checkenv -dirs /path/to/directory -exts .js,.ts,.jsx,.tsx -ignore node_modules,vendor
```

This will scan the `/path/to/directory` directory for files with the `.js`, `.ts`, `.jsx`, or `.tsx` extensions, ignoring any files in the `node_modules` or `vendor` directories. If any environment variables are referenced but not defined in the scanned files, CheckEnv will output a message indicating which variables are missing.

## License
CheckEnv is licensed under the MIT License. See the LICENSE file for more information.