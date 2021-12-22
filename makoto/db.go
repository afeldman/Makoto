package makoto

import (
	"fmt"
	"strings"
	"sync"

	"github.com/afeldman/Makoto/kpc"
	"github.com/asdine/storm/v3"

	version "github.com/mcuadros/go-version"
	log "github.com/sirupsen/logrus"
)

type KPC_DB_Entry struct {
	Name    string `storm:"name"`
	Hash    string `storm:"id"`
	Version string `storm:"version"`
	KPC     string `storm:"karel"`
}

type KPC_DB struct {
	DB      *storm.DB
	DB_Open bool
}

var (
	instance *KPC_DB
	lock     = &sync.Mutex{}
	once     sync.Once
)

func (makoto *Makoto) DBInit() *KPC_DB {

	once.Do(func() {
		sdb, err := storm.Open(makoto.KPC_DB)
		if err != nil {
			log.Panic(err)
		}

		instance = &KPC_DB{
			DB:      sdb,
			DB_Open: true,
		}
	})

	return instance
}

func Close() {
	if instance.DB_Open {
		instance.DB.Close()
		instance.DB_Open = false
	}
}

func Append(kpc *kpc.KPC) error {

	if !instance.DB_Open {
		log.Panic("Database not init")
	}

	name := *kpc.GetName()
	version := *kpc.GetVersion()

	entry := KPC_DB_Entry{Name: name + "@" + version, Version: *kpc.GetVersion(), Hash: kpc.GetMD5Hash(kpc)}

	kpc_ := Get(name, version)
	if kpc_ != nil {
		return nil
	}

	if err := instance.DB.Save(&entry); err != nil {
		return fmt.Errorf("could not save KPC, %v", err)
	}

	return nil
}

func Update(filepath string) {
	Append(kpc.ReadKPCFile(filepath))
}

func Get(name, version string) *KPC_DB_Entry {
	var entry KPC_DB_Entry
	if err := instance.DB.One("Name", name+"@"+version, &entry); err != nil {
		log.Debug(err)
		return nil
	}
	return &entry
}

func Compare(name, compair string) *KPC_DB_Entry {
	var latest_kpc KPC_DB_Entry

	var latest_version string = "0.0.0"

	for _, kpc := range All() {
		if name != kpc.KPC.Name {
			continue
		}

		if version.Compare(latest_version, kpc.KPC.Version, compair) {
			latest_version = kpc.KPC.Version
			latest_kpc = kpc
		}
	}

	log.Debug(len(strings.TrimSpace(latest_kpc.KPC.Name)))

	if len(strings.TrimSpace(latest_kpc.KPC.Name)) == 0 {
		return nil
	} else {
		return &latest_kpc
	}
}

func Latest(name string) *KPC_DB_Entry {
	var latest_kpc KPC_DB_Entry

	var latest_version string = "0.0.0"

	for _, kpc := range All() {
		if name != kpc.KPC.Name {
			continue
		}

		if version.Compare(latest_version, kpc.KPC.Version, "<=") {
			latest_version = kpc.KPC.Version
			latest_kpc = kpc
		}
	}

	log.Debug(len(strings.TrimSpace(latest_kpc.KPC.Name)))

	if len(strings.TrimSpace(latest_kpc.KPC.Name)) == 0 {
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
