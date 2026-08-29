package manifest

// DeploymentYAML is a minimal Kubernetes Deployment manifest for the capstone demo.
const DeploymentYAML = `apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-capstone
  labels:
    app: go-capstone
spec:
  replicas: 2
  selector:
    matchLabels:
      app: go-capstone
  template:
    metadata:
      labels:
        app: go-capstone
    spec:
      containers:
        - name: api
          image: ghcr.io/example/go-capstone:1.0.0
          ports:
            - containerPort: 8080
          readinessProbe:
            httpGet:
              path: /readyz
              port: 8080
            initialDelaySeconds: 3
            periodSeconds: 5
          livenessProbe:
            httpGet:
              path: /healthz
              port: 8080
            initialDelaySeconds: 10
            periodSeconds: 10
          resources:
            requests:
              cpu: 100m
              memory: 128Mi
            limits:
              cpu: 500m
              memory: 256Mi
`

func Summary() string {
	return "Deployment go-capstone: 2 replicas, probes on /readyz and /healthz"
}
