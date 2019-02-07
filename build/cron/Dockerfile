FROM golang:1.11-alpine3.8

ENV APP_PATH /task

# Install git
RUN apk update && apk upgrade && \
    apk add --no-cache git 

# Create src dir
RUN mkdir ${APP_PATH}

# Set workdir to back
WORKDIR ${APP_PATH}

# Copy the crontab file
COPY bobba-cron.txt bobba-cron.txt

# Give right to the crontab so that it can be exec
RUN chmod 755 bobba-cron.txt

# Create log for crontab
RUN touch /var/log/bobba-cron.log

# Register the crontab to the cron file
RUN /usr/bin/crontab bobba-cron.txt

# Run the startup
CMD ["sh", "startup.sh"]