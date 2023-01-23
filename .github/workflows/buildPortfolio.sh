# MUST CALL FROM BASE DIR FOR PATHS
echo -e "\nbuild a portfolio w/ bash commands\n"

echo -e "\nRun gRPC server\n"
./pokt serve &

echo -e "\nPrint ATOM price"
./pokt tokenPrice cosmos usd

echo -e "\n Run demo\n"
bash ./.github/workflows/demo.sh

echo -e "\ncreate a new portfolio\n"
./pokt createPortfolio BrandNew