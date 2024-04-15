# Air Helm

[![Docker Image Version](https://img.shields.io/docker/v/gdurandvadas/air-helm?arch=arm64&logo=docker&logoColor=white&label=Docker%20Hub&labelColor=%231D63ED&color=%23E5F2FC)](https://hub.docker.com/r/gdurandvadas/air-helm/tags)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/air-helm)](https://artifacthub.io/packages/search?repo=air-helm)

## Introduction

Welcome to Air Helm, a tool crafted to support the development of Go applications within Kubernetes environments.

Modern API servers often require a complex assembly of infrastructure components, and Air Helm is designed to navigate this complexity from your local setup. Air Helm integrates [Air](https://github.com/cosmtrek/air) for live reloading and [Helm](https://helm.sh) for managing Kubernetes applications, facilitating the development and testing of Go applications in a local Kubernetes environment. This synergy enables developers to work on their Go applications and accurately define and interact with the necessary architecture components.

## Key Features

- **Live Reloading**: Instant code reloads in Kubernetes without restarting pods.
- **Local Development Ease**: Helm chart simplifies deployment and management.
- **Architecture Integration**: Develop in a local environment as it was production.

## How It Works

Air Helm works by mounting your local Go project directory into a Kubernetes pod. This allows Air to monitor changes to the files and trigger rebuilds and restarts automatically. Here's what happens under the hood:

- **Local Directory Mounting**: Your local project directory is mounted into the Kubernetes pod, allowing the container to access your Go code in real time.
- **Hot reloading**: Air watches for file changes in your project and triggers builds and restarts as needed, ensuring that the latest changes are always running.
- **Helm for Kubernetes**: Helm manages the deployment of these settings into your local Kubernetes, configuring things like service exposure, persistent volumes, and more.

## Getting Started

### Prerequisites

- Docker
- Kubernetes cluster (e.g., Minikube or Docker Desktop's Kubernetes)
- Helm 3.x

### Installation

Add the Air Helm repository and install the Helm chart:

```bash
helm repo add gdurandvadas https://gdurandvadas.github.io/air-helm/
helm repo update
helm install my-local-env gdurandvadas/air-helm
```

### Configuration

Here's how you can configure the Air Helm chart in your `helmfile.yaml`:

```yaml
repositories:
  - name: gdurandvadas
    url: https://gdurandvadas.github.io/air-helm/

releases:
- name: api
  namespace: backend
  chart: gdurandvadas/air-helm
  values:
    - airServer:
        code:
          directory: "/Users/{username}/project" # Root directory of your project
          modulePath: "example/api" # Module path relative to the project root
          buildSource: "./cmd/server" # Path to the entry point of the application
    - service:
        type: LoadBalancer # Expose the service on a public IP
        http:
          enabled: true
          port: 8080
    - image:
        tag: commit # Docker image tag to use
```

You can see a full API example in the [example](./example/api/) directory

### Accessing Your Application

Depending on your service configuration, you can access your application typically via a local URL or an IP exposed by Kubernetes.

## Advanced Configuration

Review the `values.yaml` file for detailed options like autoscaling, ingress settings, and more. Tailor these settings to fit the needs of your development and production environments.

## License

Air Helm is open-sourced under the Apache License 2.0. For more details, see the [Apache License 2.0](http://www.apache.org/licenses/LICENSE-2.0)
