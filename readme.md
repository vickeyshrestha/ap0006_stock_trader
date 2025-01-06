<h1>Stockzilla</h1>

Stockzilla is a **microservice** platform that structures an application as a collection of services that are highly maintainable, loosely coupled, independently deployable, and organized around business capabilities.


<h2>List of Services</h2>

**Each of the following services evolves and deploys independently, reducing the risk and time associated with coordinating changes across an entire application.
Please refer to each service's readme for further detail**


| Service                                                 |                                                         Short Description                                                         | 
|---------------------------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------:| 
| Stock Trader                                            |   subscribes to a specific topic in NATS to get Stock Exchange data in a TimeSeries fashion and inserts them into Postgres db [More on ReadMe](https://github.com/vickeyshrestha/stockzilla/blob/master/services/stock-trader/readme.md)  | 
| Mongo Engine                                            |                                 connects to MongoDB to fetch the KV configurations for the client                                 | 
| File Transfer Engine                                    |                              utilizes byte streaming mechanism to transfer a huge-sized single file                               |
| Simulators/stock (Not really a Service, but a simulator) | publishes random numbers for companies listed under stock exchange into messagebus (NATS) to be consumed by other client services [More On ReadMe](https://github.com/vickeyshrestha/stockzilla/blob/master/simulators/stock/readme.md) |

<h3>prerequisite software</h3>
1. Docker is all you need if you want to just run and test.

The following dependencies should be taken care when docker-compose.yaml will run. Basically all services withing this microservice infrastructure depends on the following tools.
* Nats (we can use a sample server as defined in docker-compose.yaml)
* Mongodb - to store some KV pairs
* Postgres db (for storing our big data)
* Hasura (GraphQL server that gives instant, realtime GraphQL APIs over Postgres)

Now for dev purpose, you might need:
2. Goose for DB migrations - 
<br>`go install github.com/pressly/goose/v3/cmd/goose@latest`


3. Some tools to make things easy, espcially for Windows OS:
* Chocolatey, a package manager for Windows only 
<br> `Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))`
* Make command (Windows) - 
<br>`choco install make`

<h3>docker compose</h3>
- docker-compose.yaml can be used on the server to bring up all the required containers 
- setup docker-compose command on your linux VM using **sudo apt  install docker-compose**
- run docker-compose.yaml on one of the available servers to bring all the image's instances up and running
<br> `c:/> docker-compose -f docker-compose.yaml up -d` (To bring up all docker instances)
<br> `c:/> docker-compose -f docker-compose.yaml down` (To bring down all docker instances)
- For easiness, a Makefile is also included

<h3>Jenkins Pipeline</h3>
- comes as a container
- go to set localhost port `localhost:8082`
- for password - `docker exec jenkins-blueocean cat /var/jenkins_home/secrets/initialAdminPassword`
  
![img.png](img.png)

You can inspect the logs using regular docker logs <container_id> command

<h3>For stock trader, and stock-simulator</h3>
- use the included NATS-WebUI to monitor the nats server:
1. Go to http://localhost:8091/
2. Click Add Server button on top right
3. Set:
 <br>   `Hostname: stockzilla-core-messagebus`
 <br>   `Message Port: 4222`
 <br>   `Monitoring Port: 8222`
 <br> The UI tool interacts with a NATS server by connecting to its client port (default: 4222) to retrieve and display information about the NATS server and its JetStream features.
 <br> ![img2.png](img2.png)

**please refer to each service's readme for further detail**

<h3>troubleshooting</h3>
If you are having difficulty to sync dependencies from vickeyshrestha/sharing-services with 404 error, just add the following to bypass the GOPROXY for private dependencies:
<br> export GOPROXY=direct
<br> export GOPRIVATE=github.com/vickeyshrestha/sharing-services

<h3>Database Credentials</h3>
Refer to docker-compose file. The default username is postgres
<br> You can run stockzilla\services\stock-trader\database\stock-trader.sql against the database, **stockzilla**

<h3>Connect database through Hasura</h3>
go to http://localhost:8080/console
<br> go to Data tab
<br> Select Postgres and click "Connect existing Database"
<br> Give database name as "stockzilla"
<br> Under "Connect database via" option, select "Environment Variable"
<br> Then enter "PG_DATABASE_URL" as a variable and click "Connect"

<h3>Techstacks used</h3>
- Golang
- Docker
- Nats Message bus
- Postgres
- Mongo DB
- Hasura
