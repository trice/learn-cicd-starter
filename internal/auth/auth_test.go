package auth_test

import(
	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	goodHeader := make(http.Header)
	goodHeader.Set("Authorization", "ApiKey gloppo")
	badHeader := make(http.Header)
	badHeader.Set("Authorization", "")

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name: "Gets an Authkey",
			headers: goodHeader,
			want: "gloppo",
			wantErr: false,
		},
		{
			name: "Missing ApiKey",
			headers: badHeader,
			want: "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := auth.GetAPIKey(tt.headers)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetAPIKey() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetAPIKey() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

