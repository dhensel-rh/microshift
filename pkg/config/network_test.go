package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCNIPlugin_IsEnabled(t *testing.T) {
	type fields struct {
		CNIPlugin CNIPlugin
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "is disabled when plugin is none",
			fields: fields{
				CNIPlugin: CniPluginNone,
			},
			want: false,
		},
		{
			name: "is enabled when plugin is ovnk",
			fields: fields{
				CNIPlugin: CniPluginOVNK,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Network{
				CNIPlugin: tt.fields.CNIPlugin,
			}
			assert.Equalf(t, tt.want, n.IsEnabled(), "IsEnabled()")
		})
	}
}

func TestNetwork_cniPluginIsValid(t *testing.T) {
	type fields struct {
		CNIPlugin CNIPlugin
	}
	tests := []struct {
		name            string
		fields          fields
		wantIsSupported bool
	}{
		{
			name: "is valid when value matches one of predefined drivers",
			fields: fields{
				CNIPlugin: CniPluginOVNK,
			},
			wantIsSupported: true,
		},
		{
			name: "is valid when value is an empty string",
			fields: fields{
				CNIPlugin: "",
			},
			wantIsSupported: true,
		},
		{
			name: "is invalid when value does not match one of predefined drivers",
			fields: fields{
				CNIPlugin: "foobar",
			},
			wantIsSupported: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Network{
				CNIPlugin: tt.fields.CNIPlugin,
			}
			assert.Equalf(t, tt.wantIsSupported, n.validCNIPlugin(), "validCNIPlugin()")
		})
	}
}
