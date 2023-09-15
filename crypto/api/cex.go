package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"qmmr.xyz/go/crypto/types"
)

const apiUrl = "https://cex.io/api/ticker/%v/USD"

func GetRate(currency string) (*types.Rate, error) {
	upperCurrency := strings.ToUpper(currency)

	resp, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}

		// fmt.Println(response)

	} else {
		return nil, fmt.Errorf("status code was %v", resp.StatusCode)
	}

	rate := types.Rate{
		Currency: currency,
		Price:    response.Bid,
	}

	return &rate, nil
}
