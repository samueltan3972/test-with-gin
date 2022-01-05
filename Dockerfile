FROM golang:alpine

#ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src

COPY . /go/src
# COPY templates /go/src/go-docker-dev.to/templates

# Run the two commands below to install git and dependencies for the project. 
# RUN apk update && apk add --no-cache git
# RUN go get ./...

# COPY dependencies /go/src #if you don't want to pull dependencies from git 

RUN go build -o ./app ./main.go

EXPOSE $PORT

ENTRYPOINT ["./app"]