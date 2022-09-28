FROM golang:1.18-alpine

WORKDIR /app

COPY . .

RUN go build -o rogerdev-blogspot-frontend

EXPOSE 8082

CMD ./rogerdev-blogspot-frontend
