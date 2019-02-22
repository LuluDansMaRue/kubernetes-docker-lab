## Create our docker image for the API 🏮

As we did with Minikube. We also need to create a set of Docker images in order to make them deployable to GCP with one noticeable differences. These images will be deployed to GCP and thus, store into GCP. Let's do it.

First let's deploy the API image

## Creating our API image ⚙️

The command that need to be run should looks like to what we'd saw when we deploy our API to minikube.

```shell
docker build -t sesame_api -f build/api/Dockerfile.release <path to root folder>/kubernetes-docker-lab
```

Now tag the images with a version so that we can define a special version to use for our Deployment.

```shell
docker tag sesame_api gcr.io/<projet_name>/sesame_api:v1
```

Configure Docker to use GCloud as a credential helper by running this command:

```shell
gcloud auth configure-docker
```

Finally push your image into GCR

```shell
docker push gcr.io/<project_name>/sesame_api:v1
```

## Check our image

Now let's check that are being stored into GCP.
- Go to the **Container registry** section
- Click on the **images** options, you should see the images like the image below

<p align="center">
  <img src="../img/gcp-api-image.png" alt="drawing" width="500"/>
</p>

#### Let's move on and configure our SQL database [Deploying our application !](sql.md)
