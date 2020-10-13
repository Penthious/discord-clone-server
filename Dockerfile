FROM golang:latest as BASE
WORKDIR /app
COPY ./ /app
RUN go mod download

FROM BASE as dev
RUN go get github.com/codegangsta/gin
## Add the wait script to the image
ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait
CMD /wait && gin run server

FROM BASE as PROD
RUN go build main.go
