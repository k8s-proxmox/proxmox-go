# Proxmox VE REST API Client

Go client for the Proxmox VE REST API (https://pve.proxmox.com/wiki/Proxmox_VE_API)

## Overview

A Go package containing a client for [Proxmox VE](https://www.proxmox.com/). The client implements [/api2/json](https://pve.proxmox.com/pve-docs/api-viewer/index.html) and aims to provide better sdk solution for especially [cluster-api-provider-proxmox](https://github.com/k8s-proxmox/cluster-api-provider-proxmox) and [cloud-provider-proxmox](https://github.com/k8s-proxmox/cloud-provider-proxmox) project.

## Developing

### Unit Testing

```sh
go test ./... -v -skip ^TestSuiteIntegration
```

### Integration Testing

```sh
export PROXMOX_URL='http://localhost:8006/api2/json'
# tokenid & secret
export PROXMOX_TOKENID='root@pam!your-token-id'
export PROXMOX_SECRET='aaaaaaaaa-bbb-cccc-dddd-ef0123456789'
# or username & password
# export PROXMOX_USERNAME='root@pam'
# export PROXMOX_PASSWORD='password'

go test ./... -v -run ^TestSuiteIntegration
```
