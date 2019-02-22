# Installing dependencies
npm install

# Create environment file
touch .env

# Set env variable
echo "VUE_APP_DEPLOYMENT=$DEPLOYMENT" >> .env

# Run the serve
npm run serve