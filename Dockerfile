FROM golang:1.18
WORKDIR /app
COPY ./published ./
CMD ["./main"]