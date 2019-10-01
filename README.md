# PaySuper 1C Gateway

[![License: GPL 3.0](https://img.shields.io/badge/License-GPL3.0-green.svg)](https://opensource.org/licenses/Gpl3.0)
[![Go Report Card](https://goreportcard.com/badge/github.com/paysuper/paysuper-1c-gateway)](https://goreportcard.com/report/github.com/paysuper/paysuper-1c-gateway)

s# Motivation

1С gateway is a simple REST API proxy to handle transaction log with full set of data import from PaySuper billing 
server to 1C server. Works as very low load service with only some requests per day.

# Usage
Application designed to be launched with Kubernetes and handle part of configuration from env variables:

| Variable                            | Description                                                                                            |
|-------------------------------------|-------------------------------------------|
| GW_SERVER_PORT                      | HTTP port for billing server REST gateway.|
| GW_METRICS_PORT                     | HTTP port for metric and health check.    |
