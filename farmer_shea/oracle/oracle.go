package oracle

import (
	"context"
	"fmt"
	"math"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	bin "github.com/gagliardetto/binary"
)

const (
	pythMappingAccount = "AHtgzX45WTKfkPG53L6WYhGEXwQkN1BVknET3sVsLL8J"
)

// Oracle defines the interface for a price oracle.	ype Oracle interface {
	GetPrice(token string) (float64, error)
}

// Pyth is a price oracle that uses the Pyth network.	ype Pyth struct {
	client *rpc.Client
}

// NewPyth creates a new Pyth oracle.
func NewPyth(client *rpc.Client) *Pyth {
	return &Pyth{client: client}
}

// GetPrice returns the price of a given token.
func (p *Pyth) GetPrice(token string) (float64, error) {
	mappingAccount, err := solana.PublicKeyFromBase58(pythMappingAccount)
	if err != nil {
		return 0, err
	}

	info, err := p.client.GetAccountInfo(context.Background(), mappingAccount)
	if err != nil {
		return 0, err
	}

	priceAccountKey, err := p.findProductAccount(info.Value.Data, token)
	if err != nil {
		return 0, err
	}

	priceInfo, err := p.client.GetAccountInfo(context.Background(), *priceAccountKey)
	if err != nil {
		return 0, err
	}

	var priceAccount PythPriceAccount
	if err := bin.NewBinDecoder(priceInfo.Value.Data).Decode(&priceAccount); err != nil {
		return 0, err
	}

	return float64(priceAccount.Agg.Price) * math.Pow10(int(priceAccount.Exponent)), nil
}

func (p *Pyth) findProductAccount(data []byte, token string) (*solana.PublicKey, error) {
	var mapping PythMappingAccount
	if err := bin.NewBinDecoder(data).Decode(&mapping); err != nil {
		return nil, err
	}

	for _, productKey := range mapping.Products {
		info, err := p.client.GetAccountInfo(context.Background(), productKey)
		if err != nil {
			continue
		}

		var product PythProductAccount
		if err := bin.NewBinDecoder(info.Value.Data).Decode(&product); err != nil {
			continue
		}

		if product.Attrs["symbol"] == token {
			return &product.Price, nil
		}
	}

	return nil, fmt.Errorf("product not found for token %s", token)
}

// GetJupiterPrices returns the prices of the given tokens from the Jupiter API.
func GetJupiterPrices(mints []string) (map[string]float64, error) {
	baseURL := "https://price.jup.ag/v4/price?ids="
	ids := ""
	for i, mint := range mints {
		ids += mint
		if i < len(mints)-1 {
			ids += ","
		}
	}

	resp, err := http.Get(baseURL + ids)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data map[string]struct {
			Price float64 `json:"price"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for mint, data := range result.Data {
		prices[mint] = data.Price
	}

	return prices, nil
}

// PythMappingAccount is the structure of a Pyth mapping account.	ype PythMappingAccount struct {
	Magic       uint32
	Version     uint32
	Type        uint32
	Size        uint32
	Num         uint32
	Pad         [4]byte
	Products    [2048]solana.PublicKey
}

// PythProductAccount is the structure of a Pyth product account.	ype PythProductAccount struct {
	Magic       uint32
	Version     uint32
	Type        uint32
	Size        uint32
	Price       solana.PublicKey
	Attrs       map[string]string
}

// PythPriceAccount is the structure of a Pyth price account.	ype PythPriceAccount struct {
	Magic       uint32
	Version     uint32
	Type        uint32
	Size        uint32
	PriceType   uint32
	Exponent    int32
	Num         uint32
	NumQt       uint32
	LastSlot    uint64
	ValidSlot   uint64
	Twap        int64
	Twac        int64
	Drv         [32]byte
	Agg         PythPriceInfo
	Comps       [32]PythPriceComp
}

// PythPriceInfo is the structure of the price and confidence interval.	ype PythPriceInfo struct {
	Price       int64
	Conf        uint64
	Status      uint32
	CorpAct     uint32
	PubSlot     uint64
}

// PythPriceComp is the structure of a component price.	ype PythPriceComp struct {
	Publisher   solana.PublicKey
	Agg         PythPriceInfo
	Latest      PythPriceInfo
}
