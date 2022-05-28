# md-codedocs

## Usage
> This cli tool will parse your markdown document and extract all the codeblocks in it

md-codedocs

## Flags
|Flag|Usage|
|----|-----|
|`--debug`|enable debug messages|
|`--disable-update-checks`|disables update checks|
|`--raw`|print unstyled raw output (set it if output is written to a file)|

## Commands
|Command|Usage|
|-------|-----|
|`md-codedocs completion`|Generate the autocompletion script for the specified shell|
|`md-codedocs help`|Help about any command|
# ... completion
`md-codedocs completion`

## Usage
> Generate the autocompletion script for the specified shell

md-codedocs completion

## Description

```
Generate the autocompletion script for md-codedocs for the specified shell.
See each sub-command's help for details on how to use the generated script.

```

## Commands
|Command|Usage|
|-------|-----|
|`md-codedocs completion bash`|Generate the autocompletion script for bash|
|`md-codedocs completion fish`|Generate the autocompletion script for fish|
|`md-codedocs completion powershell`|Generate the autocompletion script for powershell|
|`md-codedocs completion zsh`|Generate the autocompletion script for zsh|
# ... completion bash
`md-codedocs completion bash`

## Usage
> Generate the autocompletion script for bash

md-codedocs completion bash

## Description

```
Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(md-codedocs completion bash)

To load completions for every new session, execute once:

#### Linux:

	md-codedocs completion bash > /etc/bash_completion.d/md-codedocs

#### macOS:

	md-codedocs completion bash > /usr/local/etc/bash_completion.d/md-codedocs

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion fish
`md-codedocs completion fish`

## Usage
> Generate the autocompletion script for fish

md-codedocs completion fish

## Description

```
Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	md-codedocs completion fish | source

To load completions for every new session, execute once:

	md-codedocs completion fish > ~/.config/fish/completions/md-codedocs.fish

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion powershell
`md-codedocs completion powershell`

## Usage
> Generate the autocompletion script for powershell

md-codedocs completion powershell

## Description

```
Generate the autocompletion script for powershell.

To load completions in your current shell session:

	md-codedocs completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... completion zsh
`md-codedocs completion zsh`

## Usage
> Generate the autocompletion script for zsh

md-codedocs completion zsh

## Description

```
Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions for every new session, execute once:

#### Linux:

	md-codedocs completion zsh > "${fpath[1]}/_md-codedocs"

#### macOS:

	md-codedocs completion zsh > /usr/local/share/zsh/site-functions/_md-codedocs

You will need to start a new shell for this setup to take effect.

```

## Flags
|Flag|Usage|
|----|-----|
|`--no-descriptions`|disable completion descriptions|
# ... help
`md-codedocs help`

## Usage
> Help about any command

md-codedocs help [command]

## Description

```
Help provides help for any command in the application.
Simply type md-codedocs help [path to command] for full details.
```


---
> **Documentation automatically generated with [PTerm](https://github.com/pterm/cli-template) on 28 May 2022**
