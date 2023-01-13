package types

import (
	"encoding/json"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

const (
	FileName = "portfolio.db"
	filePath = "../"
)

func GetPortfolio() (*Portfolio, error) {
	data, err := os.ReadFile(filePath + FileName)
	if err != nil {
		fmt.Print(err.Error())
		return &Portfolio{}, err
	}

	portfolio := &Portfolio{}

	err = json.Unmarshal(data, &portfolio)
	if err != nil {
		fmt.Print(err.Error())
		return &Portfolio{}, err
	}

	return portfolio, nil
}

func SetPortfolio(portfolio *Portfolio) error {
	b, err := proto.Marshal(portfolio)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	if err := os.WriteFile(filePath+FileName, b, 0700); err != nil {
		fmt.Print(err.Error())
	}

}
