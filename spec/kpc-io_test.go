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

	Context("check the KPC io", func(){
		It("get json", func(){
			con := Conflict_Init("test","1.12.31")

			tkpc.GetRepo().SetType("git")
			tkpc.GetRepo().SetURL("github")
			tkpc.Description = "test123123"
			tkpc.Version = "12311231"
			tkpc.Homepage = "github"
			tkpc.Requirements.Add(Requirement{
				Name: "hallo",
				Version: "0.121",
			})
			tkpc.Conflicts.Add(*con)
			Expect( tkpc.ToJSON() ).Should(Equal(`{"kpc_version":"0.0.1","name":"Test-Project","description":"test123123","version":"12311231","url":"github","requirements":[{"name":"hallo","version":"0.121"}],"conflicts":[{"name":"test","version":["1.12.31"]}],"author":[],"source":{"type":"git","url":"github"},"issues":"","prefix":"","srcdir":"","typedir":"","includedir":"","constdir":"","formdir":"","dictdir":"","main":"","dict":[],"form":[],"types":[],"includes":[],"consts":[]}`))
		})
		It("from json", func(){
			ckpc := FromJSON(`{"kpc_version":"0.0.1","name":"Test-Project","description":"test123123","version":"12311231","url":"github","requirements":[{"name":"hallo","version":"0.121"}],"conflicts":[{"name":"test","version":["1.12.31"]}],"author":[],"source":{"type":"git","url":"github"},"issues":"","prefix":"","srcdir":"","typedir":"","includedir":"","constdir":"","formdir":"","dictdir":"","main":"","dict":[],"form":[],"types":[],"includes":[],"consts":[]}`)
			Expect( *(ckpc.GetName()) ).Should(Equal("Test-Project"))
		})
	})
})
