
# Example

1. Create docker image

```shell
docker build -t sesame_front -f build/node/Dockerfile.release /Users/marcintha/workspace/kubbie
```

2. Create the configuration

3. Launch the deployment

```shell
kubectl create -f k8s/deployment/<deployment_yaml_file.yml>
```

4. Watch the deployment

```shell
kubectl rollout status deployment.v1.apps/<name_of_deployment>
```

5. Check if the pods are running properly

```shell
kubectl get pods
```

See status. If different from 'running'. Then check the pod event or log

Pod event

```shell
kubectl logs -p <pod_name>
```

Event

```shell
kubectl describe pod <pod_name>
```

6. Expose the deployment to the localhost (creating loadbalancer on minikube)

- First check if the deployment is already exposed

```shell
kubectl get svc
```

- If the deployment is not present run this command

```shell
kubectl expose deployment frontend-vue --type=NodePort
```

- Then you can generate an URL in order to access to your deployment

```shell
minikube service <deployment> --url
```