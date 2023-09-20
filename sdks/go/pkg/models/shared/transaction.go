// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type Transaction struct {
	ID        int64             `json:"id"`
	Metadata  map[string]string `json:"metadata"`
	Postings  []Posting         `json:"postings"`
	Reference *string           `json:"reference,omitempty"`
	Reverted  bool              `json:"reverted"`
	Timestamp time.Time         `json:"timestamp"`
}

func (o *Transaction) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *Transaction) GetMetadata() map[string]string {
	if o == nil {
		return map[string]string{}
	}
	return o.Metadata
}

func (o *Transaction) GetPostings() []Posting {
	if o == nil {
		return []Posting{}
	}
	return o.Postings
}

func (o *Transaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *Transaction) GetReverted() bool {
	if o == nil {
		return false
	}
	return o.Reverted
}

func (o *Transaction) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}
