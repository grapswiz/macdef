# dfs
change and save defaults settings for Mac.

// under development...

```sh
$ brew install dfs
```

```sh
$ dfs completion (bash or zsh) > ? #TODO
```

```sh
$ dfs update # update definitions in .dfs/definitions/
```

```sh
$ dfs set dock/showhidden true # completions are shown from definitions
$ dfs set dock/static-only true
$ dfs set dock/autohide true
$ dfs set dock/orientation bottom
$ dfs set finder/AppleShowAllExtensions true
```

```sh
$ cat $HOME/.dfs/dfs.json
{
    "settings": {
        "dock": {
            "showhidden": true,
            "static-only": true,
            "autohide": true,
            "orientation": "bottom"
        },
        "finder": {
            "AppleShowAllExtensions": true
        }
    }
}
```

```sh
$ dfs apply # apply changes to your Mac using dfs.json
```

```sh
$ dfs export # show settings using shell script
# Dock
defaults write com.apple.Dock showhidden -bool true # dock/showhidden
defaults write com.apple.Dock static-only -boolean true # dock/static-only
defaults write com.apple.Dock autohide -boolean true # dock/autohide
defaults write com.apple.dock orientation bottom

# Finder
defaults write NSGlobalDomain AppleShowAllExtensions -bool true
```

```
.dfs
├── definitions
│   ├── 10.11
│   │   └── dock.json
│   └── 10.12
│       ├── dock.json
│       └── finder.json
└── dfs.json
```

```sh
$ cat $HOME/.dfs/definitions/10.12/dock.json
{
    "items": {
        "showhidden": {
            "type": "bool",
            "description": "Show hidden apps",
            "commands": ["defaults write com.apple.Dock showhidden -bool {{0}}"]
        },
        "static-only": {
            "type": "boolean",
            "description": "Only show applications that are running",
            "commands": ["defaults write com.apple.Dock static-only -boolean {{0}}"]
        },
        "autohide": {
            "type": "boolean",
            "description": "Hide Dock",
            "commands": ["defaults write com.apple.Dock autohide -boolean {{0}}"]
        },
        "orientation": {
            "type": "string literal",
            "values": ["bottom", "left"],
            "description": "Dock position",
            "commands": ["defaults write com.apple.dock orientation {{0}}"]
        }
    }
}
```