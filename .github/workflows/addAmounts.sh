# MUST CALL FROM BASE DIR FOR PATHS
echo -e "\nAdd 2 amounts to cosmos on each chain for Account-0 \n"
./pokt addAmount Account-0 Chain-0 cosmos 0  
sleep 1
./pokt addAmount Account-0 Chain-1 cosmos 1  
sleep 1

echo -e "\nAdd 2 amounts to cosmos on each chain for Account-1\n"
./pokt addAmount Account-1 Chain-0 cosmos 0  
sleep 1
./pokt addAmount Account-1 Chain-1 cosmos 1  