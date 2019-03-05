---
layout: single
classes: wide
title: Deploying the API 
sidebar:
  nav: "docs"
---

Like the front-end project. Deploying the API required you to have several files like so:

- a docker image
- a yaml deployment configuration file

## Building our Docker image 🐣

For the sake of the demo a Dockerfile.release is available in the ```build/api/``` folder. This Dockerfile doesn't differ that much from the original dockerfile.

First of all from the root of the folder run this command

```shell
docker build -t sesame_api -f build/api/Dockerfile.release <path to root folder>/kubernetes-docker-lab
```

For an explanation of the command please check the [front deployment](front.md) article
Once the image is build check if the docker image is present by using the ```docker images``` command.

## Create the deployment configuration 🐥

Our API is also stateless. Therefore we're also going to use the ```Deployment``` type of deployment.
For a detailed explanation of the yaml file please check the [front deployment](front.md) article

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bobba-api
  labels:
    app: bobba-api
    tier: backend
spec:
  replicas: 2
  selector:
    matchLabels:
      app: bobba-api
      tier: backend
  template:
    metadata:
      labels:
        app: bobba-api
        tier: backend
    spec:
      containers:
      - name: bobba-api
        image: sesame_api:latest
        imagePullPolicy: Never
        ports:
          - containerPort: 8000
        args:
          - sh
          - start.sh
```

## Deploying our application to Minikube

Now that we have our yaml configuration file and our docker image. 
Let's deploy our app. Run the following command

```shell
kubectl create -f k8s/deployment/api_deployment.yml
```

Secondly listen to the deployment status of the pod by running this command

```shell
kubectl rollout status deployment.v1.apps/bobba-api
```

You should get a success message that said that your deployment is successfull. 

Finally check if your pod are running by using the command

```shell
kubectl get pods
```

#### Et voila the API is deployed. Now let's deploy the front [service](service_front.md)

## Errors

If you have any error you can check the event status of your pods by running this command.

```shell
# Get the list of available pods
kubectl get pods

# Get the event of a pod
kubectl describe pod <pod_name>

# Now look at the events section (should be at the end)
# You could also check the pod's log by running this command
kubectl logs -p <pod_name>

# if the pod is killed you could get the logs of a pod like this too
kubectl logs <pod_name>
```