# Loganalyzer

Loganalyzer est un outil en ligne de commande développé en Go, qui permet d’analyser des fichiers de logs (access, erreur, etc.) et d’en extraire des résultats consolidés dans un fichier JSON.

## Fonctionnement

Le programme lit un fichier de configuration au format JSON contenant une liste de logs à analyser. Chaque log est traité de manière concurrente. L’outil vérifie notamment :
- Si le fichier existe (`ErrFileNotFound`)
- Si le fichier est vide (`ErrFileEmpty`)

Un rapport d’analyse est ensuite exporté dans un fichier JSON avec les résultats de chaque log : succès, erreur de parsing, ou échec.

## Structure du projet

```
.
├── cmd/               # Commandes CLI (Cobra)
│   ├── analyze.go     # Commande `analyze`
│   └── root.go        # Point d’entrée CLI
├── internal/
│   ├── analyzer/
│   │   ├── analyzer.go     # Analyse concurrente des logs
│   │   ├── reporter.go     # Export JSON
│   │   └── errors.go       # Types d'erreurs customisées
│   └── config/
│       └── config.go       # Lecture du fichier de configuration
├── main.go            # Lancement de l’application
├── access.log         # Exemple de fichier de log valide
├── errors.log         # Exemple de log avec des erreurs
├── corrupted.log      # Exemple de log corrompu
├── config.json        # Exemple de fichier de configuration (à créer)
└── README.md
```

## Installation

1. Cloner le dépôt :
   ```bash
   git clone <url-du-depot>
   cd loganalyzer
   ```

2. Construire le binaire :
   ```bash
   go build -o loganalyzer
   ```

## Utilisation

```bash
./loganalyzer analyze --input config.json --output result.json
```

### Paramètres

- `--input` ou `-i` : Chemin vers le fichier de configuration JSON listant les logs à analyser (obligatoire).
- `--output` ou `-o` : Chemin du fichier JSON où seront exportés les résultats (par défaut : `result.json`).

### Exemple de `config.json`

```json
[
  {
    "id": "access",
    "path": "access.log",
    "type": "access"
  },
  {
    "id": "errors",
    "path": "errors.log",
    "type": "error"
  },
  {
    "id": "corrupted",
    "path": "corrupted.log",
    "type": "unknown"
  }
]
```

## Résultat attendu (`result.json`)

```json
[
  {
    "log_id": "access",
    "file_path": "access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "errors",
    "file_path": "errors.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "corrupted",
    "file_path": "corrupted.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  }
]
```

Si un fichier est introuvable ou vide, `status` sera `"FAILED"` et `error_details` contiendra le message correspondant.

## Gestion des erreurs

Le programme gère plusieurs cas :
- Fichier introuvable : message `"Fichier introuvable."`
- Fichier vide : message `"Erreur de parsing."`
- Autres erreurs : message `"Erreur inconnue."`

## Dépendances

- Cobra : pour la gestion de la CLI

## Auteurs

Projet réalisé dans le cadre d’un TD de Go en groupe — EFREI 2025.
