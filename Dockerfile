# Start by building the application.
FROM golang:1.13-buster as build
WORKDIR /go/src/app
ADD . /go/src/app
RUN go build -ldflags "-s -w" /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /
CMD ["/app"]

