// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2StageWaitEvent struct {
	Event string `json:"event"`
}

func (o *V2StageWaitEvent) GetEvent() string {
	if o == nil {
		return ""
	}
	return o.Event
}