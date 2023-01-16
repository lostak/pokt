package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

const (
	filePath = "poktdb/"
	FileName = "portfolio"
	FileExt  = ".db"
)

func GetPortfolio() (*Portfolio, error) {
	path := filePath + FileName + FileExt
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err.Error())
		return &Portfolio{}, err
	}

	portfolio := &Portfolio{}

	if err := proto.Unmarshal(data, portfolio); err != nil {
		fmt.Print(err.Error())
		return &Portfolio{}, err
	}

	return portfolio, nil
}

func SetPortfolio(portfolio *Portfolio) error {
	b, err := proto.Marshal(portfolio)
	if err != nil {
		return err
	}

	path := filePath + FileName + FileExt
	// check for .poktdb dir
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(filePath, 0755)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

func GetPortfolioFromId(id uint32) (*Portfolio, error) {
	path := fmt.Sprintf("%s%s%d%s", filePath, FileName, id, FileExt)

	data, err := os.ReadFile(path)
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

func SetPortfolioWithId(id uint32, portfolio *Portfolio) error {
	b, err := proto.Marshal(portfolio)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}

	path := fmt.Sprintf("%s%s%d%s", filePath, FileName, id, FileExt)

	if err := os.WriteFile(path, b, 0700); err != nil {
		fmt.Print(err.Error())
		return err
	}

	return nil
}
