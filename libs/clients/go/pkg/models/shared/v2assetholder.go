// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
)

type V2AssetHolder struct {
	Assets map[string]*big.Int `json:"assets"`
}

func (v V2AssetHolder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2AssetHolder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2AssetHolder) GetAssets() map[string]*big.Int {
	if o == nil {
		return map[string]*big.Int{}
	}
	return o.Assets
}