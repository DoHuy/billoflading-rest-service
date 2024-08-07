FROM golang:1.14.1 as builder

ARG PROJECT=vtp-apis
ARG COMMIT_HASH_SHORT
WORKDIR /app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main \
    -ldflags="-X '$PROJECT/packages/constants.ServiceCode=$PROJECT'\
              -X '$PROJECT/packages/constants.CommitHashShort=$COMMIT_HASH_SHORT'"\
    ./cmd

FROM golang:1.14.1-alpine
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

