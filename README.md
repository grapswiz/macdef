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
$ cat $HOME/.macdef/macdef.yaml
---
items:
- name: showhidden
  category: dock
  value: true
- name: orientation
  category: dock
  value: bottom
```

```sh
$ macdef apply # apply changes to your Mac using macdef.yaml
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
│   │   └── dock.yaml
│   └── 10.12
│       ├── dock.yaml
│       └── finder.yaml
└── macdef.yaml
```

```sh
$ cat $HOME/.macdef/definitions/10.12/dock.yaml
---
items:
- name: showhidden
  description: Show hidden apps
  type: bool
  commands:
  - defaults write com.apple.Dock showhidden -bool {{0}}
- name: static-only
  description: Only show applications that are running
  type: boolean
  commands:
  - defaults write com.apple.Dock static-only -boolean {{0}}
- name: autohide
  description: Hide Dock
  type: boolean
  commands:
  - defaults write com.apple.Dock autohide -boolean {{0}}
- name: orientation
  description: Dock position
  type: string literal
  values:
  - bottom
  - left
  commands:
  - defaults write com.apple.dock orientation {{0}}
```

## Install
1. `brew tap macdef`
1. `brew install macdef`
