# Installing dependencies
npm install

# Create environment file
touch .env

# Set env variable
echo "VUE_APP_DEPLOYMENT=$DEPLOYMENT" >> .env

# Run the serve
if [$DEPLOYMENT = "minikube"] || [$DEPLOYMENT = "gcloud"]; then
  npm run build
  serve -l 8080 -s dist
else
  npm run serve
fi
