# Example

In order to understand how to deploy a full web app in Kubernetes I provide a small example which is make up of :

- front: a small front-end made with VueJS
- back: a small back-end which contain Rest API made with Go
- database: a simple MySQL database with only one table

# How are we going to deploy this web app

Based on what we learn so far. We know that a cluster is made of several Nodes where in each of these Nodes there are multiples pods which are managed by the Kubelet.

Moreover we also saw different kind of deployment that allow us to manage the behavior of our pods by using different type of controller.

Let's define the constraint of our web app:

- stateless, everything is stored in the database
- we need to have the same amount of pods available
- for our convenience we need to be able to update, rollback our deployment if anything happened

Based on the types of deployment we know the one that fits the best is the ```deployment``` types.

***What about our database ?***

Our database is a statefull kind of service. With Kubernetes we have the option to create a stateful service with the ```statefulSet``` controller. However the community is divide at the idea of running a database in Kubernetes.

A deep explanation has been made on the subreddit of Kubernetes. **Credits to nkrgovic for it's detailed explanation**

> Stateful services, especially databases, rely on keeping state on a reliable media - in most cases disk (ssd or spinning rust). Also, in most cases, database performance is a big issue affecting system performance. This is partly due to design, and partly due to inherent inability of relational databases to scale horizontally.

> Simple put, the most common bottleneck in databases is the disk performance. In order to get good database performance you need disk performance. This means predictable latency and good throughput - which usually means dedicated volumes with QoS, not sharing out of a pool.

> Other issues have to do with databases using a lot of memory, and also not being good to spin up and spin down new instances quickly, exactly because of the state: They need to copy the entire data set over and that takes time.

> When you put all those together a relational database is a very bad candidate for running in a container - it just doesn't make sense. If you have a lot of data spinning up and down, starting, everything takes time, and a lot of time.

> If you have a little bit of data, like in a test or CI environment, then database in a container is a great tool. But in a production environment databases and containers just don't mix.


Moreover Kelsey Hightower a developer advocate on the GCP advise to also not use Kubernetes for deploying a database. See his [tweet](https://twitter.com/kelseyhightower/status/963413508300812295?lang=en)

Therefore we're going to seperate our database from the Kubernetes and for the sake of this demo run the database on our own host. However for the sake of the demo provided in the k8s/database folder **(WIP)**.

# Linking

As our front is tightly coupled to our API we will need to connect these 2 deployment altogether. In order to do that we'll use a service. **(WIP)**

# Demonstration

Please refer to these articles for each deployment environment

[Deploying the front](./front.md)
[Deploying the back](./api.md)