## MySQL on GCP

In today's article we're going to set up a MySQL database which is going to be use by our pods's API.

Do you remember how we did that with minikube ?

With Minikube we create a seperate StatefulSets deployment which contain a volume that was linked to our StatefulSet deployment. However as describe in previous articles, creating a database in a container is not what should be done in production environment as database need performance, storage which can grow at anytime. Thus it's recommended to use a seperate Database environment and this is what we're going to do today.

GCP has a service name Storage where we can create a relational database like MySQL or PostgreSQL

## Creating our MySQL database

Let's create our database by naming it ```db``` with the password which is ```bobbaroot```.

> Note: We use the same password that we gave with Minikube. However you should not do that in real production environment. This is intended for demonstration purposes.

Now that you have create your database we will connect to the database. Run this command in your terminal, and enter your password when ask.

```shell
gcloud sql connect db --user=root
```

Now we need to create a database. Run this command 

```shell
create database thebestbobba;
```

All right let's create our table. Run these commands

```shell
use thebestbobba;

# then
CREATE TABLE bobba (
    ID int NOT NULL,
    name varchar(255),
    price decimal,
    shop varchar(255),
    flavor varchar(255),
    calory decimal,
    PRIMARY KEY (ID)
);
```

Now you have a database with a table great isn't it ?

## Connecting our database with Cloud SQL Proxy

There are 2 ways of connecting your database with GCP. The first one is by using a private IP which is the most straightforward way.
However Google is proposing an other way of connecting your database. This one is named CloudSQL Proxy and it's made to be use with Kubernetes in mind. Therefore we're going to take a look at how to implement the CloudSQL Proxy.

CloudSQL Proxy come with a seperate image that's going to be installed as a sidecar in your pod. In order to understand what a Sidecar is. Take a look to the related article [What's a sidecar](../k8s/sidecar.md)
