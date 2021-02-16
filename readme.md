**Godzilla**

Godzilla is a microservice platform that structures an application as a collection of services that are highly maintainable, loosely coupled, independently deployable and organized around business capabilities.


**prerequisite**
<br>1. _Nats server_ 
<br> -  Some of the services require message bus to operate. Please run Windows.bat file located under prerequisite-startup folder (Windows server only) or you can simply bring up the nats docker image through docker-compose.yaml

<br>2. _Postgres sql_ 
<br>- requires Postgres db setup
<br>
<br>3. _mongo db_ 
<br>- requires mongo database db setup

**docker compose**
- docker-compose.yaml is not use for the internal godzilla repo but is for usage on the VM where all docker services will be running.
- setup docker-compose command on your linux VM using **sudo apt  install docker-compose**
- run docker-compose.yaml on one of the available servers to bring all of the image's instaces up and running
<br> _c:/> docker-compose -f docker-compose.yaml up -d_ (To bring up all docker instances)
<br> _c:/> docker-compose -f docker-compose.yaml down_ (To bring down all docker instances)