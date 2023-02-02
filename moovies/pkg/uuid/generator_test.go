package uuid_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/google/uuid"
	. "github.tools.sap/distribution-store/moovies/pkg/uuid"
)

var _ = Describe("Generator", func() {
	var generator *UUIDGenerator

	BeforeEach(func() {
		generator = NewUUIDGenerator()
	})

	It("generates a non-default uuid", func() {
		Expect(generator.Generate()).ToNot(Equal(uuid.UUID{}))
	})

	It("generates a unique uuid each time", func() {
		Expect(generator.Generate()).ToNot(Equal(generator.Generate()))
	})
})
