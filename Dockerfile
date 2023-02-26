# build container
FROM golang as builder

ENV DISTRIBUTION_DIR /go/src/github.com/rolinux/tiny

WORKDIR $DISTRIBUTION_DIR
COPY . $DISTRIBUTION_DIR

RUN go mod tidy
RUN go mod download

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o tiny .
RUN chmod +x tiny

# run container with app on top on scratch empty container
FROM scratch

COPY --from=builder /go/src/github.com/rolinux/tiny/tiny /bin/tiny

EXPOSE 8080

CMD ["tiny"]
