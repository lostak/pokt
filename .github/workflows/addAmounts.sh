# MUST CALL FROM BASE DIR FOR PATHS
echo "Add 4 amounts to cosmos on each chain for Account-0"
./pokt addAmount Account-0 Chain-0 cosmos 0  
./pokt addAmount Account-0 Chain-1 cosmos 1  

echo "Add 4 amounts to cosmos on each chain for Account-1"
./pokt addAmount Account-1 Chain-0 cosmos 0  
./pokt addAmount Account-1 Chain-1 cosmos 1  

echo "Add 4 amounts to cosmos on each chain for Account-2"
./pokt addAmount Account-2 Chain-0 cosmos 0  
./pokt addAmount Account-2 Chain-1 cosmos 1  

echo "Add 4 amounts to cosmos on each chain for Account-3"
./pokt addAmount Account-3 Chain-0 cosmos 0  
./pokt addAmount Account-3 Chain-1 cosmos 1  
