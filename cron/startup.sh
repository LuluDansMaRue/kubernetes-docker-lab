# Install dependencies for the bruteforce.go
cd bruteforce && go get ./... && cd ..

# Install dependencies for select.go
cd select && go get ./... && cd ..

# Install dependencies for temporary.go
cd temporary && go get ./... && cd ..

# Run the crontab
crond -f -l 8