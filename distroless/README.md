
```
# build distoless image using Dockerfile
docker buildx build --platform linux/amd64 -t go-http-server:0.1 .

# create a deployment with image
kubectl create deploy go-http-server --image=go-http-server:0.1

# get pods and exec
kubectl get pods
kubectl exec -ti go-http-server-7bb984744d-d7dhx -- sh

# use debug instead of exec
kubectl debug -it -c debug --image=busybox go-http-server-7bb984744d-d7dhx

# apply shareProcessNamespace: true
kubectl patch deploy go-http-server -p '{"spec": {"template": {"spec": {"shareProcessNamespace": true }}}}'

# run the debug again
kubectl debug -it -c debug --image=busybox go-http-server-657964647c-lfxxc 
  ps -ef
  ls -la /proc/7/root/
  exit

# use copy-to to create another pod
kubectl debug -it -c debug --image=busybox --share-processes --copy-to debug-pod go-http-server-7bb984744d-d7dhx  


# debug NODE
kubectl debug node/<NODE-NAME>  -it --image=ubuntu
  
```
