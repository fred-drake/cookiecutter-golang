# Golang Cookie Cutter Template
This is the template I use to begin my Go projects.  It aims to have everything that I usually use:

- Visual Studio Code Remote Container support
- Installation of just the Visual Studio Code plugins that are needed for a Go project
- Github actions CI into a container

## Container Deployment
Upon push, Github will launch the CI pipeline that creates a container in ghcr.io.  If it is the master branch, it will tag it with `latest`.  If there is another branch name, it will tag with that branch's name.

It will create containers for both `amd64` and `arm64` architectures.

## To-Do's

- [ ] A more complete reference implementation
- [ ] Test suite to cover the reference implementation
- [ ] Add more container architectures
- [ ] Make CI configurable to not deploy container
