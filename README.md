
# Simple-Mock-Server

## Use json to define services

```json
{
    "services": [
        {
            "location": "www.example.com",
            "response": {
                "status": 200,
                "header": {

                },
                "body": ""     
            }
        }   
    ]
}


```

# Deploy to k8s

```yaml

---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: mock-server
  labels:
    app: mock-server
    component: cicd
spec:
  replicas: 1
  selector:
    matchLabels:
      app: mock-server
      component: cicd
  template:
    metadata:
      labels:
        app: mock-server
        component: cicd
    spec:
      containers:
      - name: app
        image: hub.c.163.com/commonwork/mock-server:latest
        imagePullPolicy: Always
        ports:
        - containerPort: 8080

---
apiVersion: v1
kind: Service
metadata:
  name: mock-server-svc
spec:
  selector:
    app: mock-server
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080

```