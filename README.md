## Requirements

| Name | Version |
|------|---------|
| terraform | >= 0.12 |
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

