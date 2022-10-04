# Build Stage
FROM golang:{{ cookiecutter.go_version }} AS base

FROM base as development-build-stage
ENV HADOLINT_VERSION={{ cookiecutter.hadolint_version }}
WORKDIR /tmp/builder
RUN dpkg --print-architecture > arch
RUN if [ "$(cat arch)" = "amd64" ]; then echo x86_64 > arch; fi
RUN curl -LO "https://github.com/hadolint/hadolint/releases/download/v${HADOLINT_VERSION}/hadolint-Linux-$(cat arch)" && \
    chmod +x "hadolint-Linux-$(cat arch)" && \
    mv "hadolint-Linux-$(cat arch)" /usr/local/bin/hadolint

FROM base as development
COPY --from=development-build-stage /usr/local/bin/hadolint /usr/local/bin/hadolint
WORKDIR /workspaces/{{ cookiecutter.project_name }}
RUN go install golang.org/x/tools/gopls@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

FROM base AS build-stage
LABEL app="build-{{ cookiecutter.project_name }}"
LABEL REPO="https://github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}"

ENV PROJPATH=/go/src/github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

WORKDIR /src
COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN mkdir /app && \
    go mod vendor && \
    go build -v -o /app ./...

# Final Stage
FROM busybox:glibc

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/{{ cookiecutter.github_name }}/{{ cookiecutter.project_name }}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/app

WORKDIR /app

COPY --from=build-stage /app/{{ cookiecutter.project_name }} /app/
RUN chmod +x /app/{{ cookiecutter.project_name }} && \
    adduser -D -g '' {{ cookiecutter.project_name }}
USER {{ cookiecutter.project_name }}

CMD ["/app/{{ cookiecutter.project_name }}"]
