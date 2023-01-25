package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"syscall/js"
	"time"

	"github.com/lostak/pokt/client/grpc"
	"github.com/lostak/pokt/server"
	g "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	addr = flag.String("addr", "localhost:8080", "the address to connect to for grpc")
)

func getOutputText() (js.Value, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return js.Null(), fmt.Errorf("unable to get document object")
	}

	outputText := jsDoc.Call("querySelector", "#b")
	if !outputText.Truthy() {
		return js.Null(), fmt.Errorf("unable to get output textarea")
	}

	return outputText, nil
}

//exports getPortfolioWrapper
func GetPortfolioWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("GetPortfolioWrapper called")

		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			fmt.Println("unable to get ddocument")

			return map[string]any{
				"error": fmt.Errorf("unable to get document object"),
			}
		}

		outputText := jsDoc.Call("querySelector", "#b")
		if !outputText.Truthy() {
			fmt.Println("unable to get output textarea")

			return map[string]any{
				"error": fmt.Errorf("unable to get output textarea"),
			}
		}

		conn, err := g.Dial(*addr, g.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			fmt.Println("unable to dial")

			return map[string]any{
				"error": err.Error(),
			}
		}

		defer conn.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		client := server.NewQueryServiceClient(conn)

		response, err := client.GetPortfolio(ctx, &server.MsgGetPortfolio{})
		if err != nil {
			fmt.Println("error w/ getting output text ", err.Error())

			return map[string]any{
				"error": err.Error(),
			}
		}

		portfolio := response.GetPortfolio()
		if portfolio == nil {
			fmt.Println("portfolio recieved is nil")

			return map[string]any{
				"error": "portfolio recieved is nil",
			}
		}

		fmt.Println("Portfolio Recieved:")
		portfolio.Println()

		b, err := protojson.Marshal(portfolio)
		if err != nil {
			fmt.Println("error marshaling portfolio")

			return map[string]any{
				"error": err.Error(),
			}
		}

		var ugly any
		err = json.Unmarshal(b, &ugly)
		if err != nil {
			fmt.Println("unable to unmarshal bytes")

			return map[string]any{
				"error": err.Error(),
			}
		}

		pretty, err := json.MarshalIndent(ugly, "", "  ")
		if err != nil {
			fmt.Println("unable to marshalIndent")

			return map[string]any{
				"error": err.Error(),
			}
		}

		outputText.Set("value", pretty)
		fmt.Printf("Porfolio as JSON: \n%s", pretty)
		return map[string]any{
			"portfolio": pretty,
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

func main() {
	wait := make(chan struct{}, 0)

	fmt.Println("Go web assembly")

	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		fmt.Println("Wasm failed - could not get document")
		return
	}

	button := jsDoc.Call("querySelector", "#a")
	if !button.Truthy() {
		fmt.Println("wasm failed - could not get button #a")
		return
	}

	button.Call("addEventListener", "click", GetPortfolioWrapper())

	<-wait
}
