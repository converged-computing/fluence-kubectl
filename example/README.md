# Deployment Example

This is an example that shows how to make a local kind cluster with ingress, and install the plugin with it.

```bash
kind create cluster --config ./example/kind-config.yaml
```

Install fluence

```bash
git clone https://github.com/flux-framework/flux-k8s
cd flux-k8s
make prepare
cd ./upstream/manifests/install/charts
helm install  --set scheduler.enableExternalService=true schedscheduler-plugins as-a-second-scheduler/
```

This is for a custom build.

```bash
# This is a registry you can push to
make REGISTRY=ghcr.io/vsoch
helm install --set scheduler.image=ghcr.io/vsoch/fluence:latest --set scheduler.sidecarimage=ghcr.io/vsoch/fluence-sidecar:latest schedscheduler-plugins as-a-second-scheduler/
```

That will install the service and enable it for fluence.