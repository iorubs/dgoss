# dgoss
>A simple dgoss image for CI Jobs

### Description:

Goss is tool for validating server’s configuration (avoid conf. drift). Dgoss is wrapper written on top of the goss for validating docker images.
https://github.com/aelsabbahy/goss/tree/master/extras/dgoss


### Usage:

``` bash
# cd into the directory with your Dockerfile and build your image.
docker build -t app .

# dgoss is the entrypoint to this image, so running with no arguments
docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss
# Prints ERROR: USAGE: dgoss [run|edit] <docker_run_params>

# Add tests
docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss edit app

goss a process app
goss a user root
goss a user node
goss a file /app
goss a file /app/app
goss a http http://localhost:8080
exit

# Run tests
docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss edit app run app
```
