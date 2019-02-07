# Install dependencies for the bruteforce.go
cd bruteforce && go get ./... && cd ..

# Install dependencies for select.go
cd select && go get ./... && cd ..

# Install dependencies for temporary.go
cd temporary && go get ./... && cd ..

# Run the crontab
# -f (Foregroud) option is used in order for Docker to not kill the container
crond -f -l 8