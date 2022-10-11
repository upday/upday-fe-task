FROM golang:1.16.4 AS build
WORKDIR /go/src
COPY . .
COPY cmd/cms/main.go .

ENV CGO_ENABLED=0
RUN go mod vendor
RUN go build -a -installsuffix cgo -o cms .

FROM scratch AS runtime
COPY --from=build /go/src/cms ./
EXPOSE 8080/tcp
ENTRYPOINT ["./cms"]
