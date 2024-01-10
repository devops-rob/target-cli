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

### Example Use Case

Two Vault clusters are available for development (`http://dev-vault:8200`) and production (`https://prod-vault:8200`). By default, local Vault CLI commands attempt to connect to `https://localhost:8200`. Connecting to a different cluster requires remembering and setting the appropriate Vault environment variable each time.

Target simplifies this process. It allows you to store the connection details for multiple clusters in [context profiles](#what-is-a-context-profile). Switching between clusters is as easy as selecting the desired profile.

### What Is a Context Profile?

A **context profile** is a group of configuration parameters used to perform an action against one of the supported HashiCorp products.

For example, a Vault context profile is made of:

- `endpoint` set to `https://prod-vault:8200`
- `namespace` set to `admin/target`
- `token` set to `s.12345790asdfghjklpoi`

Which is translated to the following `export` commands:

```shell
export VAULT_ADDR=https://prod-vault:8200; export VAULT_NAMESPACE=admin/target; export VAULT_TOKEN=s.12345790asdfghjklpoi
```

This can then be used with the `eval` command to set these environment variables in the current shell's session.

### Example Usage

```shell
eval $(target vault select prod)
```

### Supported Tools

- [Terraform](#terraform)
- [Vault](#vault-consul-nomad-and-boundary)
- [Boundary](#vault-consul-nomad-and-boundary)
- [Consul](#vault-consul-nomad-and-boundary)
- [Nomad](#vault-consul-nomad-and-boundary)

### Configuring Target CLI For Your Shell

Target CLI allows engineers to set default context profiles that are automatically loaded into your shell session via Environment Variables. In order for this to work, you must configure your shell's start up scripts to load in Target CLI defaults.

For example, to configure Target CLI for zsh, we need to point the CLI at our `zshrc file`. For most users, this will be located at `~/.zshrc`.

To configure Target CLI for this shell type, run the following command:

```shell
target configure --path "~/.zshrc"
```

This will write a small helper script to the file that will come into effect when a new shell session is loaded.

### Terraform

This is designed to store sets of Terraform variables into profiles to allow for easy switching. 

#### Example Use case

Let's say there are three environments, dev, test, and prod. Each environment is deployed and configured with Terraform. Configuration parameters are set using Terraform variables; however, each environment should be configured with its own naming convention. 

- `dev` has resource names prefixed with `dev-`
- `test` has resource names prefixed with `test-`
- `prod` has resource names prefixed with `prod-`

To configure a set of variables for each environment, we can set up a `dev` context profile where the variable values are set accordingly, and do the same for `test` and `prod` context profiles. This means to deploy to a specific environment, an operator can switch context before applying their Terraform code.

**Create Example**

```shell
target terraform create dev \
  --var "aws_region=eu-west1" \
  --var "vpc_name=dev-vpc" \
  --var "ec2_instance_name=dev-droid-vm"
```

This will create a context profile named `dev` with 3 Terraform variables, `aws_region`, `vpc_name`, and `ec2_instance_name`

**Update Example**

```shell
target terraform update example \
  --var "aws_region=eu-west2" \
  --var "vpc_name=dev" \
  --var "ec2_instance_name=dev-starship-vm"
```

This will update the values of the `dev` context profile created in the previous step.

**Delete Example**

```shell
target terraform delete dev
```

This will delete the context profile named `dev`.

### Vault, Consul, Nomad and Boundary

Each of these tool's subcommands work in the same way using the available flags for each tool specific argument. 

Below are some examples using the `vault` subcommand:

**Create Example**

```shell
target vault create staging \
  --endpoint "https://staging-vault.target:8200" \
  --cacert "your CA cert" \
  --cert "your cert" \
  --key "your key" \
  --skip-verify false \
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
  --skip-verify false \
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

To set or change the default context profile for a tool, use the `set-default` subcommand:

```shell
target vault set-default staging
```

After setting a default, you'll need to start a new shell session for the changes to take effect. This is because the environment variables in the current shell session won't automatically update. In all new shell sessions, the environment variables will be set according to your defaults.

### Switching Context profiles

To switch between different context profiles, use the `select` subcommand as shown below:

```shell
target vault select dev
```

This subcommand prints all `export` commands for the selected context profile (`dev` in this case). To apply these `export` commands to your current shell session, wrap the `select` subcommand in an `eval` command like so:

```shell
eval $(target vault select dev)
```

The `eval` command executes the `export` commands, setting the environment variables for the selected context profile in your current shell session.
