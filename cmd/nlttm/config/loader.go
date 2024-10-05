package _config

import (
	"bufio"
	"os"
	"strings"
)

// Fonction pour lire le fichier de configuration
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

// Fonction pour lire le fichier des packages
func LoadPackagesConfig(filename string) (map[string][]string, error) {
	packages := make(map[string][]string)

	file, err := os.Open(filename)
	if err != nil {
		return packages, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 && !strings.HasPrefix(line, "#") { // Ignorer les lignes vides ou commentaires
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := parts[0]
				values := strings.Split(parts[1], " ")
				packages[key] = values
			}
		}
	}

	return packages, scanner.Err()
}
