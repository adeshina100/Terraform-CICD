# Terraform AWS Amplify Deployment Resource

AWS Amplify provides an endpoint to deploy web client as zip archive. As of 2021-07, the standard terraform AWS provider doesn't include the resource definition to deploy the app using zip archive. This module defines the resource "aws_amplify_deploy_zip".
