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

	Context("check the KPC ", func(){
		It("check main source", func(){
			Expect( *(tkpc.GetMainSourceFile()) ).Should(Equal(""))
			tkpc.SetMainSourceFile("test")
			Expect( *(tkpc.GetMainSourceFile()) ).Should(Equal("test"))
		})
		It("check consts", func(){
			Expect( tkpc.ConstsSize() ).Should(Equal(0))
			tkpc.AddConst("abc")
			Expect( *(tkpc.GetConst("abc")) ).Should(Equal("abc"))
			Expect( tkpc.ConstsSize() ).Should(Equal(1))
			tkpc.RejectConst("abc")
			Expect( tkpc.ConstsSize() ).Should(Equal(0))
			tkpc.AddConst("abc")
			tkpc.AddConst("cba")
			consts := tkpc.GetConsts()
			Expect(len(consts) ).Should(Equal(2))
		})
		It("check includes", func(){
			Expect( tkpc.IncludesSize() ).Should(Equal(0))
			tkpc.AddInclude("abc")
			Expect( *(tkpc.GetInclude("abc")) ).Should(Equal("abc"))
			Expect( tkpc.IncludesSize() ).Should(Equal(1))
			tkpc.RejectInclude("abc")
			Expect( tkpc.IncludesSize() ).Should(Equal(0))
			tkpc.AddInclude("abc")
			tkpc.AddInclude("cba")
			inc := tkpc.GetIncludes()
			Expect(len(inc) ).Should(Equal(2))
		})
		It("check forms", func(){
			Expect( tkpc.FormsSize() ).Should(Equal(0))
			tkpc.AddForm("abc")
			Expect( *(tkpc.GetForm("abc")) ).Should(Equal("abc"))
			Expect( tkpc.FormsSize() ).Should(Equal(1))
			tkpc.RejectForm("abc")
			Expect( tkpc.FormsSize() ).Should(Equal(0))
			tkpc.AddForm("abc")
			tkpc.AddForm("cba")
			forms := tkpc.GetForms()
			Expect( len(forms) ).Should(Equal(2))
		})
		It("check dict", func(){
			Expect( tkpc.DictsSize() ).Should(Equal(0))
			tkpc.AddDict("abc")
			Expect( *(tkpc.GetDict("abc")) ).Should(Equal("abc"))
			Expect( tkpc.DictsSize() ).Should(Equal(1))
			tkpc.RejectDict("abc")
			Expect( tkpc.DictsSize() ).Should(Equal(0))
			tkpc.AddDict("abc")
			tkpc.AddDict("cba")
			dicts := tkpc.GetDicts()
			Expect( len(dicts) ).Should(Equal(2))
		})
		It("check type", func(){
			Expect( tkpc.TypesSize() ).Should(Equal(0))
			tkpc.AddType("abc")
			Expect( *(tkpc.GetType("abc")) ).Should(Equal("abc"))
			Expect( tkpc.TypesSize() ).Should(Equal(1))
			tkpc.RejectType("abc")
			Expect( tkpc.TypesSize() ).Should(Equal(0))
			tkpc.AddType("abc")
			tkpc.AddType("cba")
			types := tkpc.GetTypes()
			Expect( len(types) ).Should(Equal(2))
		})
	})

})
