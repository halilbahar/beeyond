INSERT INTO template (name, description, content) VALUES ('nginx', 'A simple nginx deployment',
'apiVersion: v1
kind: Service
metadata:
  name: nginx
spec:
  selector:
    app: nginx-container
  ports:
  - port: 80
    targetPort: 80
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx-container
  template:
    metadata:
      labels:
          app: nginx-container
    spec:
      containers:
      - image: sonjacao/jee-docker_www
        name: nginx-container
        ports:
        - containerPort: 80
          name: http'
);
