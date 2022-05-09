package security

import (
	"log"
	"reflect"
	"testing"
)

func TestHashPassword(t *testing.T) {
	type args struct {
		password string
	}

	var addArgs = []args{
		{"MJl8EA,,4l^+*?&]"},
		{"x7YHY0oPOjWNrdhHKaIxaAEf"},
		{"KglQIpMYGRYwInleyGyQhGaATeOhfkPpQBHaqTzymX"},
		{"75460997590909835263211715119227873972621622506122"},
	}

	for _, tt := range addArgs {
		t.Run(tt.password, func(t *testing.T) {

			got, err := HashPassword(tt.password)

			if err != nil {
				log.Fatalf("Error in HashPassword: %v", err)
			}

			if reflect.TypeOf(got) != reflect.TypeOf(tt.password) {
				log.Fatalf("error hashPassword must be return a string")
			}

			if len(got) != 60 {
				t.Errorf("HashPassword() must be return a string with 60 caracters")
			}
		})
	}
}

func TestVerifyPasswordMatch(t *testing.T) {
	type args struct {
		passwordInput, passwordHash string
	}

	tests := []args{
		{"MJl8EA,,4l^+*?&]", "$2a$12$S.VLK/hliliMI2.mzQwqSOpIsBfK3qH9lXGM3WxfeB1DzAevxBgKC"},
		{"x7YHY0oPOjWNrdhHKaIxaAEf", "$2a$12$Ndov7kXkKVswn8y.QTyyfuG6cEm8rY8GKts8kzeIcUh3o0rHAZazS"},
		{"KglQIpMYGRYwInleyGyQhGaATeOhfkPpQBHaqTzymX", "$2a$10$SSH1L8FQ2GUZ6yfwiSuKvumVhD5wj9Lx.CA/GbTLdwdfeRC1Zuy0O"},
		{"75460997590909835263211715119227873972621622506122", "$2a$10$yuAENIA0JCFS6dxukd7b2.7MKQDqArWOFXCm28uc79Y6CAKNU/Ro6"},
	}

	for _, tt := range tests {
		t.Run(tt.passwordInput, func(t *testing.T) {
			if err := VerifyPasswordMatch(tt.passwordHash, tt.passwordInput); err != nil {
				t.Errorf("VerifyPasswordMatch() error = %v", err)
			}
		})
	}
}
