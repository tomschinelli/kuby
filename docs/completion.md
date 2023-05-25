# Kuby - Completion

## BASH
```shell
source <(kuby completion bash) # set up autocomplete in bash into the current shell, bash-completion package should be installed first.
echo "source <(kuby completion bash)" >> ~/.bashrc 
```

## ZSH
```shell
source <(kuby completion zsh)  # set up autocomplete in zsh into the current shell
echo '[[ $commands[kuby] ]] && source <(kuby completion zsh)' >> ~/.zshrc # add autocomplete permanently to your zsh shell

```
