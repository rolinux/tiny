# tiny

Using the [Building Multi-Arch Images for Arm and x86 with Docker Desktop](https://engineering.docker.com/2019/04/multi-arch-images/) documentation I was able to build images for:

* linux/amd64
* linux/arm64
* linux/arm/v7
* linux/s390x
* linux/386

Used the `docker buildx build --platform linux/amd64,linux/arm64,linux/s390x,linux/386,linux/arm/v7 -t rolinux/tiny:latest --push .` command.

Testing can be done by running `docker run --rm -p 8080:8080 docker.io/rolinux/tiny:latest@sha256:9001d549bc1f140b5571a00c07fbd327002a367d75a7da453497ade3f60cef89` on the Docker Desktop (Edge) as it is able to use the multi arch emulation.

**_tiny.go_** is a Golang small program to return at http call the port name as assigned as command line parameter or 8080 that is default. It will also return the architecture.
