# tiny
__tiny__ repo will include several parts of the same project (trying to put as many containers on the same host):

1. **_tiny.go_** is a Golang small program to return at http call the port name as assigned as command line parameter.
2. **_build.sh_** is the command line to run to build the _tiny_ static binary.
3. **_tiny_** is the result of the _tiny.go_ static compiled using [Musl](http://www.musl-libc.org) and Golang 1.5 inside [golang:1.5-alpine container](https://hub.docker.com/_/golang/).
4. **_tiny.sha256sum_** is the SHA256SUM file for my static compiled file (you are free to regenerate yours).
5. **_Dockerfile_** needed to generate the [rolinux/tiny](https://hub.docker.com/r/rolinux/tiny/) Docker container.

Further information on how to use the Docker container will be provided in future commits.

First attempt was to hardcode the port name and return the hostname (container assigned ID) over http call. Unfortunately the default docker0 bridge or the bridge one has an 1024 limit (not allowing you more than 1023 containers).
Second attempt is to assign the port dynamically so we can use '-net=host' and avoid any bridge setup.
