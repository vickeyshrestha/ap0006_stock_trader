**Maxzilla**

Maxzilla is a microservice platform that structures an application as a collection of services that are highly maintainable, loosely coupled, independently deployable and organized around business capabilities.


**prerequisite**
<br>1. _Nats server_ 
<br> -  Some of the services require message bus to operate. Please run Windows.bat file located under prerequisite-startup folder (Windows server only)
<br> - download and run the nats server on one of the VMs (later, this will be part of docker-compose.yaml):
<br> _docker run -d --name nats-main -p 4222:4222 -p 6222:6222 -p 8222:8222 nats_

<br>2. _Postgres sql_ 
<br>- requires Postgres db setup
<br>
<br>3. _mongo db_ 
<br>- requires mongo database db setup
