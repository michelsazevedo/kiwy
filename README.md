# Kiwy

Kiwy is a proof of concept to validate [Google BigTable](https://cloud.google.com/bigtable/) performance.

## Built With

- [Go](https://golang.org/)

Plus *some* of packages, a complete list of which is at [/master/go.mod](https://github.com/michelsazevedo/kiwy/blob/master/go.mod).

## Instructions

### Dependencies

#### Running with Docker
[Docker](www.docker.com) is an open platform for developers and sysadmins to build, ship, and run distributed applications, whether on laptops, data center VMs, or the cloud.

If you haven't used Docker before, it would be good idea to read this article first: Install [Docker Engine](https://docs.docker.com/engine/installation/)

1. Install [Docker](https://www.docker.com/what-docker) and then [Docker Compose](https://docs.docker.com/compose/):

2. Run `docker-compose build --no-cache` to build the image for the project.

3. Finally, run the local app with `docker-compose run app` and kiwy will perform requests.

4. Aaaaand, you can run the automated tests suite running a `docker-compose run --rm test` with no other parameters!

## License
Copyright © 2019
