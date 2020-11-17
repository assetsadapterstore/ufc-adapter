package ufc

import "testing"

func TestDecrypt(t *testing.T) {
	type args struct {
		msg     string
		fromPub string
		toPub   string
		wif     string
		nonce   uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			want: "hello boy",
			args: args{
				msg:     "17d0ac3874548d7c4ef56236698d719e",
				nonce:   5577006791947779410,
				fromPub: "UFC6fpcoYK72BxsYYRwcBEPGuVoGhy2Yki2YCNfnCZCYxL5xp56Hh",
				toPub:   "UFC6icdz8dWibXRz8PcDn9RMupFkPbwHQ4toHxP8UmLm2hDtMHUKr",
				wif:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.msg, tt.args.fromPub, tt.args.toPub, tt.args.nonce, tt.args.wif)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
