package kpc

import (
	"encoding/json"
	"log"
)


func (this *KPC) ToJSON() (string){


	con := this.GetConflicts()
	con_str := "["
	for _, val := range con{
		con_tmp, _ := val.Versions.ToJSON()
		con_str += "{\"name\":\""+val.Name+"\"," +
			"\"version\":"+string(con_tmp)+"}"
	}
	con_str += "]"

	req, _ := this.Requirements.ToJSON()
	aut, _ := this.Authors.ToJSON()
	dicts, _ := this.Dicts.ToJSON()
	forms, _ := this.Forms.ToJSON()
	types, _ := this.Types.ToJSON()
	includes, _ := this.Includes.ToJSON()
	consts, _ := this.Consts.ToJSON()

	return "{\"kpc_version\":\"" + this.GetKPCVersion() + "\","+
		"\"name\":\"" + this.Name + "\"," +
		"\"description\":\""+this.Description+"\","+
		"\"version\":\""+ this.Version+"\","+
		"\"url\":\""+this.Homepage+"\","+
		"\"requirements\":"+string(req)+","+
		"\"conflicts\":"+con_str+","+
		"\"author\":"+string(aut)+","+
		"\"source\":{\"type\":\""+this.Repository.Type+"\","+
		             "\"url\":\""+this.Repository.URL+"\"},"+
		"\"issues\":\""+this.Issue+"\","+
		"\"prefix\":\""+this.Prefix+"\","+
		"\"srcdir\":\""+this.SrcDir+"\","+
		"\"typedir\":\""+this.TypeDir+"\","+
		"\"includedir\":\""+this.IncludeDir+"\","+
		"\"constdir\":\""+this.ConstDir+"\","+
		"\"formdir\":\""+this.FormDir+"\","+
		"\"dictdir\":\""+this.DictDir+"\","+
		"\"main\":\""+this.Main+"\","+
		"\"dict\":"+string(dicts)+","+
		"\"form\":"+string(forms)+","+
		"\"types\":"+string(types)+","+
		"\"includes\":"+string(includes)+","+
		"\"consts\":"+string(consts)+
		"}"
}

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			log.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			log.Printf("}\n")
		} else {
			log.Printf("%v %v : %v\n", space, k, v)
		}
	}
}


func FromJSON(json_str string) *KPC {
	raw := make(map[string]interface{})
	err := json.Unmarshal([]byte(json_str), &raw)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	dumpMap(" ",raw)

	kpc := KPC_Init(raw["name"].(string))
	kpc.SetPrefix(raw["prefix"].(string))
	kpc.SetVersion(raw["version"].(string))
	kpc.SetDescription(raw["description"].(string))
	kpc.Homepage = raw["url"].(string)

	repo := Repo{}
	for k, v := range raw["source"].(map[string]interface{}) {
		if k == "type" {
			repo.Type = v.(string)
		}
		if k == "url" {
			repo.URL = v.(string)
		}
	}


//	repo := Repo{}
//	err_repo := json.Unmarshal([]byte(raw["source"]), &repo)
//	if err_repo != nil {
//		log.Fatal(err_repo)
//		panic(err_repo)
//	}

//	kpc.Repository = &repo

	return kpc
}
