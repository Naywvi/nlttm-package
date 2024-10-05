package _args

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
)

// Variable pour stocker le mot de passe MySQL
var mysqlPassword string

// Fonction pour exécuter une série de commandes avec un loader et des couleurs
func executeCommands(sectionName string, packageName string, commands []string) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("%s %s %s\n", green("Section :"), yellow(sectionName), green("du package"), yellow(packageName))

	bar := progressbar.NewOptions(len(commands),
		progressbar.OptionSetDescription(fmt.Sprintf("Exécution des commandes de %s...", sectionName)),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(40),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionFullWidth(),
	)

	for _, cmd := range commands {
		cmd = strings.TrimSpace(cmd)
		if cmd == "" {
			continue
		}

		bar.Add(1)

		// Exécution de la commande
		out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
		if err != nil {
			// Afficher l'erreur avec tous les détails
			fmt.Printf("%s %s\n", red("Erreur lors de l'exécution de la commande :"), cmd)
			fmt.Printf("%s %s\n", red("Détail de l'erreur :"), err.Error())
			fmt.Printf("Sortie de la commande : %s\n", out)
			continue
		}
		fmt.Printf("Sortie de la commande : %s\n", out)
	}

	fmt.Printf("%s %s terminé pour le package %s\n", green("[ Section ] "), yellow(sectionName), yellow(packageName))
}

// Fonction pour installer un package
func Install(packageName string, commands []string) {
	executeCommands("installation", packageName, commands)
}

// Fonction pour supprimer un package
func Delete(packageName string, commands []string) {
	executeCommands("delete", packageName, commands)
}

// Fonction pour vérifier un package
func Check(packageName string, commands []string) {
	executeCommands("check", packageName, commands)
}

// Exemple d'utilisation dans la gestion des prérequis
func Prerequisites(packageName string, commands []string) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Printf("%s Vérification des prérequis pour %s\n", yellow("=>"), yellow(packageName))

	// Demander le mot de passe MySQL une seule fois
	if mysqlPassword == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Veuillez entrer le mot de passe root MySQL: ")
		password, _ := reader.ReadString('\n')
		mysqlPassword = strings.TrimSpace(password)
	}

	for _, cmd := range commands {
		// Si c'est la commande mysql_secure_installation_auto, appeler la fonction SecureMySQL
		if strings.Contains(cmd, "mysql_secure_installation_auto") {
			SecureMySQL()
			continue
		}

		// Si c'est une commande MySQL, insérer le mot de passe
		if strings.Contains(cmd, "mysql -u root") {
			cmd = strings.Replace(cmd, "YourPassword", mysqlPassword, -1)
		}

		// Exécuter la commande
		fmt.Printf("Exécution de la commande : %s\n", cmd)
		out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
		if err != nil {
			fmt.Printf("%s Erreur lors de l'exécution de %s\n", red("X"), cmd)
			fmt.Printf("Sortie de la commande : %s\n", out)
		} else {
			fmt.Printf("%s Commande exécutée avec succès : %s\n", green("✔"), cmd)
		}
	}
}

// Fonction pour gérer mysql_secure_installation avec interaction utilisateur
func SecureMySQL() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Veuillez entrer le mot de passe root MySQL: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	// Automatiser la configuration sécurisée de MySQL avec des réponses automatiques
	cmd := fmt.Sprintf(`echo -e "Y\n%s\n%s\nY\nY\nY\nY" | mysql_secure_installation`, password, password)
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()

	if err != nil {
		fmt.Printf("Erreur lors de la configuration sécurisée de MySQL: %v\n", err)
	} else {
		fmt.Println("Configuration sécurisée de MySQL terminée.")
		fmt.Printf("Sortie de la commande : %s\n", out)
	}
}
