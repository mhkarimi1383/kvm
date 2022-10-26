# kvm
Easily switch between kubectl version

## Why?
Managing kubectl versions can help you when managing multiple clusters with different versions and can make you fully compatible with your clusters.

## Install

Make required paths ready by running
```shell
mkdir -p "$HOME/.kvm/bin/"
```

and put the following line into your `.bashrc` or `.zshrc`

```shell
export PATH="$HOME/.kvm/bin:$PATH"
```

then grab the package binary and put it into a dir that is covered by `$PATH`