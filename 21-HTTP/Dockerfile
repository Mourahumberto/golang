FROM golang

WORKDIR /app
ADD . /app
RUN go build -o http
ENTRYPOINT ./http