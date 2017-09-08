# dfs
change and save defaults settings for Mac.

```sh
$ brew install dfs
```

```sh
$ dfs update # update definitions
```

```sh
$ dfs set dock/hidden true # completions are shown from definitions
$ dfs set dock/static-only true
$ dfs set dock/autohide true
$ dfs set finder/AppleShowAllExtensions true
```

```sh
$ cat $HOME/.dfs/dfs.json
{
    "settings": {
        "dock": {
        "hidden": true,
        "static-only": true,
        "autohide": true
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
$ dfs export # export settings to defaults.sh
$ cat $HOME/.dfs/defaults.sh
# Dock
defaults write com.apple.Dock showhidden -bool true
defaults write com.apple.dock static-only -boolean true
defaults write com.apple.dock autohide -boolean true
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
$ cat $HOME/.dfs/10.12/dock.json
{
    "hidden": {
        "type": "bool"
    },
    "static-only": {
        "type": "boolean"
    },
    "autohide": {
        "type": "boolean"
    }
}
```