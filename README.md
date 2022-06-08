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
      - "your_awsome_repository/"
```

## Build-Push Image

Once your repository connected with github actions. You can build-push your image to registry by creating tag with prefix `v`. For example
> v1.0.0
v0.0.1

# Deployments

WIP
