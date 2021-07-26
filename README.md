# Terraform `awsamplify` Provider

AWS Amplify provides an endpoint to deploy web client as zip archive. As of 2021-07, the standard terraform AWS provider doesn't include the resource definition to deploy the app using zip archive. This module defines the resource "awsamplify_deploy_zip".

## Requirements

-	[Terraform](https://www.terraform.io/downloads.html) 0.15.x
-	[Go](https://golang.org/doc/install) 1.16 (to build and develop the provider plugin)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/kislerdm/terraform-provider-awsamplify`

```bash
git clone git@github.com:terraform-providers/terraform-provider-awsamplify $GOPATH/src/github.com/kislerdm/terraform-provider-awsamplify
```

Enter the provider directory and build the provider

```bash
cd $GOPATH/src/github.com/kislerdm/terraform-provider-awsamplify
make build
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.16+ is *required*).

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

In order to test the provider, you can simply run `make test`.

```bash
make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```bash
make testacc
```

## Install the Provider

To install the provider from repo, run the command

```bash
make install
```

It will build the binary for MacOS, ARM CPU architecture and move it to `~/.terraform.d/plugins/`. If your OS and ARCH differ from default, you can run the command

```bash
make install OS_ARCH=##GOOS##_##GOARCH##
```

For example, for linux with amd64, run

```bash
make install OS_ARCH=linux_amd64
```
