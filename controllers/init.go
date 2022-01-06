package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func InitController(cmd *cobra.Command, args []string) {

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Kuknos Cli inited before in: ", viper.ConfigFileUsed())
		return
	}

	// Main Branch [default: master/]
	fmt.Println("Initilizing...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter main branch name [default: master/]: ")
	mainBranch, _ := reader.ReadString('\n')
	mainBranch = strings.TrimSpace(mainBranch)
	if mainBranch == "" {
		mainBranch = "master/"
	}

	// Develop Branch [default: develop/]
	fmt.Print("Enter develop branch name [default: develop/]: ")
	developBranch, _ := reader.ReadString('\n')
	developBranch = strings.TrimSpace(developBranch)
	if developBranch == "" {
		developBranch = "develop/"
	}

	// Prefix of feature branch [default: feature/]
	fmt.Print("Enter feature branch prefix [default: feature/]: ")
	featurePrefix, _ := reader.ReadString('\n')
	featurePrefix = strings.TrimSpace(featurePrefix)
	if featurePrefix == "" {
		featurePrefix = "feature/"
	}

	// Prefix of bugfix branch [default: bugfix/]
	fmt.Print("Enter bugfix branch prefix [default: bugfix/]: ")
	bugfixPrefix, _ := reader.ReadString('\n')
	bugfixPrefix = strings.TrimSpace(bugfixPrefix)
	if bugfixPrefix == "" {
		bugfixPrefix = "bugfix/"
	}

	// Prefix of hotfix branch [default: hotfix/]
	fmt.Print("Enter hotfix branch prefix [default: hotfix/]: ")
	hotfixPrefix, _ := reader.ReadString('\n')
	hotfixPrefix = strings.TrimSpace(hotfixPrefix)
	if hotfixPrefix == "" {
		hotfixPrefix = "hotfix/"
	}

	// Prefix of release branch [default: release/]
	fmt.Print("Enter release branch prefix [default: release/]: ")
	releasePrefix, _ := reader.ReadString('\n')
	releasePrefix = strings.TrimSpace(releasePrefix)
	if releasePrefix == "" {
		releasePrefix = "release/"
	}

	// Prefix of improvement branch [default: improvement/]
	fmt.Print("Enter improvement branch prefix [default: improvement/]: ")
	improvementPrefix, _ := reader.ReadString('\n')
	improvementPrefix = strings.TrimSpace(improvementPrefix)
	if improvementPrefix == "" {
		improvementPrefix = "improvement/"
	}

	// Save to config file (yaml)
	viper.Set("mainBranch", mainBranch)
	viper.Set("developBranch", developBranch)
	viper.Set("featurePrefix", featurePrefix)
	viper.Set("bugfixPrefix", bugfixPrefix)
	viper.Set("hotfixPrefix", hotfixPrefix)
	viper.Set("improvementPrefix", improvementPrefix)

	err := viper.SafeWriteConfig()

	if err != nil {
		fmt.Println("Error writing config file: ", err)
		return
	}

	fmt.Println("Kuknos git cli inited successfully!")
}
