FROM golang:1.16-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build project.go
CMD ["/app/project"]