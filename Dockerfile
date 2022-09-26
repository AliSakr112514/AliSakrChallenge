#Build stage
FROM golang:1.19.1-alpine3.16 as builder
WORKDIR /app
# download the required Go dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify 
COPY . .
RUN  cd cmd && go build -o ./out

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder  app/ .
EXPOSE 8080

CMD [ "./cmd/out" ]



