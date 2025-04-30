
```
# install CNPG postgres operator
kubectl apply --server-side -f \
  https://raw.githubusercontent.com/cloudnative-pg/cloudnative-pg/release-1.25/releases/cnpg-1.25.1.yaml

# create a simple postgres cluster
kubectl apply -f simple-pg.yaml 

# use Controlled By on pods
kubectl get pods
kubectl describe pod simple-pg-1 | grep Controlled 
kubectl get cluster
kubectl describe cluster | grep API

```


