package spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/Makoto/kpc"
)

const text = `{"kpc_version":"0.0.1","name":"Test-Project","description":"","version":"","url":"","requirements":[{"name":"test","version":"0.1.1"}],"conflicts":[{"name":"test","version":["1.12.31"]}],"source":{"type":"git","url":"github"},"issues":"","prefix":"","srcdir":"","typedir":"","includedir":"","constdir":"","formdir":"","dictdir":"","main":""}`

var _ = Describe("KPC", func() {
	var (
		tkpc KPC
	)

	BeforeEach(func() {
		tkpc = *KPC_Init("Test-Project")
	})

	Context("check the KPC io", func(){
		It("get json", func(){
			con := Conflict_Init("test","1.12.31")
			requirement := Requirement_Init("test")
			requirement.SetVersion("0.1.1")

			tkpc.GetRepo().SetType("git")
			tkpc.GetRepo().SetURL("github")

			tkpc.AddConflict(*con)
			tkpc.AddRequirement(*requirement)
			Expect( string(tkpc.ToJSON()) ).Should(Equal(text))
		})
		//It("from json", func(){
	///
	//		Expect( *(ckpc.GetName()) ).Should(Equal("Test-Project"))
	//	})
	})
})
