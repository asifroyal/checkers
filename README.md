# CheckEnv
CheckEnv is a command-line tool written in Go that scans specified directories for files with specified extensions and checks if they contain references to environment variables. If any environment variables are referenced but not defined, the tool outputs a message indicating which variables are missing. This tool can be useful for ensuring that all required environment variables are defined before running an application.

> Note: The concept for CheckEnv was inspired by an incident involving the accidental failure to define environment variables while deploying a Node.js application. It's important to recognize that while CheckEnv can be useful for small codebase, there are more efficient alternatives available.

## Installation
To install CheckEnv, you can use `go get`:

```go
go get github.com/asifroyal/checkenv
```

## Usage
To use CheckEnv, you must specify the directories and file extensions to scan using the `-dirs` and `-exts` flags. You can also specify a comma-separated list of directories to ignore using the `-ignore` flag.

Here's an example usage:

```go
checkenv -dirs /path/to/directory -exts .js,.ts,.jsx,.tsx -ignore node_modules,vendor
```

This will scan the `/path/to/directory` directory for files with the `.js`, `.ts`, `.jsx`, or `.tsx` extensions, ignoring any files in the `node_modules` or `vendor` directories. If any environment variables are referenced but not defined in the scanned files, CheckEnv will output a message indicating which variables are missing.


## Building and Running the Program

### Windows

To build the program on Windows, open a command prompt and navigate to the directory containing the Go source code. Then run:
```go
go build -o checkenv.exe
```

This will compile the code and generate an executable named `checkenv.exe`.

To run the program, make sure any required environment variables are set, then run:
```go
checkenv.exe -dirs C:\path\to\scan -exts .js,.ts
```

Replace the `-dirs` and `-exts` values with your desired directories and extensions.

### Linux

To build the program on Linux, open a terminal and navigate to the directory containing the Go source code. Then run:
```go
go build -o checkenv
```

This will compile the code and generate an executable named `checkenv`.

To run the program, make sure any required environment variables are set, then run:
```go
./checkenv -dirs /path/to/scan -exts .js,.ts
```

Replace the `-dirs` and `-exts` values with your desired directories and extensions.

The executable can be run from any directory after building. Make sure to include `./` before `checkenv` to execute it from the current directory.