// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type LedgerAccountSubject struct {
	Identifier string `json:"identifier"`
	Type       string `json:"type"`
}

func (o *LedgerAccountSubject) GetIdentifier() string {
	if o == nil {
		return ""
	}
	return o.Identifier
}

func (o *LedgerAccountSubject) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}