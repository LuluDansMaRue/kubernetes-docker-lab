FROM golang:1.11-alpine3.8

ENV APP_PATH /task

# Env for MySQL
ENV MYSQL_USER=alison
ENV MYSQL_PASSWORD=bobba
ENV MYSQL_ROOT_PASSWORD=bobbaroot
ENV MYSQL_HOST=db
ENV MYSQL_PORT=3306
ENV MYSQL_DATABASE=thebestbobba

# Env for Golang module
ENV GO111MODULE=on 

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
