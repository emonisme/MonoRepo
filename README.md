# MonoRepo
Mono Repository Example

## Setup github actions
> Currently only support golang

If you want to add your repository to github action, you need to add your repository on [.github/workflows/](.github/workflows)[[your languange]].yaml. Put it under `on.push.paths`
```
name: golang
on:
  push:
    paths:
      - "golang_simple_app/"
      - "your_awsome_repository/" # add this
  pull_request:
    paths:
      - "golang_simple_app/"
      - "your_awsome_repository/" # add this
```

## Build-Push Image

Once your repository connected with github actions. You can build-push your image to registry by creating tag with prefix `v`. For example
> v1.0.0

# Deployments

Before you deploy your microservices to kubernetes, make sure to create kubernetes manifest using Helm (install it with this [guide](https://helm.sh/docs/intro/install/)). This will create manifest template to deploy it to kubernetes
```
cd helm-charts
helm create charts/[[your_awsome_repository]]
```
Adjust templates and values if needed. Dry run helm to check if your template is correct or not inside helm-charts folder
```
helm install --debug --dry-run [[Name version]] ./charts/[[your_awsome_repository]]
```
if it is ok, run actions deploy to deploy it to kubernetes [WIP]
