package makoto

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/afeldman/go-util/env"
	"github.com/kirsle/configdir"
	log "github.com/sirupsen/logrus"
)

type Makoto struct {
	ConfigPath string `toml:"configpath"`
	ConfigFile string `toml:"config"`
	KPC_DB     string `toml:"db"`
}

var ConfigPath string = env.GetEnvOrDefault("MAKOTO_PATH", configdir.LocalConfig("makoto"))

func Init() *Makoto {

	if err := configdir.MakePath(ConfigPath); err != nil {
		log.Panic(err)
	}

	log.Debugln("the makoto config files are in " + ConfigPath)

	return &Makoto{
		ConfigPath: ConfigPath,
		ConfigFile: filepath.Join(ConfigPath, "makoto.toml"),
		KPC_DB:     filepath.Join(ConfigPath, "kpc.db"),
	}
}

func Config(config_path string) *Makoto {
	if len(strings.TrimSpace(config_path)) == 0 {
		config_path = filepath.Join(ConfigPath, "makoto.toml")
	}

	tomlfile, err := os.Open(config_path)
	if err != nil {
		log.Fatal(err)
	}
	defer tomlfile.Close()

	tomlData, _ := ioutil.ReadAll(tomlfile)

	var makoto Makoto
	if err := toml.Unmarshal(tomlData, &makoto); err != nil {
		log.Panic(err)
	}
	return &makoto
}

func (makoto *Makoto) Store() {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(makoto); err != nil {
		log.Panic(err)
	}
	toml_str := buf.String()
	if err := os.WriteFile(makoto.ConfigFile, []byte(toml_str), 0644); err != nil {
		log.Panic(err)
	}
}
