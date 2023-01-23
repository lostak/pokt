
echo "Delete the token history for ATOM"
bash ./pokt deleteTokenHistory AccountName Cosmos-Hub ATOM 

echo "Add a new token to Cosmos-Hub and then clear all the chain tokens' histories"
bash ./pokt addToken AccountName Cosmos-Hub USDC && ./pokt deleteChainHistory AccountName Cosmos-Hub 

echo "Clear the Cosmos-Hub history"
bash ./pokt deleteChainHistory AccountName Cosmos-Hub

echo "Clear AccountName History"
bash ./pokt deleteAccountHistory AccountName 

echo "Clear Portfolio History"
bash ./pokt deletePortfolioHistory 

echo "Delete the token history for ATOM"
bash ./pokt deleteTokenHistory AccountName Cosmos-Hub ATOM 

echo "Add a new amount and then clear the chain token history"
bash ./pokt addAmount AccountName Cosmos-Hub ATOM 10 && ./pokt deleteChainHistory AccountName Cosmos-Hub 
