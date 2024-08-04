
# Dockerfile Linter

This tool is a simple Dockerfile linter written in Golang. It checks for common Dockerfile issues and provides suggestions for improvements.

## Features

The linter currently supports the following rules:

1. **FROM Instruction**: Ensures that the `FROM` instruction has at least one argument.
2. **RUN Instruction**: 
   - Checks for multiple consecutive `RUN` instructions and suggests combining them into one.
   - Suggests using `WORKDIR` instead of `RUN cd` for changing directories.
3. **ENTRYPOINT Instruction**: Ensures that `ENTRYPOINT` arguments are provided in JSON notation.
4. **CMD Instruction**: Ensures that `CMD` arguments are provided in JSON notation.
5. **Unknown Instructions**: Warns about unknown instructions or those without provided rules.

## Installation

To build and run the linter, you need to have Golang installed. Follow these steps:

1. **Clone the Repository**:
    ```sh
    git clone <repository-url>
    cd dockerfile-linter
    ```

2. **Build the Linter**:
    ```sh
    go build -o dockerfile-linter .
    ```

3. **Run the Linter**:
    ```sh
    ./dockerfile-linter path/to/Dockerfile
    ```

## Usage

To lint a Dockerfile, simply pass the path of the Dockerfile to the linter:

```sh
./dockerfile-linter path/to/Dockerfile
```

The linter will output any issues found in the Dockerfile along with suggestions for fixing them.

## Example

Consider the following Dockerfile:

```Dockerfile
FROM ubuntu:latest
RUN apt-get update
RUN apt-get install -y vim
RUN cd /app
ENTRYPOINT echo "Hello, world"
CMD echo "This is a test"
```

Running the linter on this Dockerfile will produce the following output:

```
Multiple consecutive RUN instructions. do this 'RUN download_a_really_big_file && \ 
    remove_the_really_big_file'
Make use of WORKDIR instead of RUN cd 'some-path'
Use JSON notation for ENTRYPOINT arguments. example 'ENTRYPOINT ['foo', 'run-server']'
Use JSON notation for CMD arguments. example 'CMD ['foo', 'run-server']'
```

## Contributing

Contributions are welcome! If you have any suggestions or find any bugs, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.