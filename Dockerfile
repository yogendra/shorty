# Start by building the application.
FROM golang:1.13-buster as build
WORKDIR /go/src/app
ADD . /go/src/app
RUN go build -ldflags "-s -w" -o /go/bin/app .

FROM scratch
COPY --from=build /go/bin/app /
CMD ["/app"]

