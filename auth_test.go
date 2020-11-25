package smartcatclient

import "testing"

func TestConfig_AuthToken(t *testing.T) {
	type fields struct {
		AccountID string
		AuthKey   string
		URL       string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"case 1", fields{AccountID: "", AuthKey: "", URL: ""}, "Basic Og=="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Config{
				AccountID: tt.fields.AccountID,
				AuthKey:   tt.fields.AuthKey,
				URL:       tt.fields.URL,
			}
			if got := c.AuthToken(); got != tt.want {
				t.Errorf("AuthToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
