FROM golang:1.22.5-alpine
WORKDIR /code
COPY . /code
RUN go install github.com/air-verse/air@latest
RUN go mod download
RUN go mod tidy

EXPOSE 4500
CMD air server --port 4500 -c .air.toml 
