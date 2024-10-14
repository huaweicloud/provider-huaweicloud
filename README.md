# Provider HuaweiCloud

`provider-huaweicloud` is a [Crossplane](https://crossplane.io/) provider that
is built using [Upjet](https://github.com/crossplane/upjet) code
generation tools and exposes XRM-conformant managed resources for the
HuaweiCloud API.

## Getting Started

To start using this provider, follow the [Configuration Guide](docs/configuration.md).

You can find a detailed API reference with all CRDs [here](https://doc.crds.dev/github.com/huaweicloud/provider-huaweicloud).

## Developing

Run code-generation pipeline:
```console
go run cmd/generator/main.go "$PWD"
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/huaweicloud/provider-huaweicloud/issues).
