# dgoss
>A simple dgoss image for CI Jobs

[![Build Status](https://travis-ci.org/iorubs/dgoss.svg?branch=master)](https://travis-ci.org/iorubs/dgoss)

## Description:

Goss is tool for validating server’s configuration (avoid conf. drift). Dgoss is wrapper written on top of the goss for validating docker images.
https://github.com/aelsabbahy/goss/tree/master/extras/dgoss


## Usage:

#### Setup:

```bash
# cd into the directory with your Dockerfile and build your image.
docker build -t app .

# dgoss is the entrypoint to this image, so running with no arguments
docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss
# Prints ERROR: USAGE: dgoss [run|edit] <docker_run_params>
```

#### Create tests:

```bash
docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss edit app

goss a process app
goss a user root
goss a user user2
goss a file /dir
goss a file /dir/file
goss a http http://localhost:8080
exit
```

#### Run tests:

```bash
docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss run app
```
