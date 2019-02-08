## Cron jobs (WIP)

Running a Cronjob in minikube required you to define a .yml file.


### Configuring a cron job

// TODO do the documentation of the yaml file...

### Steps

/!\ Must be done on the same terminal window

1. Build the local docker images for minkube from the build/cron folder

```shell
docker image build -t sesame_minikube -f build/cron/Dockerfile /Users/marcintha/workspace/kubbie
```

2. Creating the job file

3. Creating the cron job service

```shell
kubectl create -f k8s/job/cronjob-temporary.yaml
```

4. Run the command 

```shell
kubectl get cronjob <cronjob_name>
```

This display the cronjob. It's indeed not active yet as it will wait for the schedule to be elapsed in order to activate the cron

5. Debugging

- Get events & config files

```shell
kubectl describe pod temporary-1549618800-djvh9
```

- Watch for the creation of the job

```shell
kubectl get jobs --watch
```

- Get log of a pod

```shell
kubectl logs -p
```

6. Delete cron

```shell
kubectl delete cronjob <name>
```

7. Error type:

- CrashLoopBackOff: Crash from the pod good it's an application crash and not a pod config stuff damn it

### Command to run


### Resources

## Help

If you get an ```ssh error...: operation timed out ```
Locate the ```hyperkit.pid``` and delete it.

Below is an example on OSX:

```shell
rm -rf ~/.minikube/machines/minikube/hyperkit.pid
```