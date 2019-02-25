# dgoss
>A simple dgoss image for CI Jobs

### Description:

Goss is tool for validating server’s configuration (avoid conf. drift). Dgoss is wrapper written on top of the goss for validating docker images.
https://github.com/aelsabbahy/goss/tree/master/extras/dgoss


### Usage:
cd example

docker build -t app .

docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss

ERROR: USAGE: dgoss [run|edit] <docker_run_params>


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

docker run --rm -it \
  -v "$(pwd)":/src \
  -v /var/run/docker.sock:/var/run/docker.sock \
iorubs/dgoss edit app run app
