# pokt


## Architecture
     __________________              ______________________                ___________________
    |                  |   (HTTP)   |                      | (Read/Write) |                   |
    |   CoinGecko API  | <--------> |   Portfolio Module   | <----------> |   Local Storage   |
    |__________________|            |______________________|              |___________________|
     ______________                       ^         ^                                ^     
    |              |      (Cobra CLI)     |         |                                |
    |   Terminal   | <--------------------+         |  (gRPC)                        |
    |______________|                                |                                |
     ______________          __________             |                                |
    |              |        |          | <----------+                                |  
    |    Browser   | <----> |   Node   |                                             |
    |______________|  (JS)  |__________| <-------------------------------------------+ 
                    (Read/Write)                          (Read/Write)