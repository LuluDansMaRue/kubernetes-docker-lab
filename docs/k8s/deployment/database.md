# Database

Previously we saw one type of deployment. This deployment was what we call ```stateless``` this mean that when pods restarting we loose the ```datas``` that are written in the pods.

Kubernetes provide an other kind of deployment. This one is name ```StatefulSet``` as you can imagine this kind of deployment allow us to keep the state of our pods continually

# Creating the Docker image

```shell
docker build -t sesame_db -f build/db/Dockerfile <path to root folder>/kubernetes-docker-lab
```

