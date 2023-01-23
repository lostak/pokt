# MUST CALL FROM BASE DIR FOR PATHS
echo "Add 4 chains to Account-0"
./pokt addChain Account-0 Chain-0 addr
./pokt addChain Account-0 Chain-1 addr

echo "Add 4 chains to Account-1"
./pokt addChain Account-1 Chain-0 addr
./pokt addChain Account-1 Chain-1 addr