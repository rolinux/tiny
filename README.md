# tiny

Using the [Building Multi-Arch Images for Arm and x86 with Docker Desktop](https://engineering.docker.com/2019/04/multi-arch-images/) documentation and later on GitGub Actions I was able to build:

* linux/amd64
* linux/arm64
* linux/arm/v7

And binaries.

**_tiny.go_** is a Golang small program to return at http call the port name as assigned as command line parameter or 8080 that is default. It will also return the architecture.
