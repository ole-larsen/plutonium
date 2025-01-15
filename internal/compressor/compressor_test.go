package compressor

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressDecompress(t *testing.T) {
	type args struct {
		data []byte
	}

	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "compress / decompress positive 1",
			args: args{
				data: []byte(strings.Repeat("Hello World", 10)),
			},
			want:    []byte(strings.Repeat("Hello World", 10)),
			wantErr: false,
		},
		{
			name: "compress / decompress positive 1",
			args: args{
				data: []byte(strings.Repeat("Hello World", 100)),
			},
			want:    []byte(strings.Repeat("Hello World", 100)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := Compress(tt.args.data)
			require.NoError(t, err, "must compress")
			require.Greater(t, len(tt.args.data), len(b))
			t.Logf("%d bytes has been compressed to %d bytes\r\n", len(tt.args.data), len(b))

			out, err := Decompress(bytes.NewBuffer(b))
			require.NoError(t, err, "must decompress")

			if !bytes.Equal(tt.args.data, out) {
				t.Errorf(`original data %d != decompressed data %d`, len(tt.args.data), len(out))
			}
		})
	}
}
