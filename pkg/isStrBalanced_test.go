package pkg_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"main.go/pkg"
)

func Test_IsStrBalanced(t *testing.T) {
	t.Parallel()

	t.Run("case: odd number of brackets ", func(t *testing.T) {
		t.Parallel()
		str := "[){}]"

		isBalanced, err := pkg.IsStrBalanced(str)

		require.NoError(t, err)
		require.Equal(t, isBalanced, false)
	})
	//case: закрывающая ")" стоит впереди открывающей "("
	t.Run("case: closed round bracket is placed before open one ", func(t *testing.T) {
		t.Parallel()
		str := "[){}]("

		isBalanced, err := pkg.IsStrBalanced(str)

		require.NoError(t, err)
		require.Equal(t, isBalanced, false)
	})
	//case: закрывающая "]" стоит впереди открывающей "["
	t.Run("case: closed square bracket is placed before open one ", func(t *testing.T) {
		t.Parallel()
		str := "{}]["

		isBalanced, err := pkg.IsStrBalanced(str)

		require.NoError(t, err)
		require.Equal(t, isBalanced, false)
	})
	//case: закрывающая "}" стоит впереди открывающей "{"
	t.Run("case: closed curly bracket is placed before open one ", func(t *testing.T) {
		t.Parallel()
		str := "}[[]]{"

		isBalanced, err := pkg.IsStrBalanced(str)

		require.NoError(t, err)
		require.Equal(t, isBalanced, false)
	})
	//case: неправильный ввод
	t.Run("error: invalid input ", func(t *testing.T) {
		t.Parallel()
		str := "f}[[]]{{"

		isBalanced, err := pkg.IsStrBalanced(str)

		require.ErrorContains(t, err, "Invalid character f. Use only ()[]{} brackets to check")
		require.Equal(t, isBalanced, false)
	})

	// case: не все скобки закрываются
	t.Run("case: round+square+curly not equal 0 ", func(t *testing.T) {
		t.Parallel()
		str := "[]{}(){{"

		isBalanced, err := pkg.IsStrBalanced(str)

		require.NoError(t, err)
		require.Equal(t, isBalanced, false)
	})

	t.Run("success ", func(t *testing.T) {
		t.Parallel()
		str := "[]{}()"

		isBalanced, err := pkg.IsStrBalanced(str)

		require.NoError(t, err)
		require.Equal(t, isBalanced, true)
	})

}
