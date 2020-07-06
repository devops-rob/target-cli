# Target

## WIP - Tool is currently under development

A CLI tool to manage context profiles for HashiCorp products.  This allows you to save connection and configuration details, which would otherwise be set using ENV VARS into context profiles and gives the ability to switch between different profiles as required.

### Example use case

There are two vault clusters, one for Dev (<http://dev-vault:8200>) and one for Prod (<https://prod-vault:8200>).

Running Vault CLI commands locally will by default attempt to connect to <https://localhost:8200>.  To connect to another cluster, you need to know and remember the connection details and set a Vault environment variable. When you want to switch to a different cluster, you again need to set the environment variable.  Target allows you to save multiple connection details into context profiles and easily switch between them as you require.

### Example usage

```shell
target vault select dev
```
