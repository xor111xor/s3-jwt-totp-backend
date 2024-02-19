/*
Copyright © 2024 xor111xor
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/api"
	"github.com/xor111xor/s3-jwt-totp-backend/internal/domain"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Storage API server",
	Short: "S3 API Server",
	RunE: func(cmd *cobra.Command, args []string) error {
		var config domain.SysConfig

		err := viper.Unmarshal(&config)
		if err != nil {
			return err
		}

		repo, err := GetRepoDB(config.DBConfig)
		if err != nil {
			return err
		}
		storageS3, err := GetStorageS3(config)
		if err != nil {
			return err
		}
		commonConfig := domain.NewCommonConfig(config, repo, GetCache(), storageS3)
		err = api.RunApi(commonConfig)
		if err != nil {
			return err
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/storage-server.toml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("toml")
		viper.SetConfigName(".config/storage-server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
