FROM gitpod/workspace-full:latest

RUN brew install go-task/tap/go-task

RUN brew tap go-swagger/go-swagger && brew install go-swagger

RUN brew install goreleaser
