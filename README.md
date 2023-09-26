# Target CLI

A CLI tool to manage context profiles for HashiCorp products.  This allows you to save connection and configuration details, which would otherwise be set using Environment Variables into context profiles and gives the ability to switch between different profiles as required.

### Example use case

There are two vault clusters, one for Dev (<http://dev-vault:8200>) and one for Prod (<https://prod-vault:8200>).

Running Vault CLI commands locally will by default attempt to connect to <https://localhost:8200>.  To connect to another cluster, you need to know and remember the connection details and set a Vault environment variable. When you want to switch to a different cluster, you again need to set the environment variable.  Target allows you to save multiple connection details into context profiles and easily switch between them as you require.

### What Is a Context Profile?

A context profile is a grouping of configuration parameters required to perform an action against one of the supported HashiCorp Tools. A vault example is a context profile made of an `endpoint` pointing to `https://prod-vault:8200`, a `namespace` pointing to `admin/target`, and a `token` pointing to `s.12345790asdfghjklpoi`. This would render the following commands `export VAULT_ADDR=https://prod-vault:8200; export VAULT_NAMESPACE=admin/target; export VAULT_TOKEN=s.12345790asdfghjklpoi`. This can then be used with the `eval` command to set these environment variables in the current shell session.
### Example usage

```shell
eval $(target vault select prod)
```

### Current Supported Tools

- Terraform
- Vault
- Boundary
- Consul
- Nomad

### Configuring Target CLI For Your Shell

Target CLI allows engineers to set default context profiles that are automatically loaded into your shell session via Environment Variables. In order for this to work, you must configure your shell's start up scripts to load in Target CLI defaults.

For example, to configure Target CLI for zsh, we need to point the CLI at our `zshrc file`. For most users, this will be located at `~/.zshrc`.

To configure Target CLI for this shell type, run the following command

```shell
target configure --path "~/.zshrc"
```

This will write a small helper script to the file that will come into effect when a new shell session is loaded.

### Terraform

This is designed to store sets of terraform variables into profiles to allow for easy switching. 

### Example Use case

Let's say there are three environments, dev, test, and prod. Each environment is deployed and configured with Terraform. Configuration parameters are set using Terraform variables; however, each environment should be configured with its own naming convention. 

- `dev` has resource names prefixed with `dev-`
- `test` has resource names prefixed with `test-`
- `prod` has resource names prefixed with `prod-`

To configure a set of variables for each environment, we can set up a dev context profile where the variable values are set accordingly, and do the same for test and prod context profiles. This means to deploy to a specific environment, an operator can switch context before applying their terraform code.

**Create Example**

```shell
target terraform create dev \
  --var "aws_region=eu-west1" \
  --var "vpc_name=dev-vpc" \
  --var "ec2_instance_name=dev-droid-vm"
```

This will create a context profile named `dev` with 3 terraform variables, `aws_region`, `vpc_name`, and `ec2_instance_name`

**Update Example**

```shell
target terraform update example \
  --var "aws_region=eu-west2" \
  --var "vpc_name=dev" \
  --var "ec2_instance_name=dev-starship-vm"
```

This will update the values of the `dev` context profile created in the previous step

**Delete Example**

```shell
target terraform delete dev
```

This will delete the context profile named `dev`