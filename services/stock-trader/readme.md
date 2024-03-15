**Service - stock trader**

Service stock trader will subscribe to a specific topics in NATS (topics published by simulator - stock by default) to get NASDAQ Stock Exchange, NY Stock Exchange, and London Stock Exchange data in a timeseries fashion.

**architecture diagram**
<br>
![architecture](readme-resources/architecture.jpg)


**Database requirement**
<br>database named "stockzilla" should be taken care by docker-compose

**building and pushing docker image**
<br> Run the builddocker.bat file by passing the version as an argument. For eg:
<br> C:\stockzilla\simulators\stock> builddocker.bat 0.0.1
<br> Test by running docker run first before pushing into dockerhub
<br> Use **docker push image:ver** to push into dockerhub


**Service endpoints**

| EndPoint Syntax           |                    Example                    |                                    Detail |
|---------------------------|:---------------------------------------------:|------------------------------------------:|
| /api/stock_trader/status  | http://localhost:8085/api/stock_trader/status | Get the health status of this application |
| /api/stock_trader/actions | http://localhost:8085/api/stock_trader/status |  Get the most recent record from postgres |
