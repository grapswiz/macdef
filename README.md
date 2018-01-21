# macdef

[![CircleCI](https://circleci.com/gh/grapswiz/macdef.svg?style=svg)](https://circleci.com/gh/grapswiz/macdef)

change and save defaults settings for Mac.

// under development...

```sh
$ brew install macdef
```

```sh
$ macdef completion (bash or zsh) > ? #TODO
```

```sh
$ macdef update # update definitions in .macdef/definitions/
```

```sh
$ macdef set dock/showhidden true # completions are shown from definitions
$ macdef set dock/static-only true
$ macdef set dock/autohide true
$ macdef set dock/orientation bottom
$ macdef set finder/AppleShowAllExtensions true
```

```sh
$ cat $HOME/.macdef/macdef.json
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
$ macdef apply # apply changes to your Mac using macdef.json
```

```sh
$ macdef export # show settings using shell script
# Dock
defaults write com.apple.Dock showhidden -bool true # dock/showhidden
defaults write com.apple.Dock static-only -boolean true # dock/static-only
defaults write com.apple.Dock autohide -boolean true # dock/autohide
defaults write com.apple.dock orientation bottom

# Finder
defaults write NSGlobalDomain AppleShowAllExtensions -bool true
```

```
.macdef
├── definitions
│   ├── 10.11
│   │   └── dock.json
│   └── 10.12
│       ├── dock.json
│       └── finder.json
└── macdef.json
```

```sh
$ cat $HOME/.macdef/definitions/10.12/dock.json
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
