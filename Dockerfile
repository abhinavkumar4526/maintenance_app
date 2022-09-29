FROM golang:1.16-alpine

WORKDIR /app
COPY . .
RUN go build -o maintenance_build ./main.go
EXPOSE 80
CMD ["./maintenance_build"]


