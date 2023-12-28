package util

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	Vp *viper.Viper
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
func NewSetting(path, name, filetype string) (*Setting, error) {
	vp := viper.New()
	vp.AddConfigPath(path)
	vp.SetConfigName(name)
	vp.SetConfigType("env")
	//vp.AutomaticEnv()
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	s := &Setting{vp}

	if filetype == "json" {
		s.WatchSettingChange()
	}

	return s, nil
}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.Vp.WatchConfig()
		s.Vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.Vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Setting) EnvReadSection(section interface{}) error {
	err := s.Vp.Unmarshal(section)
	return err
}
