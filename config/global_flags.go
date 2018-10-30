package config

type GlobalFlags struct {
	Help      bool   `short:"h" long:"help"`
	Debug     bool   `short:"d" long:"debug"        env:"KID_DEBUG"`
	Version   bool   `short:"v" long:"version"`
	NoConfirm bool   `short:"n" long:"no-confirm"`
	StateDir  string `short:"s" long:"state-dir"    env:"KID_STATE_DIRECTORY"`
	EnvID     string `          long:"name"`
	IAAS      string `          long:"iaas"         env:"KID_IAAS"`
}
