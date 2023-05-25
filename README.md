# kummy - Kubernetes for Dummies 

kummy a tool targeting developer without or little kubernetes knowledge so that they can inspect their applications.

No Pun intended 🤔.

**Important Note:** [WIP] Currently not in a working state. Commands or scope may change.

## :question: What is kummy?
Learning Kuberne tes is hard. If you find yourself in a Team with developers or other people, with no skills in 
kubernetes, you can provide them with this tool. With kummy, they can handle basic operations, like viewing logs, 
restarting pod or finding an overview on how the application is configured.

### What it isn't: 
A replacement for [kubectl](https://kubernetes.io/de/docs/reference/kubectl/). If you want to dive deeper into working 
with kubernetes, you have to learn the basic and work with kubectl, or one of the other tools out there.


## :hammer_and_wrench: Installation

### From source

Run 
```shell
git clone https://github.com/tomschinelli/kummy.git
cd kummy
./install.sh
```

### Binaries
ToDo Build pipeline

### AUR
ToDo publish package

### Completion
It should work like [the completion for kubectl](https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-autocomplete), since I use the same packages to achieve it.

For more information see [Completion](./docs/completion.md)

## :green_book: How to use 

Get application overview: 
```shell
kummy overview
```

For a more detailed documentation, see the [user guide](./docs/user-guide.md)