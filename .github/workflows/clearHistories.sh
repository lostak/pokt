
echo -e "\nDelete the token history for ATOM\n"
./pokt deleteTokenHistory AccountName Cosmos-Hub ATOM 

echo -e "\nAdd a new token to Cosmos-Hub and then clear all the chain tokens' histories\n"
./pokt addToken AccountName Cosmos-Hub USDC 
echo -e ""
./pokt deleteChainHistory AccountName Cosmos-Hub 

echo -e "\nClear the Cosmos-Hub history\n"
./pokt deleteChainHistory AccountName Cosmos-Hub

echo -e "\nClear AccountName History\n"
./pokt deleteAccountHistory AccountName 

echo -e "\nClear Portfolio History\n"
./pokt deletePortfolioHistory 

echo -e "\nDelete the token history for ATOM\n"
./pokt deleteTokenHistory AccountName Cosmos-Hub ATOM 

echo -e "\nAdd a new amount and then clear the chain token history\n"
./pokt addAmount AccountName Cosmos-Hub ATOM 10 
echo -e ""
./pokt deleteChainHistory AccountName Cosmos-Hub 
