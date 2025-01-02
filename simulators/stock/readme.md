**1. simulator - stock**
<p> 
I have created this simulator to mimic what the service using Alpha Vantage Support should have done. Since it requires a premium membership to use its full potential of providing unlimited stock's information, I am temporarily opting out from that API service and using this simulator instead.
</p>

The goal of this simulator is very simple:
* we have 5 cross-listing companies appearing in 3 different stock exchanges.
* Theoretically, prices will be very similar across different exchanges,
  but slight variations can occur due to factors like exchange fees,
  time zone differences, and varying levels of liquidity on each exchange,
  meaning the price might appear slightly different depending on where you
  look at a given moment.
* Hence, this simulator will publish some random numbers, assuming they are 
  the current prices of these 5 different companies across 3 different stock exchanges 
  into messagebus (NATS) to be consumed by other client services like "stock-trader".


**2. building and pushing docker image**
<br> Run the builddocker.bat file by passing the version as an argument. For eg:
<br> C:\...stockzilla\simulators\stock> builddocker.bat 0.0.1
<br> Test by running docker run first before pushing into dockerhub
<br> Use **docker push image:ver** to push into dockerhub