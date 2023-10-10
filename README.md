# JSON to Binary Converter and Size Comparator

This Go (Golang) project converts a large JSON file (24MB) into a binary file and compares the size of the generated binary file (4MB) at the end.

## Prerequisites

To use this project, ensure you have the following prerequisites installed on your system:

- Go (Golang): [Installation guide](https://golang.org/doc/install)

## Usage

1. Clone the repository:
   
   ```bash
   git clone https://github.com/jorgejr568/binary-vs-json-size.git
   cd binary-vs-json-size
   ```

2. Build the Go application:

   ```bash
   go run transform.go
   ```

3. Compare the sizes:

   ```bash
   du -h du -h covid-timeseries.*
   ```

## License

This project is licensed under the [MIT License](LICENSE).
