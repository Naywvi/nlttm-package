package _args

import (
	"fmt"
	"os/exec"
	"strings"
)

// Fonction pour installer un package
func Install(packageName string, commands []string) {
	fmt.Printf("Installation du package %s\n", packageName)

	// Combine toutes les commandes en une seule chaîne
	fullCommand := strings.Join(commands, " ")

	// Exécuter la commande complète via "sh -c"
	fmt.Printf("Exécution de : %s\n", fullCommand)
	out, err := exec.Command("sh", "-c", fullCommand).CombinedOutput()
	if err != nil {
		fmt.Printf("Erreur lors de l'exécution de %s : %v\n", fullCommand, err)
		return
	}
	fmt.Printf("Résultat : %s\n", out)
}
