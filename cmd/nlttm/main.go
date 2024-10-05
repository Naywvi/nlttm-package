package main

import (
	"fmt"
	_args "nlttm/cmd/nlttm/args"
	_config "nlttm/cmd/nlttm/config"
	"os"
)

func main() {
	var (
		config, err = _config.LoadConfig("/etc/nlttm/nlttm.conf") // Chemin mis à jour pour le fichier de config principal
		args        = os.Args[1:]
	)

	if err != nil {
		fmt.Println("Erreur de chargement du fichier de configuration:", err)
		os.Exit(65)
	}

	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			_args.Help(false, args[i])
		case "-v", "--version":
			fmt.Println("Version :", config["version"])
		case "-i", "--install":
			installPackage(args, i)
		case "-d", "--delete":
			deletePackage(args, i)
		case "-c", "--check":
			checkPackage(args, i)
		default:
			_args.Help(false, args[i])
			os.Exit(2)
		}
	}
}

// Fonction pour gérer les prérequis, puis l'installation des packages
func installPackage(args []string, startIndex int) {
	for j := startIndex + 1; j < len(args); j++ {
		if args[j][0] == '-' {
			break
		}

		// Charger les prérequis depuis /etc/nlttm/config/packages/
		prereqs, err := _config.LoadPackageSectionCommands(args[j], "prerequisites")
		if err == nil {
			_args.Prerequisites(args[j], prereqs)
		}

		// Charger et installer le package depuis /etc/nlttm/config/packages/
		commands, err := _config.LoadPackageSectionCommands(args[j], "installation")
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la section installation pour le package %s: %v\n", args[j], err)
			os.Exit(2)
		}
		_args.Install(args[j], commands)
	}
}

// Fonction pour gérer la suppression des packages
func deletePackage(args []string, startIndex int) {
	for j := startIndex + 1; j < len(args); j++ {
		if args[j][0] == '-' {
			break
		}
		// Charger la section "delete" depuis /etc/nlttm/config/packages/
		commands, err := _config.LoadPackageSectionCommands(args[j], "delete")
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la section delete pour le package %s: %v\n", args[j], err)
			os.Exit(2)
		}
		_args.Delete(args[j], commands)
	}
}

// Fonction pour vérifier l'état des packages
func checkPackage(args []string, startIndex int) {
	for j := startIndex + 1; j < len(args); j++ {
		if args[j][0] == '-' {
			break
		}
		// Charger la section "check" depuis /etc/nlttm/config/packages/
		commands, err := _config.LoadPackageSectionCommands(args[j], "check")
		if err != nil {
			fmt.Printf("Erreur lors du chargement de la section check pour le package %s: %v\n", args[j], err)
			os.Exit(2)
		}
		_args.Check(args[j], commands)
	}
}
