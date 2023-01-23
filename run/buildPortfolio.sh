# TO BE CALLED FROM BASE DIR
echo "build a portfolio w/ bash commands"

echo "Serve gRPC" 
bash ./pokt serve &

echo "Create Portfolio"
bash ./pokt createPortfolio PortfolioName

echo "Add Account AccountName to Portfolio"
bash ./pokt addAccount AccountName

echo "Add Chain named Cosmos-Hub to Account named AccountName"
bash ./pokt addChain AccountName Cosmos-Hub Address 

echo "Add Token named ATOM to Cosmos-Hub Chain for Account named AccountName"
bash ./pokt addToken AccountName Cosmos-Hub ATOM

echo "Update ATOM's CoinGecko Id to be cosmos"
bash ./pokt updateGeckoId AccountName Cosmos-Hub ATOM cosmos

echo "Build portfolio using bash "
bash chmod +x run/*.sh 

echo "add 4 accounts"
bash ./run/addAccounts.sh

echo "add 4 chains to each account"
bash ./run/addChains.sh

echo "add 1 token to each chain on each account"
bash ./run/addTokens.sh

echo "add amounts to each token on each chain on each account"
bash ./run/addAmounts.sh

echo "update Coin Gecko Id for the cosmos token Account-0 Chain-0 to be btc"
bash ./pokt updateGeckoId Account-0 Chain-0 cosmos btc

echo "Print $ATOM price"
bash ./pokt tokenPrice cosmos usd


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

echo "Remove Token named ATOM from Cosmos-Hub from Account named AccountName"
bash ./pokt removeToken AccountName Cosmos-Hub ATOM

echo "Remove Chain named Cosmos-Hub from Account named AccountName"
bash ./pokt removeChain AccountName Cosmos-Hub

echo "Remove Account named AccountName"
bash ./pokt removeAccount AccountName
