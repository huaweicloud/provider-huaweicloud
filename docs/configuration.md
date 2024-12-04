# HuaweiCloud Provider Documentation

## Install the provider

Install the HuaweiCloud provider with the following configuration file

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-huaweicloud
spec:
  package: xpkg.upbound.io/hcs/provider-huaweicloud:v0.0.7
```

Define the provider version with `spec.package`.

Install the provider with `kubectl apply -f`.

Verify the configuration with `kubectl get providers`.

```shell
$ kubectl get providers
NAME                      INSTALLED   HEALTHY   PACKAGE                                                    AGE
provider-huaweicloud      True        True      xpkg.upbound.io/hcs/provider-huaweicloud:v0.0.7    54s
```

View the Crossplane [Provider CRD definition](https://doc.crds.dev/github.com/huaweicloud/provider-huaweicloud) to view all available `Provider` options.

## Configure the provider

The HuaweiCloud provider requires credentials for authentication to Huawei Cloud. This can be done in one of the following ways:

* Authenticating using AK/SK

### Generate a Kubernetes Secret

Create a secret containing the Huawei Cloud account credentials (AK/SK).

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "region": "cn-north-4",
      "access_key": "your access key",
      "secret_key": "your access key"
    }
```

### Create a ProviderConfig object

Apply the secret in a `ProviderConfig` Kubernetes configuration file.

```yaml
apiVersion: huaweicloud.crossplane.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: example-creds
      namespace: crossplane-system
      key: credentials
```

**Note:** the `spec.credentials.secretRef.name` must match the `name` in the secret above.

The possible keys for the `spec.credentails` are:

| Key         | Type     | Required                      | Default |
| :---        | :------: | :---                          | :---    |
| region      | string   | true                          |         |
| access_key  | string   | true                          |         |
| secret_key  | string   | true                          |         |
