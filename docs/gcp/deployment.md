## Before deploying

Previously we'd create the images for kubernetes which are now stored within GCP. Now let's deploy our application on GCP.

Remember how we deploy our application with minikube ? Well for GCP this is going to be the same ! but with one difference regarding our configuration file for our deployment. Let's take a look at what we change.

- First we have concat our service & deployment in the same file with the seperator ```---```
- We have change our ```ImagePullPolicy``` with the one below 

```yaml
# If the image is already present Kubernetes won't pull it 
imagePullPolicy: IfNotPresent
```

- Thirdly we had change the image so that it point to the one stored in GCP

```yaml
image: gcr.io/<project_name>/sesame_api:v1
```

## Let's deploy

Let's deploy our application by running this command below

```shell
kubectl create -f gcp/deployment/front.yml
kubectl create -f gcp/deployment/api.yml
```

If everything is running correctly. Run the command

```shell
kubectl get pods
```

And you should see the list of pods related to your deployments.
**Wait a few minutes** ...

And run the command

```shell
# svc is an alias for services
kubectl get svc 
```

You should get the list of available services. And most notably an ExternalIP for our LoadBalancer that use the ```front``` project.

**How about our API** ?

With Minikube we only had 1 Cluster and most importantly **1 Node**. Which is not the cased in our deployment. Indeed if you run the command

```shell
kubectl get nodes --output wide
```

You should see a list of 3 nodes with an ExternalIP which you can reach the API throughout the NodePort (31320)

#### Now let's take a [look at Ingress](ingress.md)