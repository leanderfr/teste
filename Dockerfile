

FROM golang:1.24.2 AS build

#   ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
#   ENV ZONEINFO /zoneinfo.zip


WORKDIR /app



COPY go.mod go.sum ./
RUN go mod download

COPY . /app
COPY .env /app
COPY leanderdev.key /app
COPY leanderdev.pem /app
copy leanderdev.crt /app

# Install the ca-certificate package
#RUN apt-get update && apt-get install -y ca-certificates

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go


FROM scratch

WORKDIR /app


COPY --from=build /app/leanderdev.crt /etc/ssl/certs/
COPY --from=build /app/api ./
COPY --from=build /app/.env ./

COPY --from=build /app/leanderdev.key ./
COPY --from=build /app/leanderdev.pem ./


# Update the CA certificates in the container
#RUN update-ca-certificates

EXPOSE 8445

CMD [ "./api" ]

#   docker build --rm -t rentacar_backend_golang .

# docker run  -d -p 8070:8070 --name rentacar_backend_golang rentacar_backend_golang




