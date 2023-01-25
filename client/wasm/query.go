package main

import (
	"fmt"
	"syscall/js"

	"github.com/lostak/pokt/client/grpc"
)

func getOutputText() (js.Value, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return js.Null(), fmt.Errorf("Unable to get document object")
	}

	outputText := jsDoc.Call("querySelector", "#jsonoutput")
	if !outputText.Truthy() {
		return js.Null(), fmt.Errorf("Unable to get output textarea")
	}

	return outputText, nil
}

func GetPortfolioWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 1 {
			return map[string]any{
				"error": "named portfolios not yet supported",
			}
		}

		output, err := getOutputText()
		if err != nil {
			return map[string]any{
				"error": err.Error(),
			}
		}

		portfolio, err := grpc.GetPortfolioJSON()
		if err != nil {
			return map[string]any{
				"error": err.Error(),
			}
		}

		output.Set("value", portfolio)
		fmt.Printf("Porfolio as JSON: \n%s", portfolio)
		return map[string]any{
			"portfolio": portfolio,
		}
	})

	return jsonFunc
}

func GetAccountWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return map[string]any{
				"error": "named portfolios not yet supported",
			}
		} else {
			fmt.Println("Using account: ", args[0].String())
		}

		output, err := getOutputText()
		if err != nil {
			return map[string]any{
				"error": err.Error(),
			}
		}

		account, err := grpc.GetAccountJSON(args[0].String())
		if err != nil {
			return map[string]any{
				"error": err.Error(),
			}
		}

		output.Set("value", account)
		fmt.Printf("Account as JSON: \n%s", account)
		return map[string]any{
			"account": account,
		}
	})

	return jsonFunc
}
