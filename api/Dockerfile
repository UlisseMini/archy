FROM golang:latest

COPY src /src
WORKDIR /src
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM scratch

COPY --from=0 /main /main

CMD ["/main"]
