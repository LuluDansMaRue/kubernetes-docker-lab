## CronJob ⏱️

In this section we will take a look at how to implement a CronJob in Kubernetes by using Minikube.

## Configuration

In order to do this deployment you'll need to have minikube installed as well as having Docker.

## Images 💿

For these deployment examples we'll use an image which is located in the **build/cron** folder. I suggested you to take a look at it. It's based on the golang image.

## Steps

### Creating the docker image

First, we will need to start minikube.

Start minikube with the following command:

```shell
minikube start
```

Now if you do 
```shell
docker ps
```

You should see a list of **docker images**. Great !
If that's the case we can then proceed to build a docker image. We will use this image as the main image for our CronJob. Run this command within the same terminal or window terminal. We don't want to loose the docker's environment variable set to minikube :)

```shell
docker image build -t sesame_minikube -f build/cron/Dockerfile <Absolute path to the project>/kubernetes-docker-lab
```

#### Image explanation

The docker image is based on the golang alpine docker. It **mount** the folder **cron** into the **task** folder within the pod.

The *-f* option is used to in order to provide the context of execution to the docker. We want to run the build from the root folder of the project in order to mount the **cron** folder properly

Now that we'd build our docker image we will configure our CronJob file.

## Configuration of a CronJob

Kubernetes make it easy to configure a cronjob. Below is a sample of what look like a CronJob

```yml
apiVersion: batch/v1beta1
kind: CronJob <Kind of controller>
metadata:
  name: <name of the service>
spec:
  schedule: "<timeout>"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: <name of the container>
            image: <name of the image>
            imagePullPolicy: Never --> We use this value as we want to use a local image
            args:
            - a list of arguments that could be use during the startup of the pod
          restartPolicy: <See the restart policy section>

```

Pretty sample right ? Below is the description of each field

| Name                       | Description                                                                                                |
|----------------------------|------------------------------------------------------------------------------------------------------------|
| Kind                       | Represent the type of controller to uses                                                                   |
| name                       | The name of the CronJob                                                                                    |
| schedulde                  | The timeout before the cronjob trigger itself. This respect the default definition of triggering a cronjob |
| containers.name            | The name of the containers                                                                                 |
| containers.image           | The name of the image use by the container. In our case it represent a local docker images                 |
| containers.imagePullPolicy | Strategy to use when dealing with update, creating an image                                                |
| containers.args            | A list of arguments that can be triggered during the executon of the pod                                   |
| restartPolicy              | The kind of restarting strategy when something wrong happened to the pod                                   |

Below is an example of how the configuration of a cronjob looks.

```yml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: bruteforce
spec:
  schedule: "*/15 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: bruteforce
            image: sesame_minikube
            imagePullPolicy: Never
            args:
            - sh
            - install.sh
            - go
            - run
            - bruteforce/bruteforce.go
          restartPolicy: OnFailure

```

## Running the CRON

First we're going to create the service within our cluster by running this command

```shell
kubectl create -f k8s/job/cronjob-temporary.yaml
```

Now let's check if the cronjob services has been created by running this command

```shell
kubectl get cronjob <cronjob_name>
```

Once you see the cronjob within the list. Your cronjob task should be able to be trigger based on the schedulde you defined. Or you could either check if a pod has been created by the CronJob with the following command

```shell
kubectl get jobs --watch
```

Once you see a new job appaear on the list this mean that a **new pod** has been created. You could also check the status of the pod by running this command

```shell
kubectl get pods
```

## Debugging the CronJob

Sometimes it can happened that you missly configured the cronjob or that something went wrong.

I suggested you to firstly get the **events** of the pod. It will tell you if there are anything wrong regarding the configuration of the CronJob

```shell
kubectl describe pod <name of failed pod>
```

Secondly if you didn't see anything wrong with the pod I suggested you to take a look at the logs of the pod.

Below is the command to check the log for an active pod

```shell
kubectl logs -p <pods_name>
```

If the pod is in terminated status or not running you could use this command

```shell
kubectl logs <pod_name>
```

Now you know how to deploy a CronJob with minikube !
