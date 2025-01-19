package httpclient_test

import (
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/ole-larsen/plutonium/internal/plutonium/httpclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewHttpClient(t *testing.T) {
	type args struct {
		transport         http.RoundTripper
		checkRedirectFunc func(req *http.Request, via []*http.Request) error
	}

	transport, ok := http.DefaultTransport.(*http.Transport)
	require.True(t, ok)

	tests := []struct {
		args args
		want *httpclient.HTTPClient
		name string
	}{
		{
			name: "test default http client",
			args: args{
				transport: transport.Clone(),
				checkRedirectFunc: func(req *http.Request, _ []*http.Request) error {
					t.Log(req.URL)
					return nil
				},
			},
			want: &httpclient.HTTPClient{
				Client: &http.Client{
					Timeout: time.Minute,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := httpclient.NewHTTPClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHttpClient() = %v, want %v", got, tt.want)
			}

			c := httpclient.NewHTTPClient()

			assert.Nil(t, c.Client.Transport)
			assert.Nil(t, c.Client.CheckRedirect)

			timeout := 1 * time.Second

			c.SetTimeout(1 * time.Second)

			require.Equal(t, timeout, c.Client.Timeout)
			c.SetTransport(tt.args.transport)

			require.Equal(t, tt.args.transport, c.Client.Transport)
			require.Equal(t, tt.args.transport, c.GetTransport())

			resp, err := c.Get("http://127.0.0.1:8080/status")

			if err != nil {
				return
			}

			require.NoError(t, err, "must be no error")

			// io.Discard выступает в качестве приёмника ненужных данных
			_, err = io.Copy(io.Discard, resp.Body)

			require.NoError(t, err, "must be no error")

			defer func() {
				e := resp.Body.Close()
				require.NoError(t, e)
			}()

			require.Equal(t, http.StatusOK, resp.StatusCode)
			require.NoError(t, err, "must be no error")

			_, err = io.ReadAll(resp.Body)

			require.NoError(t, err, "must be no error")
		})
	}
}
