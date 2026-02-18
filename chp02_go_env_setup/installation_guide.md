# Chapter 02: Go Environment Setup
## Installation Guide
* Step 1: Go to https://go.dev/dl/ 
* Step 2: Select the featured downloads based on your machine.
* Step 3: At the terminal or CLI, type "go". There will be a list of go command shown in the terminal.
* Step 4: If showing system error, this is because that the go not been added in the path through the environment variable. [Solution](https://go.dev/wiki/SettingGOPATH)


## Running Go Scripts
* Step 1: In the terminal, type `go run demo.go`. This code snippet will create a temporary `demo.exe`. After that, execute the `.exe` file.
* Step 2: This also can be done by running the `go build demo.go` to build the compiled version of `demo.exe` first. Then at the terminal, just run `demo.exe`. 

## Additional notes
* In Go, the usage is strict, if the package imported not been used. It will display error. Stricter == More accurate/catch bug before running in the production environment.