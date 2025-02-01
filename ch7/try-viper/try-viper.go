package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	switch n {
	case "pass":
		n = "password"
		break
	case "ps":
		n = "password"
		break
	}
	return pflag.NormalizedName(n)
}

func main() {
	pflag.StringP("name", "n", "Mike", "Name parameter")
	pflag.StringP("password", "p", "hardToGuess", "Password")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)
	pflag.Parse()
	// viper.BindPFlags() call makes all flags available to the viper package.
	// strictly speaking, we say that the viper.BindPFlags() call binds an existing set of pflag flags
	// (pflag.FlagSet) to viper.
	viper.BindPFlags(pflag.CommandLine)
	// To read the values of two string "name and password" command line flags using viper.GetString()
	name := viper.GetString("name")
	password := viper.GetString("password")
	fmt.Println(name, password)
	// Reading an Environment variable
	// The viper package can also work with environment variables. We first need to call viper.
	// BindEnv() to tell viper which environment variable interests us, and then we can read its value
	// by calling viper.Get().
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS: ", val)
	}
	// Setting an Environment variable
	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)
}
