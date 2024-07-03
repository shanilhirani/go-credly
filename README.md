# Go Credly API Data Fetcher

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/shanilhirani/go-credly)](https://github.com/shanilhirani/go-credly/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/shanilhirani/go-credly.svg)](https://pkg.go.dev/github.com/shanilhirani/go-credly)
[![go.mod](https://img.shields.io/github/go-mod/go-version/shanilhirani/go-credly)](go.mod)
[![LICENSE](https://img.shields.io/github/license/shanilhirani/go-credly)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shanilhirani/go-credly/build.yml?branch=main)](https://github.com/shanilhirani/go-credly/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/shanilhirani/go-credly)](https://goreportcard.com/report/github.com/shanilhirani/go-credly)
[![Codecov](https://codecov.io/gh/shanilhirani/go-credly/branch/main/graph/badge.svg)](https://codecov.io/gh/shanilhirani/go-credly)

This repository contains a Go application that fetches public data from Credly's API which results in a users badges on Credly.

## Prerequisites

- [Go](https://golang.org/dl/) installed (version 1.22 or later)
- Internet connection

## Installation

Clone the repository:

```sh
git clone https://github.com/yourusername/go_credly.git
cd go_credly
```

## Usage

To run the application, execute the following command:
go run main.goshThe application will fetch data from the API and print it to the standard output.

## Configuration

You can configure the API endpoint by modifying the code in `main.go`.
const apiUrl = "https://api.example.com/data"go## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for improvements or bug fixes.

Happy Coding!
