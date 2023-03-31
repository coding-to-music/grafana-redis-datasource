# grafana-redis-datasource

# 🚀 Redis Data Source for @grafana allows connecting to any @redis database On-Premises and in the Cloud. 🚀

https://github.com/coding-to-music/grafana-redis-datasource

From / By https://github.com/RedisGrafana/grafana-redis-datasource

https://redisgrafana.github.io/development/redis-datasource/

## Environment variables:

```java

```

## user interfaces:

- Grafana http://localhost:3000

## GitHub

```java
git init
git add .
git remote remove origin
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:coding-to-music/grafana-redis-datasource.git
git push -u origin main
```

# Redis Data Source for Grafana

![Dashboard](https://raw.githubusercontent.com/RedisGrafana/grafana-redis-datasource/master/src/img/redis-dashboard.png)

[![Grafana 8](https://img.shields.io/badge/Grafana-8-orange)](https://www.grafana.com)
[![Redis Data Source](https://img.shields.io/badge/dynamic/json?color=blue&label=Redis%20Data%20Source&query=%24.version&url=https%3A%2F%2Fgrafana.com%2Fapi%2Fplugins%2Fredis-datasource)](https://grafana.com/grafana/plugins/redis-datasource)
[![Redis Application plugin](https://img.shields.io/badge/dynamic/json?color=blue&label=Redis%20Application%20plugin&query=%24.version&url=https%3A%2F%2Fgrafana.com%2Fapi%2Fplugins%2Fredis-app)](https://grafana.com/grafana/plugins/redis-app)
[![Redis Explorer plugin](https://img.shields.io/badge/dynamic/json?color=blue&label=Redis%20Explorer%20plugin&query=%24.version&url=https%3A%2F%2Fgrafana.com%2Fapi%2Fplugins%2Fredis-explorer-app)](https://grafana.com/grafana/plugins/redis-explorer-app)
[![Go Report Card](https://goreportcard.com/badge/github.com/RedisGrafana/grafana-redis-datasource)](https://goreportcard.com/report/github.com/RedisGrafana/grafana-redis-datasource)
![CI](https://github.com/RedisGrafana/grafana-redis-datasource/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/RedisGrafana/grafana-redis-datasource/branch/master/graph/badge.svg?token=YX7995RPCF)](https://codecov.io/gh/RedisGrafana/grafana-redis-datasource)
[![Language grade: JavaScript](https://img.shields.io/lgtm/grade/javascript/g/RedisGrafana/grafana-redis-datasource.svg?logo=lgtm&logoWidth=18)](https://lgtm.com/projects/g/RedisGrafana/grafana-redis-datasource/context:javascript)

## Introduction

The Redis Data Source for Grafana is a plugin that allows users to connect to any Redis database On-Premises and in the Cloud. It provides out-of-the-box predefined dashboards and lets you build customized dashboards to monitor Redis and application data.

### Demo

Demo is available on [demo.volkovlabs.io](https://demo.volkovlabs.io):

- [Redis Overview dashboard](https://demo.volkovlabs.io/d/TgibHBv7z/redis-overview?orgId=1&refresh=1h)
- [Projects](https://demo.volkovlabs.io)

### Requirements

- **Grafana 8.0+** is required for Redis Data Source 2.X.
- **Grafana 7.1+** is required for Redis Data Source 1.X.

### Redis Application plugin

You can add as many data sources as you want to support multiple Redis databases. [Redis Application plugin](https://grafana.com/grafana/plugins/redis-app) helps manage various Redis Data Sources and provides Custom panels.

### Redis Explorer plugin

[The Redis Explorer plugin](https://grafana.com/grafana/plugins/redis-explorer-app) connects to Redis Enterprise software clusters using REST API. It provides application pages to add Redis Data Sources for managed databases and dashboards to see cluster configuration.

## Getting Started

Redis Data Source can be installed from the Grafana Marketplace or use the `grafana-cli` tool to install from the command line:

```bash
grafana-cli plugins install redis-datasource
```

![Grafana Marketplace](https://raw.githubusercontent.com/RedisGrafana/grafana-redis-datasource/master/src/img/grafana-marketplace.png)

For Docker instructions and installation without Internet access, follow the [Quickstart](https://redisgrafana.github.io/quickstart/) page.

### Configuration

Data Source allows to connect to Redis using TCP port, Unix socket, Cluster, Sentinel and supports SSL/TLS authentication. For detailed information, take a look at the [Configuration](https://redisgrafana.github.io/redis-datasource/configuration/) page.

![Datasource](https://raw.githubusercontent.com/RedisGrafana/grafana-redis-datasource/master/src/img/datasource.png)

## Documentation

Please take a look at the [Documentation](https://redisgrafana.github.io/redis-datasource/overview/) to learn more about plugin and features.

### Supported commands

List of all supported commands and how to use them with examples you can find in the [Commands](https://redisgrafana.github.io/redis-datasource/commands/) section.

![Query](https://raw.githubusercontent.com/RedisGrafana/grafana-redis-datasource/master/src/img/query.png)

## Development

[Developing Redis Data Source](https://redisgrafana.github.io/development/redis-datasource/) page provides instructions on building the data source.

Are you interested in the latest features and updates? Start nightly built [Docker image for Redis Application plugin](https://redisgrafana.github.io/development/images/), including Redis Data Source.

## Feedback

We love to hear from users, developers, and the whole community interested in this plugin. These are various ways to get in touch with us:

- Ask a question, request a new feature, and file a bug with [GitHub issues](https://github.com/RedisGrafana/grafana-redis-datasource/issues/new/choose).
- Star the repository to show your support.

## Contributing

- Fork the repository.
- Find an issue to work on and submit a pull request.
- Could not find an issue? Look for documentation, bugs, typos, and missing features.

## License

- Apache License Version 2.0, see [LICENSE](https://github.com/RedisGrafana/grafana-redis-datasource/blob/master/LICENSE).

## install go programming language

```
wget https://golang.org/dl/go1.17.6.linux-amd64.tar.gz

tar -xvf go1.17.6.linux-amd64.tar.gz

sudo mv go /usr/local
```

Set the environment variables required for Go to work correctly. Add the following lines to your ~/.bashrc file to set the Go environment variables:

```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
```

Developing Redis Data Source¶

https://redisgrafana.github.io/development/redis-datasource/

Developing Redis Data Source involves setting up the development environment (which can be either Linux-based or macOS-based), building and running tests.

## Install Grafana¶
Grafana can be started in Docker or installed locally:

Follow Installation instructions to install and start Grafana

Open Grafana UI in web-browser http://X.X.X.X:3000

## Clone repository¶
git clone https://github.com/RedisGrafana/grafana-redis-datasource.git

## Build¶

### Frontend¶

Install the latest version of Node.js using Node Version Manager or download binaries

- Install yarn globally

```
npm install yarn -g
```

Install dependencies
```
yarn install
```

Build frontend components

```
yarn build
```

## Backend¶

- Install Golang for your platform

```
yum install go
```

- Install Grafana plugin SDK for Go dependency

```
go get -u github.com/grafana/grafana-plugin-sdk-go
```

- Install Mage (make-like build tool using Go)

```
git clone https://github.com/magefile/mage
cd mage
go run bootstrap.go
```

- Build backend binaries for Linux, Windows and MacOS for supported platforms

```
yarn build:backend
```

## Start Grafana¶

Docker Compose
Prerequisite

Docker Compose should be pre-installed following documentation.

```
yarn start:dev
```

Update local Grafana Configuration

Move distribution to Grafana's plugins/ folder

```
mv dist/ /var/lib/grafana/plugins/redis-datasource
```

Add redis-datasource to allowed unsigned plugins

```
vi /etc/grafana/grafana.ini
```

```
[plugins]
;enable_alpha = false
;app_tls_skip_verify_insecure = false
# Enter a comma-separated list of plugin identifiers to identify plugins that are allowed to be loaded even if they lack a valid signature.
allow_loading_unsigned_plugins = redis-datasource
```

Restart Grafana and verify that plugin registered

```
tail -100 /var/log/grafana/grafana.log
```

```
t=2020-07-01T06:03:38+0000 lvl=info msg="Starting plugin search" logger=plugins
t=2020-07-01T06:03:38+0000 lvl=warn msg="Running an unsigned backend plugin" logger=plugins pluginID=redis-datasource pluginDir=/var/lib/grafana/plugins/redis-datasource
t=2020-07-01T06:03:38+0000 lvl=info msg="Registering plugin" logger=plugins name=redis-datasource
t=2020-07-01T06:03:38+0000 lvl=info msg="HTTP Server Listen" logger=http.server address=[::]:3000 protocol=http subUrl= socket=
```

## Configuration¶

The Redis Data Source Configuration page explains how to connect data source to Redis database.
