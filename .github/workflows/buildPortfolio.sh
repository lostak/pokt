# MUST CALL FROM BASE DIR FOR PATHS
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

echo "cur dir"
ls .

echo "add 4 accounts"
bash ./addAccounts.sh

echo "add 4 chains to each account"
bash ./addChains.sh

echo "add 1 token to each chain on each account"
bash ./addTokens.sh

echo "add amounts to each token on each chain on each account"
bash ./addAmounts.sh

echo "clear histories"
bash ./clearHistories.sh

echo "update CoinGecko Id"
bash ./updateGeckoId.sh 

echo "remove tokens, chains and accounts"
bash ./removal.sh

echo "create a new portfolio"
./pokt createPortfolio BrandNew