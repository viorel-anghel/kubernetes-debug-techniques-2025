
```

# before install
helm show values oci://registry-1.docker.io/bitnamicharts/postgresql
helm show readme oci://registry-1.docker.io/bitnamicharts/postgresql

# install helm chart
helm upgrade --install --create-namespace -n helm-pg -f values.yaml mypg oci://registry-1.docker.io/bitnamicharts/postgresql


# after install
helm -n helm-pg ls
helm -n helm-pg get values  mypg

# using helm template instead of install
helm template --create-namespace -n helm-pg -f values.yaml mypg oci://registry-1.docker.io/bitnamicharts/postgresql

```

