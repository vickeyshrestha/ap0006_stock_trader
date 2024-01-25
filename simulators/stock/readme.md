**1. simulator - stock**
<p> I have created this simulator to mimic what the service using Alpha Vantage Support should have done. Since it requires a premium membership to use its full potential of providing unlimited stock's information, I am temporarily opting out from that API service and using this simulator instead.
Simulator stock will publish random numbers for each company listed under NASDAQ Stock Exchange, NY Stock Exchange, and London Stock Exchange into messagebus (NATS) to be consumed by other client services like "stock-trader". </p>

**2. building and pushing docker image**
<br> Run the builddocker.bat file by passing the version as an argument. For eg:
<br> C:\...stockzilla\simulators\stock> builddocker.bat 0.0.1
<br> Test by running docker run first before pushing into dockerhub
<br> Use **docker push image:ver** to push into dockerhub