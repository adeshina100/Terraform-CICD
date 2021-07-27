---
page_title: "awsamplify_deploy_zip Resource - terraform-provider-awsamplify"
subcategory: ""
description: |-
    The `awsamplify_deploy_zip` resource mimics the manual deployment of a web app to AWS Amplify by submitting its build bundle zip archive.
    The triggers argument allows specifying an arbitrary set of values that, when changed, will cause the app to be redeployed.
---

# Resource: awsamplify_deploy_zip

The `awsamplify_deploy_zip` resource mimics the manual deployment of a web app to AWS Amplify by submitting its build bundle zip archive.
The triggers argument allows specifying an arbitrary set of values that, when changed, will cause the app to be redeployed.

## Example Usage

```terraform
# Location of the app build bundle archive.
data "local_file" "app_archive" {
  path = "./build/app.zip"
}

# The illustration of the app re-deployment based upon its codebase modifications.
resource "awsamplify_deploy_zip" "app" {
  # Changes of the app codebase would trigger its redeployment to AWS Amplify
  triggers = {
    hash = sha256(data.local_file.app_archive.content_base64)
  }

  # AWS Amplify Application ID
  app_id = "AWSAmplifyAppID"

  # Deployment branch
  branch_name = "master"

  # Location of the zip archive with the app
  path = data.local_file.app_archive.filename
}
```

## Schema

### Required

- **app_id** (String) Amplify app ID.

- **branch_name** (String) Amplify deployment branch.

- **path** (String) Path to zip archive with the app build.

### Optional

- **triggers** (Map of String) A map of arbitrary strings that, when changed, will force re-deployment.

### Read-Only

- **id** (String) This is set to a random value at create time.
