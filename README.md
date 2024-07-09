# log-generator

```
IMAGE_NAME=$(uuidgen)
$ docker build -t ttl.sh/${IMAGE_NAME}:1h .
$ docker push ttl.sh/${IMAGE_NAME}:1h
docker run -d --rm -it -e LOG_INTERVAL=2 ttl.sh/${IMAGE_NAME}:1h

```
