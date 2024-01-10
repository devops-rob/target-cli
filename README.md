# Target CLI

[![goreleaser](https://github.com/devops-rob/target-cli/actions/workflows/release.yaml/badge.svg)](https://github.com/devops-rob/target-cli/actions/workflows/release.yaml)

**Target CLI** is a tool designed for developers using HashiCorp products across multiple environments such as development, testing, and production. It eliminates the tedious task of manually setting and unsetting environment variables for each environment. With Target CLI, you can easily store and switch between different [context profiles](#what-is-a-context-profile), streamlining your workflow.

### Installation

**Mac / Linux via Brew**

```shell
brew tap devops-rob/tap && \
  brew install target
```

**Linux Quick Install**

```shell
curl https://raw.githubusercontent.com/devops-rob/target-cli/main/install.sh | bash
```

**Releases**

Binaries can be downloaded from the releases page:

[https://github.com/devops-rob/target-cli/releases](https://github.com/devops-rob/target-cli/releases)

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

#### Example Use case

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

### Vault, Consul, Nomad and Boundary

Each of these tool sub commands work in the same way using the available flags for each tool specific argument. Below are some examples using Vault

**Create Example**

```shell
target vault create staging \
  --endpoint "https://staging-vault.target:8200" \
  --cacert "your CA cert" \
  --cert "your cert" \
  --key "your key" \
  --skip-verify true \
  --cli-no-colour true \
  --client-timeout "60s" \
  --format "json"
```

**Update Example**

```shell
target vault create staging \
  --endpoint "https://staging-vault.com:8200" \
  --cacert "your new CA cert" \
  --cert "your cert" \
  --key "your key" \
  --skip-verify true \
  --format "json"
```

**Delete Example**

```shell
target vault delete staging
```

**List Example**

```shell
target vault list
```

### Setting Default Context Profiles

Setting the default or changing the default context profile for a tool can be done with the `set-default` sub command. 

```shell
target vault set-default staging
```

Once a default has been set, a new shell will need to be launch in order for the changes to take effect. All new shell session swill spawn with your defaults as set.

### Switching Context profiles

This can be done using the `select` sub command:

```shell
target vault select dev
```

This will print all `export` commands for the selected context profile. 

In order for this to take effect in the current shell session, the above command will need to be wrapped in an `eval` command:

```shell
eval $(target vault select dev)
```