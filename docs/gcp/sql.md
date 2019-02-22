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

CloudSQL Proxy come with a seperate image that's going to be installed as a sidecar in your pod. In order to understand what a Sidecar is. Take a look to the related article [What's a sidecar](../architecture.md#Pod).

### Creating a new user

Firstly, we're going to create a new user that's going to be used by the CloudSQL Proxy. In order to create a user click on your ```database``` then click on the ```users``` tab. Create a new user with the following credentials

```shell
username: alison
password: bobba
``` 

Now let's make a connection test in order to check that our user has been added.

```shell
gcloud sql connect db --user=alison
```

Now let's activate the **cloud sql admin api** by going into the **API** section of the GCP dashboard.
- On the searchbar type: ```cloud sql```
- Click on the **Cloud SQL Admin SQL**
- Click on the **Enable** button
- Create a service account by following this article [service account](https://cloud.google.com/sql/docs/mysql/sql-proxy#create-service-account)
- Get the instance connection name by running the command ```gcloud sql instances describe db```
- The name should be equal to the ```connectionName``` property within your yaml

### Create the credentials

We're going to need 2 types of secrets:

- cloudsql-instance-credentials: Used by application to talk to the CloudSQL instance. It uses the credentials files generated when you create the service account.

- cloudsql-db-credentials: The new username & password added to the database earlier.

Let's create the secret for the ```cloudsql-instance-credentials``` by running the command

```shell
kubectl create secret generic cloudsql-instance-credentials --from-file=credentials.json=<path_to_generate_credentials_file.json>

# If everything goes as expected the shell should print
# secret "cloudsql-instance-credentials" created
```

Let's create the secret for the ```cloudsql-db-credentials``` by running this command

```shell
kubectl create secret generic cloudsql-db-credentials --from-literal=username=alison --from-literal=password=bobba

# If everything goes as expected the shell should print
# secret "cloudsql-db-credentials" created
```

### Update the configuration file

Now that we have create our credentials we can update our deployment. By adding a new container into our pod. This container is based on the cloudsql-docker image.

```yaml
spec:
    containers:
    #... the bobba-api container
    # Add the cloudsql-proxy below
        - name: cloudsql-proxy
        image: gcr.io/cloudsql-docker/gce-proxy:1.11
        # Command that's going to be executed when the pod will start
        command: ["/cloud_sql_proxy",
                    "-instances=kubernetes-demo-232217:us-central1:db=tcp:3306",
                    "-credential_file=/secrets/cloudsql/credentials.json"]
        securityContext:
            runAsUser: 2  # non-root user
            allowPrivilegeEscalation: false
        # Reference the volume inside the container
        volumeMounts:
            - name: cloudsql-instance-credentials
            mountPath: /secrets/cloudsql
            readOnly: true
    # Link the volumes where our credentials are stored
    volumes:
    - name: cloudsql-instance-credentials
        secret:
            secretName: cloudsql-instance-credentials
```

Futhermore let's update our bobba-api container in order to pass the credentials that we generated from the ```cloudsql-db-credentials```

```yaml
containers:
- name: bobba-api
  image: gcr.io/kubernetes-demo-232217/sesame_api:v1
  imagePullPolicy: IfNotPresent
  # Override the environment file that we'd define on the Dockerfile
  env:
    - name: MYSQL_HOST
        value: 127.0.0.1
    - name: MYSQL_USER
        valueFrom:
        secretKeyRef:
            name: cloudsql-db-credentials
            key: username
    - name: MYSQL_PASSWORD
        valueFrom:
        secretKeyRef:
            name: cloudsql-db-credentials
            key: password    
  ports:
    - containerPort: 8000
  args:
    - sh
    - start.sh
```
