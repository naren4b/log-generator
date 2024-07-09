# log-generator

```
IMAGE_NAME=$(uuidgen)
IMAGE_PATH="ttl.sh/${IMAGE_NAME}:1h"
docker build -t $IMAGE_PATH .
docker push $IMAGE_PATH
docker run -d --rm --name=log-generator -e LOG_INTERVAL=2 $IMAGE_PATH

```
pod yaml
```
cat<<EOF >pod.yaml 
apiVersion: v1
kind: Pod
metadata:
  name: log-generator
spec:
  containers:
    - name: log-generator
      image: $IMAGE_PATH
      env:
        - name: LOG_DURATION
          value: "300000"  # 5 minutes in milliseconds
EOF
```

k apply -f pod.yaml 