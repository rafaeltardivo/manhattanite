FROM golang:1.14-buster as build

# build image
ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV POINTS_FILE_RELATIVE_PATH=../data/data.json 
ENV HTTP_SERVER_PORT=8080

RUN mkdir /build
ADD . /build/
WORKDIR /build/main
RUN go test -v ../utils ../services ../serializers
RUN go build -o app .

# final image
FROM gcr.io/distroless/base-debian10

ENV POINTS_FILE_RELATIVE_PATH=data.json
ENV HTTP_SERVER_PORT=8080
COPY --from=build /build/data/data.json .
COPY --from=build /build/main /
EXPOSE 8080
ENTRYPOINT ["/app"]
