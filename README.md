# pokt

## Approach
1. Portfolio and children CRUD for CLI
    - Process for adding new messages w/o adding methods to the type:
        1. Write proto rpc message(s) w/ format: `rpc ActionType(MsgActionType) returns (MsgActionTypeResponse);`
        2. Write proto rpc messages for `MsgActionType` and `MsgActionTypeResponse`
        3. Generate `.pb.go` files
        4. Write `ActionType(MsgActionType)`  
        5. Write cobra command to call `ActionType` as a gRPC client 
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
