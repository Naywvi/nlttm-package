# Variables
APP_NAME = nlttm
VERSION = 1.0.0
BUILD_DIR = build
DEB_DIR = debian

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
	@echo "$(BLUE)==> Compilation de $(APP_NAME) pour $(GOOS)...$(RESET)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)/
	@echo "$(GREEN)==> Compilation terminée$(RESET)"

# Création de l'arborescence Debian et du package .deb
deb: build fix-control loading
	@echo "$(BLUE)==> Création de l'arborescence Debian...$(RESET)"
	mkdir -p $(DEB_DIR)/usr/local/bin
	cp $(BUILD_DIR)/$(APP_NAME) $(DEB_DIR)/usr/local/bin/

	mkdir -p $(DEB_DIR)/etc/$(APP_NAME)
	cp config/$(APP_NAME).conf $(DEB_DIR)/etc/$(APP_NAME)/

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
