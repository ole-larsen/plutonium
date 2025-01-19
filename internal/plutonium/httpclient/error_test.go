package httpclient_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ole-larsen/plutonium/internal/plutonium/httpclient"
	"github.com/stretchr/testify/require"
)

func TestNewHTTPError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		args    args
		name    string
		wantErr bool
	}{
		{
			name: "test err",
			args: args{
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "test error1",
			args: args{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
		{
			name: "test  error2",
			args: args{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := httpclient.NewHTTPError(tt.args.err, 0)
			if tt.args.err == nil {
				require.Nil(t, err)
			} else {
				if (err != nil) != tt.wantErr {
					t.Errorf("NewHTTPError() error = %v, wantErr %v", err, tt.wantErr)
				}

				require.Equal(t, "*httpclient.HTTPError", fmt.Sprintf("%T", err))

				require.Equal(t, fmt.Sprintf("%v %s", tt.args.err, time.Now().Format("2006/01/02 15:04:05")), err.Error())
			}
		})
	}
}

func TestNewError(t *testing.T) {
	type args struct {
		err error
	}

	tests := []struct {
		args    args
		name    string
		wantErr bool
	}{
		{
			name: "test err",
			args: args{
				err: nil,
			},
			wantErr: true,
		},
		{
			name: "test error1",
			args: args{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
		{
			name: "test  error2",
			args: args{
				err: errors.New("some error"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := httpclient.NewError(tt.args.err)
			if tt.args.err == nil {
				require.Nil(t, err)
			} else {
				if (err != nil) != tt.wantErr {
					t.Errorf("NewError() error = %v, wantErr %v", err, tt.wantErr)
				}

				require.Equal(t, "*httpclient.Error", fmt.Sprintf("%T", err))

				require.Equal(t, fmt.Sprintf("[client]: %v", tt.args.err), err.Error())
			}
		})
	}
}
