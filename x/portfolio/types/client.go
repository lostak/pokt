package types

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetTokenPrice(coinId string, denom string) (float64, error) {
	const baseURL = "https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=%s"
	url := fmt.Sprintf(baseURL, coinId, denom)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
		return 0, err
	}

	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		fmt.Print(err.Error())
		return 0, err
	}

	// Unmarshal using a generic interface
	var f interface{}
	err = json.Unmarshal(body, &f)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return 0, err
	}

	// JSON object parses into a map with string keys
	chainIdMap := f.(map[string]interface{})
	var price float64

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
						fmt.Println("Incorrect type for ", baseDenom)
					}
				default:
					fmt.Println("Unkown key")
				}
			}
		default:
			fmt.Println("Expected JSON object, got something else")
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
