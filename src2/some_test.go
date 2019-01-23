package any

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_starts_with_the_word_test(t *testing.T) {
	t.Log("This test will pass")
}

func Test_testing_is_build_in(t *testing.T) {
	if addNumbers(1, 3) != 4 {
		t.Error("this shouldn't happen")
		t.Error("more problems")
		t.FailNow()
		t.Error("this won't get reported")
	}
}

func Test_simple_testing_help(t *testing.T) {
	if addNumbers(4, 6) != 10 {
		assert.Equal(t, 1, 2)
		assert.Equal(t, "more", "issues")

		require.Equal(t, "the", "end")

		assert.Equal(t, "no more", "issues")
	}
}
func Test_with_sub_tests(t *testing.T) {
	t.Log("beginning")

	t.Run("sub test one", func(t *testing.T) {
		t.Log("woot 1")
	})

	t.Log("middle")

	t.Run("sub test two", func(t *testing.T) {
		t.Log("woot 2")
	})

	t.Log("end")
}
