# Terraform AWS module for AWS Step Function.

## Introduction  
This module creates an AWS Step function.

## Usage  
Checkout [examples](./examples) on how to use this module.
## Authors

Module managed by [Comtravo](https://github.com/comtravo).

## License

MIT Licensed. See [LICENSE](./LICENSE) for full details.

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.13 |
| aws | ~> 2.0 |

## Providers

| Name | Version |
|------|---------|
| aws | ~> 2.0 |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| config | Step function configuration | <pre>object({<br>    name : string<br>    definition : string<br>  })</pre> | n/a | yes |
| role_arn | IAM role | `string` | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| state_machine_arn | AWS step function arn |

