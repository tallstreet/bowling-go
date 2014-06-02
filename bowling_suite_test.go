package bowling_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBowling(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bowling Suite")
}
