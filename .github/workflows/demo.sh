echo -e "\nCreate Portfolio\n"
./pokt createPortfolio PortfolioName

echo -e "\nAdd Account AccountName to Portfolio\n"
./pokt addAccount AccountName

echo -e "\nAdd Chain named Cosmos-Hub to Account named AccountName\n"
./pokt addChain AccountName Cosmos-Hub Address 

echo -e "\nAdd Token named ATOM to Cosmos-Hub Chain for Account named AccountName\n"
./pokt addToken AccountName Cosmos-Hub ATOM

echo -e "\nUpdate ATOM's CoinGecko Id to be cosmos\n"
./pokt updateGeckoId AccountName Cosmos-Hub ATOM cosmos

echo -e "\nadd types"
bash ./.github/workflows/addTypes.sh

echo -e "\nclear histories\n"
bash ./.github/workflows/clearHistories.sh

echo -e "\nupdate CoinGecko Id\n"
bash ./.github/workflows/updateGeckoId.sh 

echo -e "\nremove tokens, chains and accounts\n"
bash ./.github/workflows/removal.sh