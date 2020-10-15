package app

import (
	flag "github.com/spf13/pflag"
)

type KVServerFlags struct {
	Host      string
	Port      int
	DebugMode bool
}

const (
	defaultHost       = "localhost"
	defaultListenPort = 10020
)

func NewKVServerFlags() *KVServerFlags {
	return &KVServerFlags{
		Host:      defaultHost,
		Port:      defaultListenPort,
		DebugMode: false,
	}
}

func (f *KVServerFlags) Set(fs *flag.FlagSet) {
	fs.StringVar(&f.Host, "host", f.Host, "listen host")
	fs.IntVarP(&f.Port, "port", "p", f.Port, "listen port")
	fs.BoolVarP(&f.DebugMode, "debug", "d", f.DebugMode, "debug mode")
}

func GetLogLevel(debugMode bool) string {
	if debugMode {
		return "debug"
	}
	return "info"
}
