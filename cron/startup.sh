
# Install dependencies for the bruteforce.go
cd bruteforce && go get ./... && cd ..

# Install dependencies for select.go
cd select && go get ./... && cd ..

# Install dependencies for temporary.go
cd temporary && go get ./... && cd ..

# Run the crontab only in the context of Docker
if [ "$1" = "docker" ] 
then
  crond -f -l 8
fi

# Run the go file in the kubernetes context
if [ "$1" = "kube" ]
then
  for i in $(ls -d */ | cut -f1 -d'/');
  do
    if [ "$2" = ${i} ]
    then
      echo "runing task ${i}"
      go run "$i/$i.go"
    fi
  done
fi