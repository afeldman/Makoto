package spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/afeldman/Makoto/kpc"
)

const text = `{"kpc_version":"0.0.1","name":"Test-Project","version":"","requirements":[{"name":"test","version":"0.1.1"}],"conflicts":[{"name":"test","version":["1.12.31"]}],"author":[],"source":{"type":"git","url":"github"},"prefix":"","main":""}`

var _ = Describe("KPC", func() {
	var (
		tkpc KPC
	)

	Context("check the KPC io", func(){
		It("get json", func(){
			tkpc = *KPC_Init("Test-Project")
			con := Conflict_Init("test","1.12.31")
			requirement := Requirement_Init("test")
			requirement.SetVersion("0.1.1")

			tkpc.GetRepo().SetType("git")
			tkpc.GetRepo().SetURL("github")

			tkpc.AddConflict(*con)
			tkpc.AddRequirement(*requirement)
			_,json_s := tkpc.ToJSON()
			Expect( string(json_s) ).Should(Equal(text))
		})
		It("from json", func(){
			_, tkpc := FromJSON([]byte(text))
			Expect( *(tkpc.GetName()) ).Should(Equal("Test-Project"))
			Expect( *(tkpc.GetVersion()) ).Should(Equal(""))
			Expect( *(tkpc.GetRepo().GetType()) ).Should(Equal("git"))
		})
	})
})
