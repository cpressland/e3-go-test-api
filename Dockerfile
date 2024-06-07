FROM docker.io/golang:1.22 as build

WORKDIR /go/src
ADD . .

RUN CGO_ENABLED=0 GOOS=linux go build -o e3-go-test-api

FROM scratch
COPY --from=build /go/src/e3-go-test-api /e3-go-test-api
ENTRYPOINT ["/e3-go-test-api"]
EXPOSE 8080
