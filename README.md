# Kubernetes debug techniques
# workshop at CloudNativeDays.ro 2025

- Original location: https://github.com/viorel-anghel/kubernetes-debug-techniques-2025
- Author: Viorel Anghel (Vang) https://www.linkedin.com/in/viorelanghel/

# "Controlled by"

kubectl create deploy nginx --image=nginx
kubectl get pods
kubectl describe pod nginx-8f458dc5b-gbzkc | grep Controlled

# kubectl get all

```
kubectl api-resources --namespaced --verbs=list  -o name | \
  xargs -n 1 kubectl get --show-kind --ignore-not-found -n <namespace>

```

