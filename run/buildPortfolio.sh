# TO BE CALLED FROM BASE DIR
echo "build a portfolio w/ bash commands"
echo "add 4 accounts"
bash .run/addAccounts.sh
echo "add 4 chains to each account"
bash .run/addChains.sh
echo "add 1 token to each chain on each account"
bash .run/addTokens.sh
echo "add amounts to each token on each chain on each account"
bash .run/addAmounts.sh