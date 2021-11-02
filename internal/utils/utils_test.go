package utils

import (
	"testing"
)

func TestKeyToAddr(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"503f38a9c967ed597e47fe25643985f032b072db8075426a92110f82df48dfcb",
			"0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"},
		{"7e5bfb82febc4c2c8529167104271ceec190eafdca277314912eaabdb67c6e5f",
			"0xAb8483F64d9C6d1EcF9b849Ae677dD3315835cb2"},
		{"cc6d63f85de8fef05446ebdd3c537c72152d0fc437fd7aa62b3019b79bd1fdd4",
			"0x4B20993Bc481177ec7E8f571ceCaE8A9e22C02db"},
	}

	for _, test := range tests {
		got, err := KeyToAddr(test.input)
		if err != nil {
			t.Error(err)
		}

		if got.String() != test.want {
			t.Errorf("want:%q got:%q", test.want, got)
		}
	}
}
