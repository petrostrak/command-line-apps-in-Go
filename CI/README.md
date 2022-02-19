#### Continuous Integration (CI)

A useful implementation of a Continuous Integration (CI) tool for your Go
programs, using the `os/exec` package. A typical CI pipeline consists of several automated steps that
continuously ensure a code base or an application is ready to be merged
with some other developer’s code, usually in a shared version control
repository.

For this example, the CI pipeline consists of:
* Building the program using `go build` to verify if the program structure is
valid.
* Executing tests using `go test` to ensure the program does what it’s intended to do.
* Executing `gofmt` to ensure the program’s format conforms to the standards.
* Executing `git push` to push the code to the remote shared Git repository that hosts the program code.