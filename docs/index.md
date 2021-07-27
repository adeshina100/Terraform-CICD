---
page_title: "Provider: AWSAmplify"
description: |-
  The AWSAmplify provider provides the resource to deploy the web app from its build zip archive to AWS Amplify environment.
---

# AWSAmplify Provider

The `awsamplify` provider the option of web apps deployment to [AWS Amplify](https://aws.amazon.com/amplify/) frontend environment by submitting [zip archive of the app build bundle](https://docs.aws.amazon.com/amplify/latest/userguide/manual-deploys.html).

## Example Usage

```terraform
terraform {
  required_providers {
    awsamplify = {
      version = "~> 1.0"
      source  = "hashicorp.com/kislerdm/awsamplify"
    }
  }
}

# AWS Region where to deploy the app
provider "awsamplify" {
  region = "eu-west-1"
}
```

## Schema
