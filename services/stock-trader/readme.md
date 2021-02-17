**service - stock trader**

Service stock trader will subscribe to a specific topics in NATS (topics published by simulator - stock by default) to get NASDAQ Stock Exchange, NY Stock Exchange, and London Stock Exchange data in a timeseries fashion.

**building and pushing docker image**
<br> Run the builddocker.bat file by passing the version as an argument. For eg:
<br> C:\godzilla\simulators\stock> builddocker.bat 0.0.1
<br> Test by running docker run first before pushing into dockerhub
<br> Use **docker push image:ver** to push into dockerhub