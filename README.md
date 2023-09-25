# Target CLI

A CLI tool to manage context profiles for HashiCorp products.  This allows you to save connection and configuration details, which would otherwise be set using Environment Variables into context profiles and gives the ability to switch between different profiles as required.

### Example use case

There are two vault clusters, one for Dev (<http://dev-vault:8200>) and one for Prod (<https://prod-vault:8200>).

Running Vault CLI commands locally will by default attempt to connect to <https://localhost:8200>.  To connect to another cluster, you need to know and remember the connection details and set a Vault environment variable. When you want to switch to a different cluster, you again need to set the environment variable.  Target allows you to save multiple connection details into context profiles and easily switch between them as you require.

### What Is a Context Profile?

A context profile is a grouping of configuration parameters required to perform an action against one of the supported HashiCorp Tools. A vault example is a context profile made of an `endpoint` pointing to `https://prod-vault:8200`, a `namespace` pointing to `admin/target`, and a `token` pointing to `s.12345790asdfghjklpoi`. This would render the following commands `export VAULT_ADDR=https://prod-vault:8200; export VAULT_NAMESPACE=admin/target; export VAULT_TOKEN=s.12345790asdfghjklpoi`. This can then be used with the `eval` command to set these environment variables in the current shell session.
### Example usage

```shell
target vault select prod
```

### Current Supported Tools

- Terraform
- Vault
- Boundary
- Consul
- Nomad

### Terraform

This is designed to store sets of terraform variables into profiles to allow for easy switching.

**Create Example**

```shell
target terraform create example \
  --var "aws_region=eu-west1" \
  --var "vpc_name=main" \
  --var "ec2_instance_name=droid-vm"
```

This will create a context profile named `example` with 3 terraform variables, `aws_region`, `vpc_name`, and `ec2_instance_name`

**Update Example**

```shell
target terraform update example \
  --var "aws_region=eu-west2" \
  --var "vpc_name=dev" \
  --var "ec2_instance_name=starship-vm"
```

This will update the values of the `example` context profile created in the previous step
