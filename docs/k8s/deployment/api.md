# API deployment Description

Deploying the API to Kubernetes required you several things

- kubectl
- minikube
- a docker image
- a yaml deployment configuration file

# Building our Docker image

For the sake of the demo a Dockerfile.release is available in the ```build/api/``` folder. This Dockerfile doesn't differ that much from the original dockerfile. It just add the project to the Dockerfile.

First of all from the root of the folder run this command

```shell
docker build -t sesame_api -f build/api/Dockerfile.release <path to root folder>/kubernetes-docker-lab
```

For an explanation of the command please check the [front deployment](front.md) article

Once the image is build check if the docker image is present by using the ```docker images``` command.

# Create the deployment configuration

Like the front-end deployment we will also use the ```Deployment``` type for deploying our API

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

# Deploying our application to Minikube

Now that we have our yaml configuration file and our docker image. Let's deploy our app. Run the following command

```shell
kubectl create -f k8s/deployment/api_deployment.yml
```

Secondly listen the deployment status of the pod

```shell
kubectl rollout status deployment.v1.apps/bobba-api
```

You should get a success message that said that your deployment is successfull. 

Finally check if your pod are running by using the command

```shell
kubectl get pods
```

An other step we can do is to check one of the pod itself by using this command

```shell
kubectl exec -ti <pod_name> -- /bin/sh
```

This will enable us to inspect the pod itself. Now let's check if the ```CompileDaemon``` process is running

```shell
ps aux
```

If you see the ```CompileDaemon``` process this mean that our application is running !

Et voilà you made your second deployment is done !