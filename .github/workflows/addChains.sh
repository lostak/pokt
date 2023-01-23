# MUST CALL FROM BASE DIR FOR PATHS
echo -e "\nAdd 2 chains to Account-0\n"
./pokt addChain Account-0 Chain-0 addr
./pokt addChain Account-0 Chain-1 addr

echo -e "\nAdd 2 chains to Account-1\n"
./pokt addChain Account-1 Chain-0 addr
./pokt addChain Account-1 Chain-1 addr