# ![Beeyond](frontend-beeyond/src/assets/images/beeyond-logo-with-text.png)

<a href="https://github.com/badges/shields/graphs/contributors" alt="Contributors">
  <img src="https://img.shields.io/github/last-commit/halilbahar/beeyond/master"/>
</a>

**Kubernetes provisioning tool** for the HTBLA Leonding (Dept. of Informatics and Media Technology)

## Sample ingress for nginx (yaml)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 2 # tells deployment to run 2 pods matching the template
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
```
