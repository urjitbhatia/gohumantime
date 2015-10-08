package gohumantime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGohumantime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gohumantime Suite")
}
