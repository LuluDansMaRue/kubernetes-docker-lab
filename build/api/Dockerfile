FROM golang:1.11-alpine3.8

# Set environment variables
# App
ENV APP_PATH /data
ENV GO111MODULE=on  

# Install git
RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    apk add build-base

# Create src dir
RUN mkdir $GOPATH/${APP_PATH}

# Set workdir to back
WORKDIR $GOPATH/${APP_PATH}

# Run the server
CMD ["sh", "start.sh"]
