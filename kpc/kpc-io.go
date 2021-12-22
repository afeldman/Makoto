// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"os"

	toml "github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

// To() creates a TOML file with all informations
func (kpc_obj *KPC) To() *string {
	// create empty buffer
	buf := new(bytes.Buffer)
	// encode kpc into the buffer
	if err := toml.NewEncoder(buf).Encode(kpc_obj); err != nil {
		return nil
	}
	// convert the buffer to a string
	toml_str := buf.String()
	// return the toml string
	return &toml_str
}

// convert the TOML string to a kpc object
func From(toml_str []byte) *KPC {
	// default kpc object
	var kpc KPC
	// decode the byte object into the kpc structure
	if err := toml.Unmarshal(toml_str, &kpc); err != nil {
		log.Panic(err)
	}
	// return the kpc structure
	return &kpc
}

// read kpc structure from toml file
func ReadKPCFile(filepath string) *KPC {

	// open the file
	input_file, err := os.Open(filepath)
	if err != nil {
		return nil
	}

	// close the file handle on function exit
	defer input_file.Close()

	// read the file
	byteValue, _ := ioutil.ReadAll(input_file)
	// return the kpc structure
	return From(byteValue)
}

func GetMD5Hash(kpc *KPC) string {
	out, err := json.Marshal(kpc)
	if err != nil {
		panic(err)
	}

	hash := md5.Sum([]byte(out))
	return hex.EncodeToString(hash[:])
}
