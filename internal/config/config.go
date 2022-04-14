package config

type (
	Config struct {
		I18n     I18n     `yaml:"i18n"`
		Bot      Bot      `yaml:"bot"`
		Log      Log      `yaml:"log"`
		Database Database `yaml:"database"`
	}

	I18n struct {
		BundlePath string `yaml:"bundle_path"`
	}

	Bot struct {
		Token         string   `yaml:"token"`
		UseProxy      bool     `yaml:"use_proxy"`
		ProxyAddress  string   `yaml:"proxy_address"`
		AllowedGroups []string `yaml:"allowed_groups"`
	}

	Log struct {
		Path string `yaml:"path"`
	}

	Database struct {
		Type       string `yaml:"type"`
		Path       string `yaml:"path"`
		JsonIndent int    `yaml:"json_indent"`
	}
)
