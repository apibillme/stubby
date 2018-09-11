package stubby

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStubEnv(t *testing.T) {
	os.Setenv("STUBBY_T1", "V1")
	os.Setenv("STUBBY_T2", "V2")
	os.Unsetenv("STUBBY_NONE")

	stubs := New()

	stubs.SetEnv("STUBBY_NONE", "a")
	stubs.SetEnv("STUBBY_T1", "1")
	stubs.SetEnv("STUBBY_T1", "2")
	stubs.SetEnv("STUBBY_T1", "3")
	stubs.SetEnv("STUBBY_T2", "4")
	stubs.UnsetEnv("STUBBY_T2")

	assert.Equal(t, "3", os.Getenv("STUBBY_T1"), "Wrong value for T1")
	assert.Equal(t, "", os.Getenv("STUBBY_T2"), "Wrong value for T2")
	assert.Equal(t, "a", os.Getenv("STUBBY_NONE"), "Wrong value for NONE")
	stubs.Reset()

	_, ok := os.LookupEnv("STUBBY_NONE")
	assert.False(t, ok, "NONE should be unset")

	assert.Equal(t, "V1", os.Getenv("STUBBY_T1"), "Wrong reset value for T1")
	assert.Equal(t, "V2", os.Getenv("STUBBY_T2"), "Wrong reset value for T2")
}
