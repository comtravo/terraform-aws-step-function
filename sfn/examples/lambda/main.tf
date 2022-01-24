variable "sfn_name" {
  type = string
}

locals {
  lambda_name = "${var.sfn_name}-lambda"
}

data "aws_iam_policy_document" "sfn_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["states.amazonaws.com"]
    }
  }
}

data "aws_iam_policy_document" "lambda_assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
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
  assume_role_policy    = data.aws_iam_policy_document.sfn_assume_role.json
  force_detach_policies = true
}

resource "aws_iam_role_policy" "sfn" {
  name   = var.sfn_name
  role   = aws_iam_role.sfn.id
  policy = data.aws_iam_policy_document.policy.json
}


resource "aws_iam_role" "lambda" {
  name                  = local.lambda_name
  assume_role_policy    = data.aws_iam_policy_document.lambda_assume_role.json
  force_detach_policies = true
}

resource "aws_iam_role_policy" "lambda" {
  name = local.lambda_name
  role = aws_iam_role.lambda.id

  policy = data.aws_iam_policy_document.policy.json
}


locals {
  sfn_definition = {
    Comment = "SFN with lambda",
    StartAt = "FirstState",
    States = {
      FirstState = {
        Type     = "Task",
        Resource = module.lambda.arn,
        Next     = "wait_using_seconds"
      },
      wait_using_seconds = {
        Type    = "Wait",
        Seconds = 10,
        Next    = "FinalState"
      },
      FinalState = {
        Type     = "Task",
        Resource = module.lambda.arn,
        End      = true
      }
    }
  }
}

module "sfn" {

  source = "../../"

  role_arn = aws_iam_role.sfn.arn
  config = {
    name       = var.sfn_name
    definition = jsonencode(local.sfn_definition)
  }
}

output "arn" {
  value = module.sfn.state_machine_arn
}

module "lambda" {

  source = "github.com/comtravo/terraform-aws-lambda?ref=5.0.0"

  file_name     = "${path.module}/../../test/fixtures/foo.zip"
  function_name = local.lambda_name
  handler       = "index.handler"
  role          = aws_iam_role.lambda.arn
  trigger = {
    type = "step-function"
  }
  environment = {
    "LOREM" = "IPSUM"
  }
  region = "us-east-1"
  tags = {
    "Foo" : local.lambda_name
  }
}
