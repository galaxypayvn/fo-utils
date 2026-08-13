package messagecode

import (
	"context"
	"net/http"
	"testing"
)

func TestLoadEmbeddedJSON(t *testing.T) {
	codes, err := loadEmbeddedMessageCodes()
	if err != nil {
		t.Fatalf("unmarshal catalog: %v", err)
	}
	if len(codes) != 132 {
		t.Fatalf("catalog size: got %d want 132", len(codes))
	}
}

func TestGetMessageFromJSON(t *testing.T) {
	c := &Client{messageMap: map[string]messageCode{}}
	if err := c.LoadMessageCode(context.Background(), 22); err != nil {
		t.Fatalf("LoadMessageCode: %v", err)
	}

	tests := []struct {
		name     string
		locale   string
		code     int
		wantMsg  string
		wantHTTP int
	}{
		{"success vi", "vi", GeneralSuccessCode, "Thực thi API thành công", http.StatusOK},
		{"bad request vi", "vi", GeneralBadRequestCode, "Yêu cầu không hợp lệ", http.StatusBadRequest},
		{"bad request en", "en", GeneralBadRequestCode, "Invalid request", http.StatusBadRequest},
		{"unauthorized last-write en", "en", GeneralUnauthorizedCode, "Request denied", http.StatusUnauthorized},
		{"unauthorized last-write vi", "vi", GeneralUnauthorizedCode, "Xác thực không thành công", http.StatusUnauthorized},
		{"group 22", "vi", 224055, "Không tìm thấy yêu cầu thanh toán", http.StatusNotFound},
		{"missing code", "vi", 999999, "Không tìm thấy message", http.StatusInternalServerError},
		{"success en missing in catalog", "en", GeneralSuccessCode, "Message not found", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.GetMessage(tt.locale, tt.code); got != tt.wantMsg {
				t.Errorf("GetMessage = %q want %q", got, tt.wantMsg)
			}
			if got := c.GetHTTPCode(tt.locale, tt.code); got != tt.wantHTTP {
				t.Errorf("GetHTTPCode = %d want %d", got, tt.wantHTTP)
			}
		})
	}
}

func TestNewClientDoesNotNeedStrapi(t *testing.T) {
	client, err := NewClient(Config{MessageGroup: []int{22}})
	if err != nil {
		t.Fatalf("NewClient: %v", err)
	}
	if got := client.GetMessage("vi", 102000); got != "Thực thi API thành công" {
		t.Fatalf("got %q", got)
	}
}
