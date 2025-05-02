FROM golang:1.22-alpine
WORKDIR /app
COPY backend/go.* ./backend/
WORKDIR /app/backend
RUN go mod download
COPY backend .
RUN go build -o /server ./cmd/server
CMD ["/server"]