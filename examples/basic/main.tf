variable "sfn_name" {
  type = string
}

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["states.amazonaws.com"]
    }
  }
}

# Do not use the below policy anywhere
data "aws_iam_policy_document" "policy" {
  statement {
    actions   = ["*"]
    resources = ["*"]
  }
}

resource "aws_iam_role" "sfn" {
  name                  = var.sfn_name
  assume_role_policy    = data.aws_iam_policy_document.assume_role.json
  force_detach_policies = true
}

resource "aws_iam_role_policy" "lambda" {
  name = var.sfn_name
  role = aws_iam_role.sfn.id

  policy = data.aws_iam_policy_document.policy.json
}


module "sfn" {

  source = "../../"

  role_arn =  aws_iam_role.sfn.arn
  config = {
    name = var.sfn_name
    definition = <<DEF
{
  "Comment": "A Hello World example of the Amazon States Language using Pass states",
  "StartAt": "Hello",
  "States": {
    "Hello": {
      "Type": "Pass",
      "Result": "Hello",
      "Next": "World"
    },
    "World": {
      "Type": "Pass",
      "Result": "World",
      "End": true
    }
  }
}
DEF
  }
}

output "arn" {
  value       = module.sfn.state_machine_arn
}
