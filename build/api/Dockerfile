FROM golang:1.11-alpine3.8

# Set environment variables
# App
ENV APP_PATH /data

# Install git
RUN apk update && apk upgrade && \
    apk add --no-cache git

# Create src dir
RUN mkdir $GOPATH/${APP_PATH}

# Set workdir to back
WORKDIR $GOPATH/${APP_PATH}

# Install a hot reloading tool for our server
RUN go get github.com/githubnemo/CompileDaemon

# Run the server
CMD ["sh", "start.sh"]
