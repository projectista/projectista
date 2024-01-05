<p align="center">
  <img src="https://github.com/projectista/static/blob/main/logo/projectista.png">
</p>

# What is Projectista?

Projectista is a single binary written in go that helps developers
bootstrap a new application in their favorite programming language / framework.

# How to use it

Download the binary from the `Releases` section on `GitHub` and place it in your path. 

You can proceed with the scaffolding. 

This version only supports php packages, others options will be added.

## PHP package

To scaffold a new project: 

```bash
projectista php package myproject
```

The project will be scaffolded in the current directory. To change directory use the provided flag:

```bash
projectista php package myproject --folder="myproject"
```

**The folder must exists.**


# Sponsor

This project is sponsored by [illegal studio](https://illegal.studio)