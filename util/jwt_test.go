package util

import "testing"

func testCreateToken(t *testing.T) {
	claims := &Claims{
		UserId: 123456,
		Name:   "fugui",
	}
	token, err := CreateToken(claims)
	t.Log(token)
	t.Log(err)
}
