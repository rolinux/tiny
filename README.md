# tiny
__tiny__ repo will include several parts of the same project (trying to put as many containers on the same host):
1. **_tiny.go_** is a Golang small program to return the hostname when connecting to port 8000. _It will try to be as small and simple as possible but still able to return the hostname when queried over http._
2. **_build.sh_** is the command line to run to build the _tiny_ static binary.
3. **_tiny_** is the result of the _tiny.go_ static compiled using [Musl](http://www.musl-libc.org) and Golang 1.5 inside [golang:1.5-alpine container](https://hub.docker.com/_/golang/).
4. **_tiny.sha256sum_** is the SHA256SUM file for my static compiled file (you are free to regenerate yours).
5. **_Dockerfile_** needed to generate the [rolinux/tiny](https://hub.docker.com/r/rolinux/tiny/) Docker container.

Further information on how to use the Docker container will be provided in future commits.
