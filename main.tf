/**
* # Terraform AWS module for AWS Step Function.
*
* ## Introduction
* This module creates an AWS Step function.
*
* ## Usage
* Checkout [examples](./examples) on how to use this module.
* ## Authors
*
* Module managed by [Comtravo](https://github.com/comtravo).
*
* ## License
*
* MIT Licensed. See [LICENSE](./LICENSE) for full details.
*/

variable config {
  type = object({
    name : string
    definition : string
  })
  description = "Step function configuration"
}

variable role_arn {
  type        = string
  description = "IAM role"
}

resource "aws_sfn_state_machine" "sfn_state_machine" {
  name       = var.config.name
  role_arn   = var.role_arn
  definition = var.config.definition
}

output state_machine_arn {
  value       = aws_sfn_state_machine.sfn_state_machine.id
  description = "AWS step function arn"
}
