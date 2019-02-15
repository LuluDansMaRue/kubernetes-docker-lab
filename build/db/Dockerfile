FROM mysql:latest

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=bobbaroot
ENV MYSQL_DATABASE=thebestbobba
ENV MYSQL_USER=alison
ENV MYSQL_PASSWORD=bobba

# Add SQL dump
ADD sql/dump.sql /docker-entrypoint-initdb.d/