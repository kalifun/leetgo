package config

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/yaml.v3"
)

const (
	CmdName           = "leetgo"
	configFile        = "config.yml"
	leetcodeCacheFile = "cache/leetcode-questions.json"
)

var (
	cfg   *Config
	Debug = os.Getenv("DEBUG") != ""
)

type Site string
type Language string

const (
	LeetCodeCN Site     = "https://leetcode.cn"
	LeetCodeUS Site     = "https://leetcode.com"
	ZH         Language = "zh"
	EN         Language = "en"
)

type Config struct {
	Language Language       `yaml:"language" mapstructure:"language" comment:"Language of the questions, zh or en"`
	LeetCode LeetCodeConfig `yaml:"leetcode" mapstructure:"leetcode" comment:"LeetCode configuration"`
	Editor   Editor         `yaml:"editor" mapstructure:"editor"`
	Go       GoConfig       `yaml:"go" mapstructure:"go"`
	Python   PythonConfig   `yaml:"python" mapstructure:"python"`
	// Add more languages here
	dir string
}

type Editor struct {
}

type PythonConfig struct {
	Enable bool   `yaml:"-" mapstructure:"enable"`
	OutDir string `yaml:"out_dir" mapstructure:"out_dir" comment:"Output directory for Python files"`
}

type GoConfig struct {
	Enable           bool   `yaml:"-" mapstructure:"enable"`
	OutDir           string `yaml:"out_dir" mapstructure:"out_dir" comment:"Output directory for Go files"`
	SeparatePackage  bool   `yaml:"separate_package" mapstructure:"separate_package" comment:"Generate separate package for each question"`
	FilenameTemplate string `yaml:"filename_template" mapstructure:"filename_template" comment:"Filename template for Go files"`
}

type LeetCodeConfig struct {
	Site Site `yaml:"site" mapstructure:"site" comment:"LeetCode site"`
}

func (c Config) ConfigDir() string {
	return c.dir
}

func (c Config) ConfigFile() string {
	return filepath.Join(c.dir, configFile)
}

func (c Config) LeetCodeCacheFile() string {
	return filepath.Join(c.dir, leetcodeCacheFile)
}

func (c Config) WriteTo(w io.Writer) error {
	enc := yaml.NewEncoder(w)
	enc.SetIndent(2)
	node, _ := toYamlNode(c)
	err := enc.Encode(node)
	return err
}

func Default() Config {
	home, _ := homedir.Dir()
	configDir := filepath.Join(home, ".config", CmdName)
	return Config{
		dir:      configDir,
		Language: ZH,
		LeetCode: LeetCodeConfig{
			Site: LeetCodeCN,
		},
		Go: GoConfig{
			Enable:           false,
			OutDir:           "go",
			SeparatePackage:  true,
			FilenameTemplate: ``,
		},
		Python: PythonConfig{
			Enable: false,
			OutDir: "python",
		},
		// Add more languages here
	}
}

func Get() Config {
	if cfg == nil {
		return Default()
	}
	return *cfg
}

func Set(c Config) {
	cfg = &c
}

func Verify(c Config) error {
	if c.LeetCode.Site != LeetCodeCN && c.LeetCode.Site != LeetCodeUS {
		return fmt.Errorf("invalid site: %s", c.LeetCode.Site)
	}

	return nil
}
