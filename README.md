# Rollout Proxy

**Rollout Restart Kubernetes Deployments from inside a k8s cluster**

# Usage

This service is for restarting k8s Deployments from inside the cluster either for CICD purposes or your Cluster is Private

# Environment Variables

`USERNAME`: string, required: Username for Basic Auth

`PASSWORD`: string, required: Password for Basic Auth

`INCLUSTER`: TRUE or FALSE, defaults to FALSE, optional: Whether the service is running inside or outside a cluster.
