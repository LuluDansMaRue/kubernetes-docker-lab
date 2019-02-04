FROM golang:1.11-alpine3.8

ENV APP_PATH /data

# Install git
RUN apk update && apk upgrade && \
    apk add --no-cache git

# Create src dir
RUN mkdir ${APP_PATH}

# Set workdir to back
WORKDIR ${APP_PATH}

# Run the server
CMD ["sh", "start.sh"]
