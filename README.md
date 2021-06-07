# ![Beeyond](frontend-beeyond/src/assets/images/beeyond-logo-with-text.png)

**Kubernetes provisioning tool** for the HTBLA Leonding (Dept. of Informatics and Media Technology)

<a href="https://github.com/halilbahar/beeyond/graphs/contributors" alt="Contributors">
  <img src="https://img.shields.io/github/last-commit/halilbahar/beeyond/master"/>
</a>
<a href="https://github.com/halilbahar?tab=packages&repo_name=beeyond" alt="Version">
  <img src="https://img.shields.io/github/v/tag/halilbahar/beeyond"/>
</a>
<a href="https://github.com/halilbahar/beeyond/pulls" alt="PullRequests">
  <img src="https://img.shields.io/github/issues-pr/halilbahar/beeyond"/>
</a>
<a href="https://halilbahar.github.io/beeyond/" alt="Documentation">
  <img src="https://img.shields.io/static/v1?label=docs&message=here&color=orange"/>
</a>
<a href="https://halilbahar.github.io/beeyond/reports.html" alt="Reports">
  <img src="https://img.shields.io/static/v1?label=reports&message=here&color=orange"/>
</a>

## Links
- [GitHub Pages: AHIF](https://2021-4ahif-syp.github.io/assigment02-system-specification-kubernetes-provisioning-tool/)
- [GitHub Pages: AHITM](https://halilbahar.github.io/beeyond)

## Sample deployment for nginx (yaml)

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
