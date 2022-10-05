# Golang Cookie Cutter Template
This is the template I use to begin my Go projects.  It aims to have everything that I usually use:

- Visual Studio Code Remote Container support
- Installation of just the Visual Studio Code plugins that are needed for a Go project
- Github actions CI into a container

## Project Structure
The template aims to follow the unofficial [golang standards project layout definition](https://github.com/golang-standards/project-layout).

## Formatting
The application will automatically format the code when a `go` file is saved.  It uses the [golines](https://github.com/segmentio/golines) library which runs the formatting of the standard `gofmt` but also breaks up long lines.  This gives a closer experience to those who are used to more fully-featured formatters such as Python's `black`.

## Containerized Development Environment
This template is intended to utilize Visual Studio Code's [remote container technology](https://code.visualstudio.com/docs/remote/containers).  The first time you load the project into VS Code, it will detect the remote container settings are in place and ask you if you wish to re-open in a container.  It will take a few moments for the container to build, but 
then you will have a full development ecosystem, complete with the exact versions of dependent software.  If you completely screw something up somewhere in your environment during development, hit `⌘⇧P`, select `Remote-Containers: Rebuild Container` and you can start again with a clean slate.

The only requirements for this are:
- The Visual Studio Code [Remote Development extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack)
- [Docker desktop](https://www.docker.com/products/docker-desktop/) (or equivalent third-party container technology, such as podman)

## Configuration
Configuration variables can be set the following ways, in the given order of priority:

1. Command line flags (highest preference order)
2. Environment variables
3. YAML config file
4. Hard-coded defaults

### Configuration Settings
In `main.go` you can configure the following structures to alter configuration behavior:

`cliFlags` is a list of strings that tie command-line parameters into configuration keys.  Example:
```go
	cliFlags := []string{
		"configFile",
		"logFormat",
		"logLevel",
	}
```

You are not limited to string types.  You may also add a colon-separated modifier altering the type. (for example, `age:int` or `notifyOnFailure:bool`).

`envToKeys` is a map of environment variables names to configuration keys.  This would allow you to use the standard format of environment variables.  Example:
```go
    envToKeys := map[string]string{
		"CONFIG_FILE": "configFile",
		"LOG_FORMAT": "logFormat",
		"LOG_LEVEL": "logLevel",
	}
```

Any application-level hard-coded default configuration values can be stored in `configDefaults`.  Example:

```go
	configDefaults := map[string]interface{}{
		"configFile": "/etc/my-project/config.yaml",
		"logFormat": "text",
		"logLevel": "info",
	}
```

Required configuration keys are set in the `requiredConfigKeys` list.  If a key is in this list but not defined anywhere, a fatal-level
log message will output and the program will exit.  Example:

```go
    requiredConfigKeys := []string{ "url", "username", "password" }
```

## Container Dockerfile
The `Dockerfile` is located in the `build/package` directory.  It is built into several phases so that the development tools you use are
separated from the deployment binary.

### Image Phases

#### Development
The `base` image is the [official Docker golang image](https://hub.docker.com/_/golang), using the `go_version` you select when calling `cookiecutter` on this template.

A `development-build-stage` image extends from the `base` to install [hadolint](https://github.com/hadolint/hadolint) based on the `hadolint_version` you give when calling `cookiecutter` on this template, and based on the machine's architecture.

A `development` image extends from the `base` which copies over the `hadolint` binary and installs the necessary `go` libraries for development.  This is the image used by Visual Studio Code.

### Deployment
A `build-stage` image extends from the `base` which downloads and verifies the `go` modules you have defined in your code.

The final image for deployment is a small `busybox` image which copies the binary that was built in the previous stage.

## Continuous Integration
Upon push, Github will launch the CI pipeline that creates a container in ghcr.io.  If it is the master branch, it will tag it with `latest`.  If there is another branch name, it will tag with that branch's name.

It will create containers for both `amd64` and `arm64` architectures.

## Packages Used
- [Logrus](https://github.com/sirupsen/logrus) logging framework
- [Gookit Config](https://github.com/gookit/config) configuration framework

## To-Do's
- [ ] A more complete reference implementation
- [ ] Test suite to cover the reference implementation
- [ ] Add more container architectures
- [ ] Make CI configurable to not deploy container
