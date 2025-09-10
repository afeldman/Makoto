// Makoto stores the information for the project
package makoto

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/afeldman/go-util/env"
	"github.com/kirsle/configdir"
	log "github.com/sirupsen/logrus"
)

// makoto config file will be stored in the app config folder
type Makoto struct {
	ConfigPath string `toml:"configpath"`
	ConfigFile string `toml:"config"`
	KPC_DB     string `toml:"db"`
}

// the configpath is either a env or the default config path
var ConfigPath string = env.GetEnvOrDefault("MAKOTO_PATH", configdir.LocalConfig("makoto"))

// init makoto
//
// @return (*Makoto) Makoto structure
func InitMakoto() *Makoto {
	// if the config fir not exists, then create the path
	if err := configdir.MakePath(ConfigPath); err != nil {
		log.Panic(err)
	}

	// user informations
	log.Debugln("the makoto config files are in " + ConfigPath)

	// makoto project information
	return &Makoto{
		ConfigPath: ConfigPath,
		ConfigFile: filepath.Join(ConfigPath, "makoto.toml"),
		KPC_DB:     filepath.Join(ConfigPath, "kpc.db"),
	}
}

// open and read the makoto config fike
func ConfigMakoto(config_path string) *Makoto {

	config_path = strings.TrimSpace(config_path)

	// check if the path is not null or empty
	if len(config_path) == 0 {
		// set default path
		config_path = filepath.Join(ConfigPath, "makoto.toml")
	}

	// load toml file
	tomlfile, err := os.Open(config_path)
	if err != nil {
		log.Fatal(err)
	}
	defer tomlfile.Close()

	// read the file as whole
	tomlData, _ := io.ReadAll(tomlfile)

	// build the Makoto object
	var makoto Makoto
	if err := toml.Unmarshal(tomlData, &makoto); err != nil {
		log.Panic(err)
	}

	// return the makoto
	return &makoto
}

// Store the makoto configuration into a file
func (makoto *Makoto) StoreMakoto() {
	// set buffer
	buf := new(bytes.Buffer)
	//encode the toml structure into a byte buffer
	if err := toml.NewEncoder(buf).Encode(makoto); err != nil {
		log.Panic(err)
	}

	// the toml information a string
	toml_str := buf.String()
	// write the data into file
	if err := os.WriteFile(makoto.ConfigFile, []byte(toml_str), 0644); err != nil {
		log.Panic(err)
	}
}
