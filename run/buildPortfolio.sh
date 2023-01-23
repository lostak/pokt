# TO BE CALLED FROM BASE DIR
echo "build a portfolio w/ bash commands"

echo "Run gRPC server"
./pokt serve &

echo "Print $ATOM price"
./pokt tokenPrice cosmos usd

echo "Create Portfolio"
./pokt createPortfolio PortfolioName

echo "Add Account AccountName to Portfolio"
./pokt addAccount AccountName

echo "Add Chain named Cosmos-Hub to Account named AccountName"
./pokt addChain AccountName Cosmos-Hub Address 

echo "Add Token named ATOM to Cosmos-Hub Chain for Account named AccountName"
./pokt addToken AccountName Cosmos-Hub ATOM

echo "Update ATOM's CoinGecko Id to be cosmos"
./pokt updateGeckoId AccountName Cosmos-Hub ATOM cosmos

echo "add 4 accounts"
bash ./run/addAccounts.sh

echo "add 4 chains to each account"
bash ./run/addChains.sh

echo "add 1 token to each chain on each account"
bash ./run/addTokens.sh

echo "add amounts to each token on each chain on each account"
bash ./run/addAmounts.sh

echo "clear histories"
bash ./run/clearHistories.sh

echo "update CoinGecko Id"
bash ./run/updateGeckoId.sh 

echo "remove tokens, chains and accounts"
bash ./run/removal.sh

echo "create a new portfolio"
./pokt createPortfolio BrandNew