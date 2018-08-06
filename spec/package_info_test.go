package spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/Makoto/kpc"
)

var _ = Describe("KPC", func() {
	var (
		tkpc KPC
	)

	BeforeEach(func() {
		tkpc = *KPC_Init("Test-Project")
	})

	Context("check the KPC ", func() {
		It("check the KPC Version ", func(){
			Expect( tkpc.GetKPCVersion() ).Should(Equal(KPC_VERSION))
		})
		It("check version ", func(){
			Expect( *tkpc.GetVersion() ).Should(Equal(""))
			tkpc.SetVersion("0.1.0")
			Expect( *tkpc.GetVersion() ).Should(Equal("0.1.0"))
		})
		It("check name ", func(){
			Expect( *tkpc.GetName() ).Should(Equal("Test-Project"))
			Expect( *tkpc.GetName() ).ShouldNot(Equal("Test-Pro"))
			tkpc.SetName("Test1")
			Expect( *tkpc.GetName() ).Should(Equal("Test1"))
		})
		It("check description ", func(){
			Expect( *tkpc.GetDescription() ).Should(Equal(""))
			tkpc.SetDescription("The life is long so fill it with as much fun as possible")
			Expect( *tkpc.GetDescription() ).ShouldNot(Equal(""))
		} )
	})
})
