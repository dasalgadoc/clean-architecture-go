FROM golang:1.20.0-alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/dasalgadoc/clean-architecture-go
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/clean-architecture-go ./cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/clean-architecture-go /go/bin/clean-architecture-go
ENTRYPOINT ["/go/bin/clean-architecture-go"]