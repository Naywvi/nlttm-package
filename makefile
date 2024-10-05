# Variables
APP_NAME = nlttm
VERSION = 1.0.0
BUILD_DIR = build
DEB_DIR = debian
CONF_DIR = /etc/nlttm

# Couleurs ANSI
RED = \033[31m
GREEN = \033[32m
YELLOW = \033[33m
BLUE = \033[34m
RESET = \033[0m

# Cibles par défaut
.PHONY: all build clean deb fix-control

all: build deb

# Simulation de chargement
loading:
	@i=0; \
	while [ $$i -lt 10 ]; do \
		printf "$(YELLOW)Loading... %s\r$(RESET)" $$i; \
		sleep 0.1; \
		i=$$(($$i + 1)); \
	done

# Correction automatique du fichier control
fix-control:
	@echo "==> Correction du fichier control"
	@sed -i '/^$$/d' debian/DEBIAN/control
	@echo "" >> debian/DEBIAN/control  # Ajoute une ligne vide à la fin

# Compilation du programme Go
build:
	@echo "$(BLUE)==> Compilation de $(APP_NAME)...$(RESET)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)/
	@echo "$(GREEN)==> Compilation terminée$(RESET)"

# Copie des fichiers de configuration et du binaire
install:
	@echo "$(BLUE)==> Création de la structure /etc/nlttm...$(RESET)"
	sudo mkdir -p $(CONF_DIR)/config/packages
	sudo mkdir -p $(CONF_DIR)/debian

	@echo "$(BLUE)==> Copie des fichiers de configuration dans $(CONF_DIR)...$(RESET)"
	sudo cp config/nlttm.conf $(CONF_DIR)/nlttm.conf
	sudo cp -r config/packages/* $(CONF_DIR)/config/packages/ # Copier tous les fichiers de packages

	@echo "$(GREEN)==> Fichiers de configuration copiés avec succès$(RESET)"

	@echo "$(BLUE)==> Vérification et création de /usr/local/bin si nécessaire...$(RESET)"
	sudo mkdir -p /usr/local/bin

	@echo "$(BLUE)==> Installation du binaire dans /usr/local/bin...$(RESET)"
	sudo cp $(BUILD_DIR)/$(APP_NAME) /usr/local/bin/$(APP_NAME)
	@echo "$(GREEN)==> Binaire installé avec succès$(RESET)"

# Création de l'arborescence Debian et du package .deb
deb: build install fix-control loading
	@echo "$(BLUE)==> Création de l'arborescence Debian...$(RESET)"
	mkdir -p $(DEB_DIR)/usr/local/bin
	cp $(BUILD_DIR)/$(APP_NAME) $(DEB_DIR)/usr/local/bin/

	mkdir -p $(DEB_DIR)/etc/nlttm/config/packages
	cp config/nlttm.conf $(DEB_DIR)/etc/nlttm/nlttm.conf
	cp -r config/packages/* $(DEB_DIR)/etc/nlttm/config/packages/ # Copier tous les fichiers de packages

	@echo "$(BLUE)==> Construction du fichier .deb...$(RESET)"
	dpkg-deb --build $(DEB_DIR) $(APP_NAME)_$(VERSION)_amd64.deb
	@echo "$(GREEN)==> Package .deb créé : $(APP_NAME)_$(VERSION)_amd64.deb$(RESET)"
	@echo "$(BLUE)==> Installation de $(APP_NAME)...$(RESET)"
	sudo dpkg -i $(APP_NAME)_$(VERSION)_amd64.deb
	@echo "$(GREEN)==> Installation terminée$(RESET)"

# Nettoyer les fichiers compilés
clean:
	@echo "$(RED)==> Nettoyage des fichiers...$(RESET)"
	rm -rf $(BUILD_DIR)
	rm -f $(APP_NAME)_$(VERSION)_amd64.deb
	@echo "$(GREEN)==> Fichiers nettoyés$(RESET)"
