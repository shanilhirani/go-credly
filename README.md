# Go Credly API Data Fetcher

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/shanilhirani/go-credly)](https://github.com/shanilhirani/go-credly/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/shanilhirani/go-credly.svg)](https://pkg.go.dev/github.com/shanilhirani/go-credly)
[![go.mod](https://img.shields.io/github/go-mod/go-version/shanilhirani/go-credly)](go.mod)
[![LICENSE](https://img.shields.io/github/license/shanilhirani/go-credly)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/shanilhirani/go-credly/build.yml?branch=main)](https://github.com/shanilhirani/go-credly/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/shanilhirani/go-credly)](https://goreportcard.com/report/github.com/shanilhirani/go-credly)
[![Codecov](https://codecov.io/gh/shanilhirani/go-credly/branch/main/graph/badge.svg)](https://codecov.io/gh/shanilhirani/go-credly)

## About

Go Credly is a Go App which enables users to obtain Certification Badges earned on Credly's Certification Platform.

If your Credly Badges are made public then all you need is your Credly `username` and you'll be able to pull your badges programmatically.

### Use case

- You could use this tool to dynamically update a CV/Resume, Portfolios and Personal Websites using a Github Action, or just running the binary on a cron.

## Prerequisites

- Credly with Public Badges
- [Go](https://golang.org/dl/) installed (version 1.22 or later)
- Internet connection

## Installation

Clone the repository:

```bash
git clone https://github.com/yourusername/go_credly.git
cd go_credly
```

## Usage

To run the application, execute the following command:

`go run main.go <yourcredlyusername>`

Then _go-credly_ will attempt to fetch data from Credly's API and return the result in JSON to standard out.

## Configuration

TBC

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for improvements or bug fixes.

Happy Coding!
