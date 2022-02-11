# Okta Golang Gin Resource Server Example

This example shows you how to use the [Okta JWT verifier library][] to secure a Golang REST API built with Gin.

## Prerequisites

Before running this sample, you will need the following:

- [Go 1.13 +](https://go.dev/dl/)
- [The Okta CLI Tool](https://github.com/okta/okta-cli/#installation)
- An Okta Developer Account, create one using `okta register`, or configure an existing one with `okta login`

## Get the Code

Grab and configure this project using `okta start go-api`

You can also clone this project from GitHub.

```bash
git clone https://github.com/okta-samples/okta-go-api-sample.git
cd okta-go-api-sample
```

> **Note**: Don't EVER commit `.okta.env` into source control. Add it to the `.gitignore` file.

## Run the Example

```bash
go run main.go
```

Now, send requests to http://localhost:8080.

```bash
curl -X GET  http://localhost:8080/api/hello

# or

curl -X GET  -H 'Authorization: Bearer ${TOKEN}’ http://localhost:8080/api/whoami
```

You can find more Golang sample in [this repository](https://github.com/okta/samples-golang)

[Okta JWT verifier library]: github.com/okta/okta-jwt-verifier-golang
