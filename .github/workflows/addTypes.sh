echo -e "\nadd 4 accounts\n"
bash ./.github/workflows/addAccounts.sh

echo -e "\nadd 4 chains to each account\n"
bash ./.github/workflows/addChains.sh

echo -e "\nadd 1 token to each chain on each account\n"
bash ./.github/workflows/addTokens.sh

echo -e "\nadd amounts to each token on each chain on each account\n"
bash ./.github/workflows/addAmounts.sh
