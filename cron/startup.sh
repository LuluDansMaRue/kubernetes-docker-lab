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