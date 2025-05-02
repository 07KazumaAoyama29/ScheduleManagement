FROM golang:1.22-alpine
WORKDIR /app
# install git so go modules can be fetched
RUN apk add --no-cache git

# copy go mod files and source
COPY backend/ ./

# ensure go.mod & go.sum are in sync and dependencies are cached
RUN go mod tidy

# build binary
RUN go build -o /server ./cmd/server
CMD ["/server"]