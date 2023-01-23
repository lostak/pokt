# MUST CALL FROM BASE DIR FOR PATHS
echo "Add cosmos token to each chain on Account-0"
./pokt addToken Account-0 Chain-0 cosmos 
./pokt addToken Account-0 Chain-1 cosmos 

echo "Add cosmos token to each chain on Account-0"
./pokt addToken Account-1 Chain-0 cosmos 
./pokt addToken Account-1 Chain-1 cosmos 
