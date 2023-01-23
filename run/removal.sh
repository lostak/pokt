echo "Remove Token named ATOM from Cosmos-Hub from Account named AccountName"
bash ./pokt removeToken AccountName Cosmos-Hub ATOM

echo "Remove Chain named Cosmos-Hub from Account named AccountName"
bash ./pokt removeChain AccountName Cosmos-Hub

echo "Remove Account named AccountName"
bash ./pokt removeAccount AccountName
