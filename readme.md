**Godzilla**

Godzilla is a microservice platform that structures an application as a collection of services that are highly maintainable, loosely coupled, independently deployable and organized around business capabilities.


**prerequisite software**
1. Docker
2. Nats (we can use a sample server as defined in docker-compose.yaml)
3. Mongodb - to store some KV pairs
4. Postgres db (for storing our big data)
5. Dockerhub account to push images

**docker compose**
- docker-compose.yaml is not for the internal godzilla repo but is for usage on the VM where all docker services will be running.
- setup docker-compose command on your linux VM using **sudo apt  install docker-compose**
- run docker-compose.yaml on one of the available servers to bring all of the image's instaces up and running
<br> _c:/> docker-compose -f docker-compose.yaml up -d_ (To bring up all docker instances)
<br> _c:/> docker-compose -f docker-compose.yaml down_ (To bring down all docker instances)