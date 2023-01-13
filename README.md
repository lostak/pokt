# pokt

## Approach
1. Portfolio and children CRUD for CLI
2. Portfolio updates w/ API using CLI
3. Browser charts from local storage
4. Browser CRUD with portofio  
## Architecture
~~~
 __________________          ______________________                         ___________________
|                  | (HTTP) |                      | (Read/Write Protobuf) |                   |
|   CoinGecko API  | <----> |   Portfolio Module   | <-------------------> |   Local Storage   |
|__________________|        |______________________|                       |___________________|
 ______________                    ^        ^                                        ^     
|              |      (Cobra CLI)  |        |                                        |
|   Terminal   | <-----------------+        |  (gRPC)                                |
|______________|                            |                                        |
 ______________          __________         |                                        |
|              |  (JS)  |          | <------+                                        |  
|    Browser   | <----> |   Node   |                 (Read/Write Protobuf)           |
|______________|        |__________| <-----------------------------------------------+ 
~~~