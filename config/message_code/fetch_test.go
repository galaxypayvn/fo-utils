package messagecode_test

import (
	"context"
	"testing"

	messagecode "code.finan.cc/finan-one-be/fo-utils/config/message_code"
)

func TestFetchMessageCode(t *testing.T) {
	cfg := messagecode.Config{
		RedisAddr:            "localhost:6379",
		StrapiMessageCodeURL: "https://strapi.sieu.re/api/message-codes",
		StrapiToken:          "2d3602ab9fe41ae9c669586d54d0173a8a2815a25810e3df92a86bd1cbfd31ed1941d1ea1346fdc1d48014620353f2e30704a63f82b10d217f75f0b2e0fcb4152c0568b75b8a4d465bf2dfd9a0880eaa8faf29656a74f9ea90cb9941b145c49407ca9491feea6eca37ba0378509152d77c76116ba6f7d8c932174d1f83c3ad1a",
	}

	client := messagecode.NewClient(cfg)
	messageCodes, err := client.GetMessage(context.Background(), 11, []int{102000, 102001})
	if err != nil {
		t.Errorf("want no error, got %s", err)
	}
	t.Logf("message_code: %+v\n", messageCodes)
}
