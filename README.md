# pokt

## Approach
1. Portfolio and children CRUD for CLI
2. Portfolio updates w/ API using CLI
3. Browser charts from local storage
4. Browser CRUD with portofio  
## Architecture
~~~
 __________________            ______________________                         ___________________
|                  |  (HTTP)  |                      | (Read/Write Protobuf) |                   |
|   CoinGecko API  | <------> |   Portfolio Module   | <-------------------> |   Local Storage   |
|__________________|          |      _____________   |                       |___________________|
 ______________               |     |             |  |                                 ^     
|              | (Cobra CLI)  |     | gRPC Server |  |                                 |
|   Terminal   | <----------------->|_____________|  |                 (Read Protobuf) |
|______________|              |_____________^________|                                 |
 ______________          __________         |                                          |
|              |  (JS)  |          |        | (Proto Request/Response)                 |  
|    Browser   | <----> |   Node   | <------+                                          |
|______________|        |__________| <-------------------------------------------------+ 
~~~
