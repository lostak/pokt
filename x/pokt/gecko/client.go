package gecko

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTokenPrice(coinId string, denom string) (price float64, err error) {
	const baseURL = "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s"
	url := fmt.Sprintf(baseURL, coinId, denom)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return price, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return price, err
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
		return price, err
	}

	// Unmarshal using a generic interface
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return price, err
	}

	// JSON object parses into a map with string keys
	chainIdMap := f.(map[string]interface{})

	for _, v := range chainIdMap {
		// type assert value is JSON
		switch jsonObj := v.(type) {
		case interface{}:
			for baseDenom, amount := range jsonObj.(map[string]interface{}) {
				switch baseDenom {
				case denom:
					// make sure value is a number
					switch amount := amount.(type) {
					case float64:
						price = float64(amount)
					default:
						return price, fmt.Errorf("Incorrect type for %s", baseDenom)
					}
				default:
					return price, fmt.Errorf("Unkown key: %s", baseDenom)
				}
			}
		default:
			return price, fmt.Errorf("Expected JSON object, got something else")
		}
	}

	return price, nil
}

func PrintAllCoinIds() {
	const url = "https://api.coingecko.com/api/v3/coins/list"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
		return
	}

	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	response, _ := json.MarshalIndent(f, "", "\t")
	fmt.Println(string(response))
}
