FROM golang:1.11-alpine3.8

ENV APP_PATH /task

# Install git
RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    apk add build-base

# Create src dir
RUN mkdir ${APP_PATH}

# Set workdir to back
WORKDIR ${APP_PATH}

# Copy the folder to the task folder
COPY cron /${APP_PATH}

# Install the dependencies
RUN go build ./...
