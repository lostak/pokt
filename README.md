# pokt


## Architecture
     __________________              ______________________                ___________________
    |                  |   (http)   |                      | (Read/Write) |                   |
    |   CoinGecko API  | <--------> |   Portfolio Module   | <----------> |   Local Storage   |
    |__________________|            |______________________|              |___________________|
     ______________                       ^                                         ^     
    |              |      (CLI)           |                                         |
    |   Terminal   | <--------------------                                          |
    |______________|                                                                |
     ______________          __________                                             |
    |              |        |          |                                            |  
    |    Browser   | <----> |   Node   | <------------------------------------------
    |______________|  (JS)  |__________|                 (Read/Write)