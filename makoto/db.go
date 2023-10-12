// Copyright Anton Feldmann
//
// store the downloaded kpc package into a local catabase
//
package makoto


// load the packages
import (
	"fmt"
	"strings"
	"sync"

	kpc "github.com/afeldman/Makoto/kpc"
	"github.com/asdine/storm/v3"

	version "github.com/mcuadros/go-version"
	log "github.com/sirupsen/logrus"
)

// the database entry is set by name, the id (the hash of the package,
// the package version and the KPC)
type KPC_DB_Entry struct {
	Name    string `storm:"name"`
	Hash    string `storm:"id"`
	Version string `storm:"version"`
	KPC     string `storm:"kpc"`
}

// KPC Database information
type KPC_DB struct {
     // database
     DB      *storm.DB
     // is the database connected
     DB_Open bool
}

// package variable
var (
	instance *KPC_DB
	lock     = &sync.Mutex{}
	once     sync.Once
)

// Open Database
func (makoto *Makoto) DBInit() *KPC_DB {
     // make singleton
	once.Do(func() {
		// open the database
		sdb, err := storm.Open(makoto.KPC_DB)
		if err != nil {
			log.Panic(err)
		}

		// set the database engine
		instance = &KPC_DB{
			DB:      sdb,
			DB_Open: true,
		}
	})

	// return the instance
	return instance
}

// close the database
func Close() {
     // if the database is open then close the database
	if instance.DB_Open {
	   //close the database
		instance.DB.Close()
		// set the open flag to false
		instance.DB_Open = false
	}
}

// append a kpc to the database
func Append(kpc_s *kpc.KPC) error {

     // is the database conected
	if !instance.DB_Open {
		log.Panic("Database not init")
	}

	// geht the kpc name
	name := *kpc_s.GetName()
	// get the kpc version
	version := *kpc_s.GetVersion()

	hash, kpc_err := kpc.GetMD5Hash(kpc_s)
	if kpc_err != nil {
		return kpc_err
	}

	// create the entry
	entry := KPC_DB_Entry{
	      Name: name + "@" + version,
	      Version: *kpc_s.GetVersion(),
	      Hash: *hash}

	// if the entry exists finish
	kpc_ := Get(name, version)
	if kpc_ != nil {
		return nil
	}

	// safe the new entry into the database
	if err := instance.DB.Save(&entry); err != nil {
	   // return the error, the kpc cound not stored
	   return fmt.Errorf("could not save KPC, %v", err)
	}

	// return nil
	return nil
}

// update the entry
func Update(filepath string) {
     // kpc file to storm database
	kpc, err := kpc.ReadKPCFile(filepath)
	if err != nil {
		log.Error(err)
		return
	}
	Append(kpc)
}

// get the kpc entry from database
func Get(name, version string) *KPC_DB_Entry {
	var entry KPC_DB_Entry
	// is the kpc in the database ?
	if err := instance.DB.One("Name", name+"@"+version, &entry); err != nil {
		log.Debug(err)
		return nil
	}
	// return entry
	return &entry
}

// compare versions of kpc package
func Compare(name, compair string) *KPC_DB_Entry {
     // last entry
	var latest_kpc KPC_DB_Entry

	// last kpc entry package version
	var latest_version string = "0.0.0"

	// search kpc in database
	for _, kpc_e := range All() {
		if name != kpc_e.Name {
			continue
		}

		if version.Compare(latest_version, kpc_e.Version, compair) {
			latest_version = kpc_e.Version
			latest_kpc = kpc_e
		}
	}

	log.Debug(len(strings.TrimSpace(latest_kpc.Name)))

	if len(strings.TrimSpace(latest_kpc.Name)) == 0 {
		return nil
	} else {
		return &latest_kpc
	}
}

// latest version of kpc
func Latest(name string) *KPC_DB_Entry {
     // kpc entry in database
	var latest_kpc KPC_DB_Entry

	var latest_version string = "0.0.0"

	// loop over all entrys of a given kpc name
	for _, kpc_e := range All() {
	    // next if name is not the entry (TODO: select direct in database)
		if name != kpc_e.Name {
			continue
		}
		// get latest version
		if version.Compare(latest_version, kpc_e.Version, "<=") {
			latest_version = kpc_e.Version
			latest_kpc = kpc_e
		}
	}

	log.Debug(len(strings.TrimSpace(latest_kpc.Name)))

	if len(strings.TrimSpace(latest_kpc.Name)) == 0 {
		return nil
	} else {
		return &latest_kpc
	}
}

func All() []KPC_DB_Entry {
	var kpcs []KPC_DB_Entry
	err := instance.DB.All(&kpcs)
	if err != nil {
		log.Fatal(err)
	}
	for _, kpc := range kpcs {
		log.Debugln(kpc)
	}
	return kpcs
}
