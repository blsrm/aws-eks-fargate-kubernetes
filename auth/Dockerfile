FROM golang:latest AS compiler

RUN mkdir /app
COPY . /app
WORKDIR /app
ENV GO111MODULE=on
RUN go mod download

RUN go build -o main .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o auth


FROM golang:1.14.4-alpine
WORKDIR /root
EXPOSE 3000
COPY --from=compiler /app/auth .
CMD ["./auth"]