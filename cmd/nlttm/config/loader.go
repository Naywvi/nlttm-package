package _config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Fonction pour lire le fichier de configuration générique
func LoadConfig(filename string) (map[string]string, error) {
	config := make(map[string]string)

	file, err := os.Open(filename)
	if err != nil {
		return config, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") { // Ignorer les lignes vides ou commentaires
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				config[parts[0]] = strings.TrimSpace(parts[1])
			}
		}
	}

	return config, scanner.Err()
}

// Fonction pour charger les commandes d'une section spécifique du fichier de configuration du package
func LoadPackageSectionCommands(packageName string, section string) ([]string, error) {
	// Construire le chemin absolu vers le fichier de configuration du package dans /etc/nlttm/config/packages/
	filename := fmt.Sprintf("/etc/nlttm/config/packages/%s.conf", packageName)

	// Vérifier si le fichier existe
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Le fichier de configuration pour le package %s est introuvable : %v", packageName, err)
	}
	defer file.Close()

	// Lire les commandes de la section spécifiée
	var commands []string
	var inSection bool
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			// Détecter le début d'une section
			inSection = strings.Trim(line, "[]") == section
			continue
		}

		// Ajouter les commandes à partir de la section courante
		if inSection && len(line) > 0 && !strings.HasPrefix(line, "#") { // Ignorer les lignes vides ou les commentaires
			commands = append(commands, line)
		}
	}

	if len(commands) == 0 {
		return nil, fmt.Errorf("Section %s introuvable ou vide dans %s.conf", section, packageName)
	}

	return commands, scanner.Err()
}
