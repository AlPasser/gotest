docker build -t goweb .
docker tag goweb alpasser/goweb:20221218
docker push alpasser/goweb:20221218
docker run --rm --name goweb -p 8080:80 -d goweb
docker logs -f goweb
curl -i localhost:8080/healthz
#export PID=`docker inspect goweb | grep Pid | head -1 | awk '{print $2}' | awk -F, '{print $1}'`
export PID=`docker inspect -f {{.State.Pid}} goweb`
sudo nsenter -n -t$PID ifconfig
