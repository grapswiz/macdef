# dfs
change and save defaults settings for Mac.

```sh
$ brew install dfs
```

```sh
$ dfs install dock/hidden true
$ dfs install dock/static-only true
$ dfs install dock/autohide true
```

```sh
$ cat $HOME/.dfs/dfs.json
{
  "installed": {
    "dock/hidden": true,
    "dock/static-only": true,
    "dock/autohide": true
  }
}
```

```sh
$ dfs apply
```

```sh
$ dfs export
$ cat $HOME/.dfs/defaults.sh
# Dock
defaults write com.apple.Dock showhidden -bool YES
defaults write com.apple.dock static-only -boolean true
defaults write com.apple.dock autohide -boolean true
```
