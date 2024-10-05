package main

import (
	"fmt"
	_args "nlttm/cmd/nlttm/args"
	_config "nlttm/cmd/nlttm/config"
	"os"
)

func main() {
	var (
		save_args         [][]string
		config, err       = _config.LoadConfig("config/nlttm.conf")
		packages, errPack = _config.LoadPackagesConfig("config/packages.conf")
		args              = os.Args[1:]
	)

	if err != nil {
		fmt.Println("Erreur de chargement du fichier de configuration:", err)
		os.Exit(65)
	}

	if errPack != nil {
		fmt.Println("Erreur de chargement du fichier des packages:", errPack)
		os.Exit(66)
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			save_args = append(save_args, []string{"help"})
		case "-v", "--version":
			save_args = append(save_args, []string{"version"})
			fmt.Println("Version :", config["version"])
		case "-i", "--install":
			var installArgs []string
			for j := i + 1; j < len(args); j++ {
				if args[j][0] == '-' {
					break
				}
				if isValidPackage(args[j], packages["install"]) {
					installArgs = append(installArgs, args[j])
				} else {
					fmt.Print(args[i][0])
					fmt.Printf("Package inconnu : %s\n", args[j])
					os.Exit(2)
				}
				i++
			}
			save_args = append(save_args, append([]string{"install"}, installArgs...))
		default:
			_args.Help(false, args[i])
			os.Exit(2)
		}
	}
	for _, arg := range args[1:] {
		if commands, ok := packages[arg]; ok {
			_args.Install(arg, commands)
		} else if arg != "-i" && arg != "--install" {
			fmt.Printf("Package inconnu : %s\n", arg)
			os.Exit(20)
		}
	}
	//fmt.Print(save_args)

}

// Fonction pour vérifier si les packages sont valides
func isValidPackage(arg string, validPackages []string) bool {
	for _, pkg := range validPackages {
		if arg == pkg {
			return true
		}
	}
	return false
}
