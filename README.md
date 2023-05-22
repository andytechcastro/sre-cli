# SRE-CLI

The idea of this project is build a tool that connect to the main cli tools for SRE

## Curret tools
- Kubernetes

## How to use
### Kubernetes 
#### Deployments
For get all deployments for all your clusters
```sre-cli kube deployments```

#### Pods
For get all pods for all your clusters
```sre-cli kube pods```

In both cases you can use -A for all namespaces or -n {namespace} for select a namespace
