package spec_test

import (
	//"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/Makoto/kpc"
)

var _ = Describe("Requirements", func() {
	var (
		req1 Requirement
		req2 Requirement
	)

	BeforeEach(func() {
		req1 = *Requirement_Init("test")
		req2 = *Requirement_Init("test2")
		req2.SetVersion("0.1.1")
	})

	Context("version of requirements", func() {
		It("Requirement with Version 0.1.1", func(){
			Expect(*(req2.GetVersion())).Should(Equal("0.1.1"))
		})
		It("Requirement with Version 0.1.0", func(){
			Expect(*(req1.GetVersion())).Should(Equal("0.1.0"))
		})
		It("Requirement not with 0.1.1", func(){
			req2.SetVersion("0.1.2")
			Expect(*(req2.GetVersion())).ShouldNot(Equal("0.1.1"))
		})
		It("Requirement not with 0.1.0", func(){
			req1.SetVersion("0.2.1")
			Expect(*(req1.GetVersion())).ShouldNot(Equal("0.1.0"))
		})
	})

	Context("change the version", func(){
		It("Change the Requirement Version Requirement 1", func(){
			Expect(*(req1.GetVersion())).Should(Equal("0.1.0"))
			req1.SetVersion("1.2.3")
			Expect(*(req1.GetVersion())).Should(Equal("1.2.3"))
		})
		It("Change the Requirement Version Requirement 2", func(){
			Expect(*(req2.GetVersion())).Should(Equal("0.1.1"))
			req2.SetVersion("2.2.3")
			Expect(*(req2.GetVersion())).Should(Equal("2.2.3"))
		})
	})
})
