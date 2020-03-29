# Secret

Cli encrypted caching tool for storing key-value pairs in encrypted file.

## Prerequsites

* Clone the project
* Create a file named `secret_`
* Set `$GOBIN` variable

## Install

Run `go install` to build the project and install binary.

## Usage

Get command: `secret get test -k 123412341234`

Set command: `secret set test test_value -k 123412341234`


## Todo

- [x] Write tests
- [ ] Store the file within binary
- [ ] Command input validation
- [ ] Error handling 