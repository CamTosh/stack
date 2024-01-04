package modulr

import (
	"encoding/json"
	"fmt"

	"github.com/formancehq/payments/cmd/connectors/internal/connectors"
	"github.com/formancehq/payments/cmd/connectors/internal/connectors/configtemplate"
)

type Config struct {
	Name          string              `json:"name" bson:"name"`
	APIKey        string              `json:"apiKey" bson:"apiKey"`
	APISecret     string              `json:"apiSecret" bson:"apiSecret"`
	Endpoint      string              `json:"endpoint" bson:"endpoint"`
	PollingPeriod connectors.Duration `json:"pollingPeriod" yaml:"pollingPeriod" bson:"pollingPeriod"`
	PageSize      int                 `json:"pageSize" yaml:"pageSize" bson:"pageSize"`
}

// String obfuscates sensitive fields and returns a string representation of the config.
// This is used for logging.
func (c Config) String() string {
	return fmt.Sprintf("endpoint=%s, apiSecret=***, apiKey=****", c.Endpoint)
}

func (c Config) Validate() error {
	if c.APIKey == "" {
		return ErrMissingAPIKey
	}

	if c.APISecret == "" {
		return ErrMissingAPISecret
	}

	if c.Name == "" {
		return ErrMissingName
	}

	return nil
}

func (c Config) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c Config) ConnectorName() string {
	return c.Name
}

func (c Config) BuildTemplate() (string, configtemplate.Config) {
	cfg := configtemplate.NewConfig()

	cfg.AddParameter("name", configtemplate.TypeString, true)
	cfg.AddParameter("apiKey", configtemplate.TypeString, true)
	cfg.AddParameter("apiSecret", configtemplate.TypeString, true)
	cfg.AddParameter("endpoint", configtemplate.TypeString, false)
	cfg.AddParameter("pollingPeriod", configtemplate.TypeDurationNs, false)
	cfg.AddParameter("pageSize", configtemplate.TypeDurationUnsignedInteger, false)

	return Name.String(), cfg
}
