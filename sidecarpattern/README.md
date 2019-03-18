refer to the url: [hand crafting a sidecar proxy like istio](https://venilnoronha.io/hand-crafting-a-sidecar-proxy-like-istio)

1. build image for sidecar proxy mage

```
docker build -t jeffcai/sidecar-proxy-image:latest -f Dockerfile .
docker push jeffcai/sidecar-proxy-image:latest
```

2. build init container image

```
docker build -t jeffcai/init-networking:latest -f Dockerfile-init .
docker push jeffcai/init-networking:latest
```

3. test

- to get the pod cluster ip

```
kubectl get pods -o wide
```

- send a get request

```
kubectl run -i --rm --restart=Never busybox --image=odise/busybox-curl \
          -- sh -c "curl -i 172.17.0.2:80/get?query=param"
```

- send a post request

```
kubectl run -i --rm --restart=Never busybox --image=odise/busybox-curl \
          -- sh -c "curl -i -X POST -d 'body=parameters' 172.17.0.2:80/post"
```

- send a get request

```
kubectl run -i --rm --restart=Never busybox --image=odise/busybox-curl \
          -- sh -c "curl -i http://172.17.0.2:80/status/429"
```

- logs to the proxy container

```
kubectl logs httpbin-pod --container="proxy"
```
