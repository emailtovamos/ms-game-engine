# 1st stage: build Go binary

FROM golang:1.10
# TODO: Insert your repo name
WORKDIR /go/src/github.com/teach/ms-game-engine/

# Copy only Go package directories each separately
# then `docker build` does not
COPY vendor ./vendor/
COPY cli ./cli/
COPY internal ./internal/

RUN CGO_ENABLED=0        \
    GOOS=linux           \
    go install           \
      -a                 \
      -installsuffix cgo \
      --ldflags="-s"     \
      ./cli/server


# 2nd stage: embed Go binary in damn small Linux distro (== Alpine)

FROM alpine:latest
WORKDIR /app/
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata 

# Copy the binary from the first build stage
COPY --from=0 /go/bin/ app

EXPOSE 8000

CMD app/server

