# qualitytrace

This is the Helm chart for [qualitytrace] installation.

## Installation

### Chart installation

Add repo:

```sh
helm repo add qualitytrace https://intelops.github.io/qualitytrace
helm repo update

```

```sh
helm install qualitytrace qualitytrace/qualitytrace
```

## Uninstall

```sh
helm delete qualitytrace
```
