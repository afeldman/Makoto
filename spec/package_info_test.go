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
		})
		It("check homepage ", func(){
			Expect( *tkpc.GetHomepage() ).Should(Equal(""))
			tkpc.SetHomepage("github...")
			Expect( *tkpc.GetHomepage() ).ShouldNot(Equal(""))
			Expect( *tkpc.GetHomepage() ).Should(Equal("github..."))
		})
		It("check Requirements ", func(){
			Expect( tkpc.RequirementSize() ).Should(Equal(0))
			req := Requirement_Init("test")
			tkpc.AddRequirement(*req)
			Expect( (*tkpc.GetRequirement("test")).Name ).Should(Equal("test"))
			Expect( tkpc.RequirementSize() ).Should(Equal(1))
			tkpc.RejectRequirement("test")
			Expect( tkpc.RequirementSize() ).Should(Equal(0))
		})
		It("check Requirements ", func(){
			Expect( tkpc.ConflictSize() ).Should(Equal(0))
			con := Conflict_Init("test","0.1.4")
			tkpc.AddConflict(*con)
			Expect( (tkpc.GetConflict("test")).Name ).Should(Equal("test"))
			Expect( tkpc.ConflictSize() ).Should(Equal(1))
			tkpc.RejectConflict("test")
			Expect( tkpc.ConflictSize() ).Should(Equal(0))
		})
		It("check Authors ", func(){
			Expect( tkpc.AuthorsSize() ).Should(Equal(0))
			aut := Author_Init("anton")
			tkpc.AddAuthor(*aut)
			Expect( (tkpc.GetAuthor("anton")).Name ).Should(Equal("anton"))
			Expect( tkpc.AuthorsSize() ).Should(Equal(1))
			tkpc.RejectAuthor("anton")
			Expect( tkpc.AuthorsSize() ).Should(Equal(0))
		})
		It("check Repo ", func(){
			repo := Repo_Init()

			tkpc.AddRepo(*repo)
			Expect( *((tkpc.GetRepo()).GetType()) ).Should(Equal(""))
			Expect( *((tkpc.GetRepo()).GetURL()) ).Should(Equal(""))

			repo2 := tkpc.GetRepo()
			repo2.SetType("git")
			repo2.SetURL("github")
			Expect( *(repo2).GetType() ).Should(Equal("git"))
			Expect( *(repo2).GetURL() ).Should(Equal("github"))
		})
		It("Check Issue ", func(){
			Expect( *(tkpc.GetIssue()) ).Should(Equal(""))
			tkpc.SetIssue("github.issuepage")
			Expect( *(tkpc.GetIssue()) ).Should(Equal("github.issuepage"))
		})
	})
})
