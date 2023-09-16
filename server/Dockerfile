FROM golang:1.21.1 as builder

WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server


FROM scratch AS runtime
COPY --from=builder /app/server /server
CMD [ "/server" ]
