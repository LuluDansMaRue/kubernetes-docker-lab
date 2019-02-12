# Example of creating a development environment with Docker

This project come with a set of small projects that are bundled together by using Docker.

## Dockerfile

Using a raw image can be useful for quick prototyping. However when you need to make internal work (e.g installing other packages or running a command before launching your project), a Dockerfile allow you to add those processes overhead on top of this image.

### Example with the API

This project's API is made with Go, a programming language made by Google. For this project, we will use the ```alpine``` image. Alpine is a lightweight linux environment. In order to run our API, we need some third party plugins.

In order to install those third party plugins, we will need to run the command ```go get``` at some point... However ```go get``` need ```git``` to be installed **which is not the case** on `alpine`.

```
This can be a good way to make a custom image so we can add git !
```

So let's create a Dockerfile

#### Dockerfile

A Dockerfile always start from a base image. Therefore, we will use the keyword *FROM*

```dockerfile
FROM image_name:version
```

Dockerfile has a set of instructions available such as ADD, RUN, CMD... I suggest you take a look at the list [here](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#dockerfile-instructions)

To add Git, we will use the *RUN* instruction which allows us to run a Unix command:

```dockerfile
FROM golang:1.11-alpine3.8-need-lao-sausage

RUN apk update && apk upgrade && \
    apk add --no-cache git
```

We could also copy the project's directory, but we won't do that. We will use docker-compose.

Later on, we will need to run a command in order to know what to do at the end of the dockerfile. Usually this process is attached, which is **running in foreground**. Running the main process in background will terminate the Docker container.

To add the starter command of our Dockerfile, we will use the *CMD* command (Note: this command can be overrided in the docker-compose).

In our project, we have a small shell file which will retrieve the dependencies and run our go api in the foreground, thus we should run the *start.sh*:

```dockerfile
# Example of usage
# CMD ["args1", "args2"] 

CMD ["sh", "start.sh"]
```

Et voilà, we've created our Dockerfile. Let's move on to docker-compose.

#### Docker-compose

Our API is not running alone. Indeed, the API is used by other projects and uses other projects. Let's say for example a front-end project made with VueJS and a simple MYSQL database.

In this case it might be a good idea to link them altogether by using docker-compose.

So what is Docker-compose ? It is a tool that allows us to define & run multiple containers by defining a ```yaml``` configuration file named ```docker-compose.yml```. Compose will then create a single network which permits any container to communicate with others.

Let's see how we can define a docker-compose file

First, it starts with the version of Compose.

```yml
version: '3'
```

Then, we will define the list of containers that we want:

```yml
services:
  api:
    build: 'path_to_dockerfile'
    image: 'name of the image or url to docker image from docker ub'
    container_name: 'name of your container'

    # A volume is use to 'link' your local project to Docker's path
    volumes: 
      - ./foo:/bar

    # Ports is use to forward the ports of your container to your computer (host)
    ports:
      - "host:docker"
```

With this basic configuration, you can have multiple services. Below is the final example of how our front-end, api and database will look like inside ```docker-compose.yml```:

```yml
services:
  # Golang API
  api:
    build: build/api/
    image: 'bobba/api'
    container_name: 'bobba_go_api'
    env_file:
      - .env
    volumes:
      - ./api:/data
    ports:
      - "9000:8000"
    depends_on:
      - db

  # Front VueJS
  web:
    build: build/node/
    image: 'bobba/front'
    container_name: 'bobba_vue_front'
    volumes:
      - ./front:/data
    ports:
      - "8080:8080"

  # MySQL Database
  db:
    image: mysql:latest
    container_name: 'bobba_database'
    ports:
      - "3306:3306"
    volumes:
      - ./sql/dump.sql:/docker-entrypoint-initdb.d/dump.sql
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD}
      MYSQL_DATABASE: ${MYSQL_DATABASE}
      MYSQL_USER: ${MYSQL_USER}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD}
```
