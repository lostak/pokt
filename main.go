/*
Copyright © 2023 Bradley Lostak lostak@engineer.com
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"fmt"
	"syscall/js"

	"github.com/lostak/pokt/client/cmd"
)

func GetPortfolioWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 1 {
			return map[string]any{
				"error": fmt.Sprintf("named portfolios not yet supported"),
			}
		}

		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			return map[string]any{
				"error": "Unable to get document object",
			}
		}

		outputText := jsDoc.Call("querySelector", "#jsonoutput")
		if !outputText.Truthy() {
			return map[string]any{
				"error": "Unable to get output textarea",
			}
		}

		portfolio, err := cmd.GetPortfolioJSON()
		if err != nil {
			return map[string]any{
				"error": err.Error(),
			}
		}

		outputText.Set("value", portfolio)
		fmt.Printf("Porfolio as JSON: \n%s", portfolio)
		return map[string]any{
			"portfolio": portfolio,
		}
	})

	return jsonFunc
}

func main() {
	js.Global().Set("formatPortfolioJSON", GetPortfolioWrapper())
	cmd.Execute()
}
