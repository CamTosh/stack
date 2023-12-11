// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PoolRequest struct {
	AccountIDs []string `json:"accountIDs"`
	Name       string   `json:"name"`
}

func (o *PoolRequest) GetAccountIDs() []string {
	if o == nil {
		return []string{}
	}
	return o.AccountIDs
}

func (o *PoolRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}