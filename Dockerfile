

FROM golang:1.26.3 AS build

#   ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
#   ENV ZONEINFO /zoneinfo.zip


WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app
COPY privkey.pem /app
COPY fullchain.pem /app
#COPY leanderdev.crt /app

# Install the ca-certificate package
#RUN apt-get update && apt-get install -y ca-certificates

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM scratch

WORKDIR /app

#COPY --from=build /app/leanderdev.crt /etc/ssl/certs/

COPY --from=build /app/fullchain.pem ./
COPY --from=build /app/privkey.pem ./
COPY --from=build /app/api ./

# Update the CA certificates in the container
#RUN update-ca-certificates

EXPOSE 8443

CMD [ "./api" ]

#   docker build --rm -t https .

# docker run  -d -p 8443:8443 --name https https




