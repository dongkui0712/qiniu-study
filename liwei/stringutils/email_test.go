package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_VerifyEmail(t *testing.T) {
	assert.True(t, VerifyEmail("contack@gmail.com"))

}
func Test_VerifyEmail2(t *testing.T) {
	assert.False(t, VerifyEmail("contackidrmfly.com"))
}
