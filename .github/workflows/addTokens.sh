# MUST CALL FROM BASE DIR FOR PATHS
echo -e "\nAdd cosmos token to each chain on Account-0\n"
./pokt addToken Account-0 Chain-0 cosmos 
echo -e ""
./pokt addToken Account-0 Chain-1 cosmos 

echo -e "\nAdd cosmos token to each chain on Account-0\n"
./pokt addToken Account-1 Chain-0 cosmos 
echo -e ""
./pokt addToken Account-1 Chain-1 cosmos 
