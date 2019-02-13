# Installing dependencies
npm install

# Run the serve
if [ "$1" = "minikube"]
then
  npm run serve-kube
else
  npm run serve
fi