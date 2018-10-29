package config

type GlobalFlags struct {
	Help      bool   `short:"h" long:"help"`
	Debug     bool   `short:"d" long:"debug"        env:"BBL_DEBUG"`
	Version   bool   `short:"v" long:"version"`
	NoConfirm bool   `short:"n" long:"no-confirm"`
	StateDir  string `short:"s" long:"state-dir"    env:"BBL_STATE_DIRECTORY"`
	EnvID     string `          long:"name"`
	IAAS      string `          long:"iaas"         env:"BBL_IAAS"`
}
