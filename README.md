# kubectl fluence

This repository implements a single kubectl plugin for interacting with the fluence scheduler.
This is currently intended to help with debugging and development.

## Usage

### Deploy Fluence

You will need to deploy Fluence with the following flag for `scheduler.enableExternalService=true`. See
the [kind-config.yaml](example/kind-config.yaml) for how to deploy a kind cluster with ingress.

```bash
helm install \
  --set scheduler.image=ghcr.io/vsoch/fluence:latest \
  --set scheduler.enableExternalService=true \
  --set scheduler.sidecarimage=ghcr.io/vsoch/fluence-sidecar:latest \
        schedscheduler-plugins as-a-second-scheduler/
```

Note that if you are using a cloud, you might need to expose the port for ingress.

```bash
gcloud compute firewall-rules create test-cluster --allow tcp:4242
```

Note that I haven't gotten it to work on Google Cloud - I am not sure why.

### Cluster Resources

A simple `kubectl fluence` will show cluster resources.

```console
 Resource Graph Summary      
        COUNT  CORES  MEMORY 
            1     10   30720 
 TOTAL      1     10   30720 
```

### Development

Note that for this setup if you are developing locally with kind, you will need to enable the ingress. Here is `kind-config.yaml`

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
    kind: InitConfiguration
    nodeRegistration:
      kubeletExtraArgs:
        node-labels: "ingress-ready=true"
  extraPortMappings:
  - containerPort: 4242
    hostPort: 4242
    protocol: TCP
```

### Make Protocol Buffers

Install the needed tools to your local bin (already on the PATH for the development container):

```bash
make protoc
```

If you aren't in the dev container:

```
export PATH=$PATH:$PWD/bin
```

Then compile:

```bash
make proto
```

### Build Plugin

To install the plugin from GitHub:

```bash
go get github.com/converged-computing/fluence-kubectl
```

Or build it locally. This assumes you have a working `KUBECONFIG`

```bash
make
```

This will generate the binary, and you'll need to copy it into somewhere on the path to access.

```bash
sudo cp ./kubectl-fluence /usr/local/bin
```

### Shell Completion

For shell completion, copy [kubectl_complete-fluence](kubectl_complete-fluence) also somewhere on your
path and make it executable.

```bash
/usr/local/bin ./kubectl_complete-fluence
sudo cp ./kubectl_complete-fluence /usr/local/bin
```

### Build and Install

To do both of the above (and warning, the install requests sudo!)

```bash
make install
```

### Uninstall

To uninstall, just delete both.

```bash
sudo rm -rf /usr/local/bin/kubectl-fluence /usr/local/bin/kubectl_complete-fluence
```

## TODO

- I don't like that the grpcs are now shared across modules and projects - I'd like a simple strategy (maybe submodules) that can get them from the same source instead.
- We need a way to get group name / metadata directly from fluxion.
- We can add the Info/Stats exposed endpoints to the fluxion GRPC to get better views of the cluster (stats)

## Thank You

 - This plugin is based on logic from [kubernetes/sample-cli-plugin](https://github.com/kubernetes/sample-cli-plugin)

## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/converged-computing/cloud-select/blob/main/LICENSE),
[COPYRIGHT](https://github.com/converged-computing/cloud-select/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/converged-computing/cloud-select/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
