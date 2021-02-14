[![Build Status](https://travis-ci.org/IBM/schematics-go-sdk.svg?branch=main)](https://travis-ci.org/IBM/schematics-go-sdk)

# IBM Cloud Schematics Go SDK 0.0.2
Go client library to interact with the various [IBM Cloud Schematics APIs](https://cloud.ibm.com/apidocs?category=schematics).

Disclaimer: this SDK is being released initially as a **pre-release** version.
Changes might occur which impact applications that use this SDK.

## Table of Contents
<!--
  The TOC below is generated using the `markdown-toc` node package.

      https://github.com/jonschlinkert/markdown-toc

  You should regenerate the TOC after making changes to this file.

      npx markdown-toc -i README.md
  -->

<!-- toc -->

- [IBM Cloud Schematics Go SDK 0.0.2](#ibm-cloud-schematics-go-sdk-002)
	- [Table of Contents](#table-of-contents)
	- [Overview](#overview)
	- [Prerequisites](#prerequisites)
	- [Installation](#installation)
			- [`go get` command](#go-get-command)
			- [Go modules](#go-modules)
			- [`dep` dependency manager](#dep-dependency-manager)
	- [Authentication](#authentication)
	- [Getting Started](#getting-started)
	- [Error handling](#error-handling)
	- [Using the SDK](#using-the-sdk)
	- [Questions](#questions)
	- [Issues](#issues)
	- [Open source @ IBM](#open-source--ibm)
	- [Contributing](#contributing)
	- [License](#license)

<!-- tocstop -->

## Overview

The IBM Cloud Schematics Go SDK allows developers to programmatically interact with the following IBM Cloud services:

Service Name | Package name 
--- | --- 
[Schematics](https://cloud.ibm.com/apidocs/schematics) | schematicsv1 

## Prerequisites

[ibm-cloud-onboarding]: https://cloud.ibm.com/registration

* An [IBM Cloud][ibm-cloud-onboarding] account.
* An IAM API key to allow the SDK to access your account. Create one [here](https://cloud.ibm.com/iam/apikeys).
* Go version 1.12 or above.

## Installation
The current version of this SDK: 0.0.2

There are a few different ways to download and install the Schematics Go SDK project for use by your
Go application:

#### `go get` command  
Use this command to download and install the SDK to allow your Go application to
use it:

```
go get -u github.com/IBM/schematics-go-sdk
```

#### Go modules  
If your application is using Go modules, you can add a suitable import to your
Go application, like this:

```go
import (
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)
```

then run `go mod tidy` to download and install the new dependency and update your Go application's
`go.mod` file.

#### `dep` dependency manager  
If your application is using the `dep` dependency management tool, you can add a dependency
to your `Gopkg.toml` file.  Here is an example:

```
[[constraint]]
  name = "github.com/IBM/schematics-go-sdk"
  version = "0.0.2"

```

then run `dep ensure`.

## Authentication

The library requires Identity and Access Management (IAM) to authenticate requests. There are several ways to set the properties for authentication

1. As environment variables
2. The programmatic approach
3. With an external credentials file

Authenticate with environment variables

For Schematics IAM authentication set the following environmental variables by replacing <apikey> with your proper service credentials. 

```
SCHEMATICS_URL = https://schematics.cloud.ibm.com
SCHEMATICS_APIKEY = <apikey>
```

Authenticate with external configuration

To use an external configuration file, see the related documentation in the [Go SDK Core document about authentication](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md).

Authenticate programmatically

To learn more about how to use programmatic authentication, see the related documentation in the [Go SDK Core document about authentication](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md).


## Getting Started

A quick example to get you up and running with Schematics Go SDK service

```
package main

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/go-sdk-core/v4/core"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func main() {

	schematicsService, _ := schematicsv1.NewSchematicsV1(&schematicsv1.SchematicsV1Options{
		URL: "https://schematics.cloud.ibm.com",
		Authenticator: &core.IamAuthenticator{
			URL:    "https://iam.cloud.ibm.com/identity/token",
			ApiKey: ". . . . . .",
		},
	})

	version, detailedResponse, err := schematicsService.GetSchematicsVersion(&schematicsv1.GetSchematicsVersionOptions{})

	if err != nil {
		fmt.Printf("Failed to get the version : %v and the response is %s", err, detailedResponse)
	}

	j, err := json.Marshal(version)
	fmt.Println("Schematics Version:", string(j))

	result, detailedResponse, err := schematicsService.ListWorkspaces(&schematicsv1.ListWorkspacesOptions{})

	if err != nil {
		fmt.Printf("Failed to list the workspaces : %v and the response is %s", err, detailedResponse)
	}

	w, err := json.Marshal(result)
	fmt.Println("Workspaces:", string(w))

}
```

## Error handling

For sample code on handling errors, please see [Schematics API docs](https://cloud.ibm.com/apidocs/schematics#error-handling).

## Using the SDK
For general SDK usage information, please see [this link](https://github.com/IBM/ibm-cloud-sdk-common/blob/master/README.md)

## Questions

If you are having difficulties using this SDK or have a question about the IBM Cloud services,
please ask a question at 
[Stack Overflow](http://stackoverflow.com/questions/ask?tags=ibm-cloud).

## Issues
If you encounter an issue with the project, you are welcome to submit a
[bug report](https://github.com/IBM/schematics-go-sdk/issues).
Before that, please search for similar issues. It's possible that someone has already reported the problem.

## Open source @ IBM
Find more open source projects on the [IBM Github Page](http://ibm.github.io/)

## Contributing
See [CONTRIBUTING](CONTRIBUTING.md).

## License

This SDK project is released under the Apache 2.0 license.
The license's full text can be found in [LICENSE](LICENSE).
