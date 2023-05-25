# kummy - Completion

## BASH
```shell
source <(kummy completion bash) # set up autocomplete in bash into the current shell, bash-completion package should be installed first.
echo "source <(kummy completion bash)" >> ~/.bashrc 
```

## ZSH
```shell
source <(kummy completion zsh)  # set up autocomplete in zsh into the current shell
echo '[[ $commands[kummy] ]] && source <(kummy completion zsh)' >> ~/.zshrc # add autocomplete permanently to your zsh shell

```
