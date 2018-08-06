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
		It("check the KPC prefix ", func(){
			Expect( *(tkpc.GetPrefix()) ).Should(Equal(""))
			tkpc.SetPrefix("/usr/local/lib/karel")
			Expect( *(tkpc.GetPrefix()) ).Should(Equal("/usr/local/lib/karel"))
		})
		It("check the KPC source dir ", func(){
			Expect( *(tkpc.GetSourceDir()) ).Should(Equal(""))
			tkpc.SetSourceDir("sourcedir")
			Expect( *(tkpc.GetSourceDir()) ).Should(Equal("sourcedir"))
		})
		It("Check the KPC type dir ", func(){
			Expect( *(tkpc.GetTypeDir()) ).Should(Equal(""))
			tkpc.SetTypeDir("typedir")
			Expect( *(tkpc.GetTypeDir()) ).Should(Equal("typedir"))
		})
		It("check the KPC include dir ", func(){
			Expect( *(tkpc.GetIncludeDir()) ).Should(Equal(""))
			tkpc.SetIncludeDir("include/packagename")
			Expect( *(tkpc.GetIncludeDir()) ).Should(Equal("include/packagename"))
		})
		It("check the KPC constant file dir ", func(){
			Expect( *(tkpc.GetConstDir()) ).Should(Equal(""))
			tkpc.SetConstDir("const")
			Expect( *(tkpc.GetConstDir()) ).Should(Equal("const"))
		})
		It("Check the KPC from dir ", func(){
			Expect( *(tkpc.GetFormDir()) ).Should(Equal(""))
			tkpc.SetFormDir("form")
			Expect( *(tkpc.GetFormDir()) ).Should(Equal("form"))
		})
		It("Check the KPC dict dir ", func(){
			Expect( *(tkpc.GetDictDir()) ).Should(Equal(""))
			tkpc.SetDictDir("dict")
			Expect( *(tkpc.GetDictDir()) ).Should(Equal("dict"))
		})
	})
})
