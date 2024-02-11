#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' 
for ((i=1;i<100;i++))
do
math=$(./linear-stats)
main=$(go run main.go data.txt) 
echo -e "---linear-stats---"
echo -e "$math\n"
echo -e "-----main.go-----"
echo -e "$main\n"
if [[ $math == $main ]]
  then
    printf "${GREEN}------EQUAL------\n\n${NC}"
    else 
    printf "${RED}----NOT-EQUAL----\n\n${NC}"
    break
fi
done

