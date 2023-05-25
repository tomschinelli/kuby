# Kuby - Kubernetes for Dummies 

Kuby a tool targeting developer without or little kubernetes knowledge so that they can inspect their applications.

**Important Note:** [WIP] Currently not in a working state. Commands or scope may change.

## What is kuby?

Learning Kubernetes is hard. If you find yourself in a Team with developers or other people, with no skills in 
kubernetes, you can provide them with this tool. With kuby, they can handle basic operations, like viewing logs, 
restarting pod or finding an overview on how the application is configured.

### What it isn't: 
A replacement for [kubectl](https://kubernetes.io/de/docs/reference/kubectl/). If you want to dive deeper into working 
with kubernetes, you have to learn the basic and work with kubectl, or one of the other tools out there.


## Installation
ToDo


### Completion
It should work like [the completion for kubectl](https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-autocomplete), since I use the same packages to achieve it.

#### BASH
```shell
source <(kuby completion bash) # set up autocomplete in bash into the current shell, bash-completion package should be installed first.
echo "source <(kuby completion bash)" >> ~/.bashrc 
```

### ZSH
```shell
source <(kuby completion zsh)  # set up autocomplete in zsh into the current shell
echo '[[ $commands[kuby] ]] && source <(kuby completion zsh)' >> ~/.zshrc # add autocomplete permanently to your zsh shell

```


## Usage 

### Get application overview
```shell
kuby get apps
```

### Accessing logs
```shell
kuby logs <application-name>
```

