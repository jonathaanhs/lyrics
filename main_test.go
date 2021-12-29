package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecite(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectation := "This is the house that Jack built "
		actual := Recite(1)
		assert.Equal(t, expectation, actual)
	})

	t.Run("failed", func(t *testing.T) {
		expectation := "This is the house that Jack built "
		actual := Recite(2)
		assert.NotEqual(t, expectation, actual)
	})
}

func TestReciteRand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		actual := ReciteRand()
		assert.NotEmpty(t, actual)
	})

	t.Run("failed", func(t *testing.T) {
		expectation := ""
		actual := ReciteRand()
		assert.NotEqual(t, expectation, actual)
	})
}

func TestReciteSubject(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		expectation := "This is the house "
		actual := ReciteSubject(1)
		assert.Equal(t, expectation, actual)
	})

	t.Run("failed", func(t *testing.T) {
		expectation := "This is the house that Jack built "
		actual := ReciteSubject(2)
		assert.NotEqual(t, expectation, actual)
	})

	t.Run("recite subject 3 ", func(t *testing.T) {
		expectation := "This is the rat, the malt and the house "
		actual := ReciteSubject(3)
		assert.Equal(t, expectation, actual)
	})
}

func TestReciteSubjectRand(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		actual := ReciteSubjectRand()
		assert.NotEmpty(t, actual)
	})

	t.Run("failed", func(t *testing.T) {
		expectation := ""
		actual := ReciteSubjectRand()
		assert.NotEqual(t, expectation, actual)
	})
}
