package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sumanmukherjee03/gotils/cmd/awsutils"
	"github.com/sumanmukherjee03/gotils/cmd/k8s"
	"github.com/sumanmukherjee03/gotils/cmd/logparser"
	"github.com/sumanmukherjee03/gotils/cmd/sshutils"
	"github.com/sumanmukherjee03/gotils/cmd/utils"
)

var (
	rootShortDesc = "Gotils is a single utilty to run various mixes of commands"
	rootLongDesc  = `Gotils is a Flexible tool built with Go.
	It does a mix of various things to help with developer and operations related work.`
	cfgFile string
	Dryrun  bool
)

func main() {
	var rootCmd = &cobra.Command{
		Use:              "gotils [sub]",
		Short:            rootShortDesc,
		Long:             rootLongDesc,
		TraverseChildren: true,
	}

	cobra.OnInitialize(func() { utils.InitConfig(cfgFile) })
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gotils.yml)")
	rootCmd.PersistentFlags().BoolVarP(&Dryrun, "dryrun", "", false, "dryrun")
	rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	viper.BindPFlag("dryrun", rootCmd.PersistentFlags().Lookup("dryrun"))
	viper.BindPFlag("viper", rootCmd.PersistentFlags().Lookup("viper"))
	rootCmd.AddCommand(logparser.InitLogParser())
	rootCmd.AddCommand(awsutils.InitAws())
	rootCmd.AddCommand(k8s.InitK8s())
	rootCmd.AddCommand(sshutils.InitSsh())
	rootCmd.Execute()
}
