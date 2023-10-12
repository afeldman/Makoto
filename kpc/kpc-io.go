// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
package kpc

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"

	toml "github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

// To() creates a TOML file with all informations
func (kpc_obj *KPC) To() (*string, error) {
	// create empty buffer
	buf := new(bytes.Buffer)
	// encode kpc into the buffer
	if err := toml.NewEncoder(buf).Encode(kpc_obj); err != nil {
		return nil, err
	}
	// convert the buffer to a string
	toml_str := buf.String()
	// return the toml string
	return &toml_str, nil
}

// convert the TOML string to a kpc object
func From(toml_str []byte) (*KPC, error) {
	// default kpc object
	var kpc KPC
	// decode the byte object into the kpc structure
	if err := toml.Unmarshal(toml_str, &kpc); err != nil {
		log.Panic(err)
		return nil, err
	}
	// return the kpc structure
	return &kpc, nil
}

// read kpc structure from toml file
func ReadKPCFile(filepath string) (*KPC, error) {

	// open the file
	input_file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("cannot open file %s", filepath)
		return nil, err
	}

	// close the file handle on function exit
	defer input_file.Close()

	// read the file
	byteValue, io_err := ioutil.ReadAll(input_file)
	if io_err != nil {
		return nil, io_err
	}

	// return the kpc structure
	kpc_s, kpc_err := From(byteValue)
	if kpc_err != nil {
		return nil, kpc_err
	}
	return kpc_s, nil
}

func GetMD5Hash(kpc *KPC) (*string, error) {
	toml_str, err := kpc.To()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	h := md5.New()
	io.WriteString(h, *toml_str)

	data := hex.EncodeToString(h.Sum(nil))

	return &data, nil
}
