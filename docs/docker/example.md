# Example of creating a development environment with Docker

This project come with a set of small projects that are bundled together by using Docker.

## Dockerfile

Using an image in a raw way can be useful for quick prototyping. However sometimes when you need to do more things internally e.g installing other packages or running command that need to be done before launching your project. a Dockerfile allow you to add more process overhead on top of a based image.

### Example with the API

In this project the API is made with Go a programming language made by Google. For this project we use the ```alpine``` image. Alpine is a lightweight linux environment. In order to run our API. Our API need some third party plugins.

In order to get these third party plugins we will need to run the command ```go get``` at some point... However ```go get``` need ```git``` to be installed **which is not the cased** on this image.

```
This can be a good way to make a custom image so we can add git !
```

So let's create a Dockerfile

#### Dockerfile

A Dockerfile always start from a base image therefore we will use the keyword *FROM*

```dockerfile
FROM image_name:version
```

Dockerfile has a set of instruction available such as ADD, RUN, CMD... I suggest you to take a look at the list of instruction available here [Dockerfile instruction](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#dockerfile-instructions)

For adding Git we will use the *RUN* instruction which allow us to run an unix command.

```dockerfile
FROM golang:1.11-alpine3.8-need-lao-sausage

RUN apk update && apk upgrade && \
    apk add --no-cache git
```

Now we could also copy the directory of our project but we won't do that we will use docker-compose...

Later on we will need to run a command in order for our Dockerfile to know what to do at the end of the dockerfile. Usually this process is a process attach which is **running in foreground**. Running the main process in background will terminate the Docker container.

In order to add the starter command to our Dockerfile we will use the command *CMD* (Note: this command can be override in the docker-compose)

In the case of our project we have a small shell file which will retrieve the dependencies & run our go api in foreground thus we will run the *start.sh*

```dockerfile
# Example of usage
# CMD ["args1", "args2"] 

CMD ["sh", "start.sh"]
```

Et voilà we have create our Dockerfile. Let's move on to the docker-compose

#### Docker-compose

Our API is not running alone. Indeed the API is used by other project and use other project. Let's say a front-end made with VueJS and a simple MYSQL database.

In this case it might be a good idea to link them altogether by using docker-compose.

So what's Docker-compose. Docker-compose is a tool that allow us to define & running multiples containers by defining a ```yaml``` configuration file named docker-compose.yml. Compose will then create a single network which every containers will be able to communicate with others in the same network.

Let's see how we can define a docker-compose file

First. it all start with the version of the Compose.

```yml
version: '3'
```

Secondly we will define the list of containers that we need to have

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

With this basic configuration you could configure multiple services. Below is the final example of how our front-end, api & database is:

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
