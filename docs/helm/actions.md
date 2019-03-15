---
layout: single
classes: wide
title: Simple action with Helm
sidebar:
  nav: "docs"
---

### Upgrade

Now that we'd deploy our project with Helm. We'll do some operation with our ```bobba-api``` project. As you can see we'd define a set of replicasCount to ```5```. Let's update this field and set 3 replicas instead of 5.

Edit the ```replicasCount``` field in the ```values-api.yaml```
Let's now update our helm chart by running this command

```shell
helm upgrade -f bobba-helm-chart/values-api.yaml <release_name> bobba-helm-chart
```
If everything is successfull you should get the list of objects that has been re-deployed to Kubernetes as well as a message which said ```Release "<release_name>" has been upgraded. Happy Helming!```

Now if everything went as planned we should get **3 pods of the bobba-api**. Let's check it out by running the command

```shell
kubectl get pods | grep bobba-api
```

You should only see **3 pods !**

### Rollback

Helm is keeping a history of your release allowing you to rollback to a release version if needed. Let's rollback to our previous version by firstly getting the history of our release

```shell
helm history <release_name>
```

Now let's rollback to our first release

```shell
helm rollback <release_name> 1
```

Now if you're checking the pods you should see that we have our **5 pods** back ! awesome !

If you're doing the command

```shell
helm history <release_name>
```

You should see that rolling back will create an other release with the description saying that we'd rollback to the 1st version.

## Wrap up

In this small example we saw how to deploy our small app to Minikube with the help of Helm. We saw that using Helm allow us to be more flexible as well as doing rollback and upgrade easily with only 1 command.

Helm is a great tool that can be used for deploying our app easily with less headache than configuring every yaml file manually. In the next section we're going to see if our current Helm configuration is applicable to GCP