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
Apply to cluster 
```
k apply -f pod.yaml 
k get pod log-generator
k logs -f log-generator

```

Testing 
```
kubectl port-forward --namespace monitoring svc/loki 3100:3100 &


# Inject 
curl -H "Content-Type: application/json" -XPOST -s "http://127.0.0.1:3100/loki/api/v1/push"  \
--data-raw "{\"streams\": [{\"stream\": {\"job\": \"test\"}, \"values\": [[\"$(date +%s)000000000\", \"fizzbuzz\"]]}]}" \
-H X-Scope-OrgId:foo


# retrive 
curl "http://127.0.0.1:3100/loki/api/v1/query_range" --data-urlencode 'query={job="test"}' -H X-Scope-OrgId:foo | jq .data.result

# for the pod logs

curl "http://127.0.0.1:3100/loki/api/v1/query_range" --data-urlencode 'query={container="log-generator"}' -H X-Scope-OrgId:1
```
