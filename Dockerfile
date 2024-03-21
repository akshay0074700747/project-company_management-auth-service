FROM golang:1.21.5-bullseye AS build

RUN apt-get update

WORKDIR /app

COPY . .

RUN go mod download

WORKDIR /app/cmd

RUN go build -o auth-service

FROM busybox:latest

WORKDIR /auth-service

COPY --from=build /app/cmd/auth-service .

COPY --from=build /app/cmd/.env .

EXPOSE 50004

CMD [ "./auth-service" ]