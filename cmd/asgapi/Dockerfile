# build stage
FROM golang:1.12 AS build-env
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/asgapi

# final stage
FROM scratch
COPY --from=build-env /go/bin/asgapi /go/bin/asgapi
EXPOSE 9797
CMD ["/go/bin/asgapi"]