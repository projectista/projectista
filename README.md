<p align="center">
  <img src="https://github.com/projectista/static/blob/main/logo/projectista@2x.png">
</p>

---

# What is Projectista?

Projectista is a single binary written in go that helps developers
bootstrap a new application in their favorite programming language / framework.

# How to use it

Download the `pi` binary from the `Releases` section on `GitHub` and place it in your path.

You can proceed with the scaffolding.

This version only supports php packages, others options will be added.

**Known issues**:

On `MacOS`, if you downloaded the binary using a browser, you need to remove the quarantine flag from the binary before being able to execute it:

```bash
xattr -d com.apple.quarantine pi
```

# Download from command line

To download the latest version of the `pi` executable, using command line, you can use the following command (Replace the file name with the correct one for your architecture / operating system):

```bash
curl -s -L https://github.com/projectista/projectista/releases/latest/download/projectista_Darwin_arm64.tar.gz | gunzip -c - | tar xopf - pi
```

This command will place the `pi` executable in the current folder, you are free to move it in your path.

## Laravel Application

To scaffold a new Laravel Application:

```bash
pi laravel application myapplication
```

The project will be scaffolded in the current directory. To change directory use the provided flag:

```bash
pi laravel application myapplication --folder="myapplication"
```

**The folder must exists.**

The command provides other flags to specify:

- **Author** of the project,
- **Email** of the author of the project,
- **Vendor** of the project,
- **Description** of the project.

You can have more information using the bundled help:

```bash
pi laravel application --help
```

Some examples:

```bash
pi laravel application myapplication --author="Vincenzo Petrucci" --vendor="illegal studio"
```

```bash
pi laravel application myapplication --description="My new Laravel application"
```

## Laravel package

To scaffold a new Laravel Package:

```bash
pi laravel package mypackage
```

The package will be scaffolded in the current directory. To change directory use the provided flag:

```bash
pi laravel package mypackage --folder="myapplication"
```

**The folder must exists.**

The command provides other flags to specify:

- **Author** of the project,
- **Email** of the author of the project,
- **Vendor** of the project,
- **Description** of the project.

You can have more information using the bundled help:

```bash
pi laravel package --help
```

Some examples:

```bash
pi laravel package mypackage --author="Vincenzo Petrucci" --vendor="illegal studio"
```

```bash
pi laravel package mypackage --description="My new Laravel package"
```

## PHP package

To scaffold a new project:

```bash
pi php package myproject
```

The project will be scaffolded in the current directory. To change directory use the provided flag:

```bash
pi php package myproject --folder="myproject"
```

**The folder must exists.**

The command provides other flags to specify:

- **Author** of the project,
- **Email** of the author of the project,
- **Vendor** of the project,
- **Description** of the project.

You can have more information using the bundled help:

```bash
pi php package --help
```

Some examples:

```bash
pi php package mypackage --author="Vincenzo Petrucci" --vendor="illegal studio"
```

```bash
pi php package mypackage --description="My awesome project"
```

# Build on your machine

If you want to build the project on your machine you have to first clone it:

```bash
git clone --recurse-submodules git@github.com:projectista/projectista.git
```

Than you can proceed to the build stage:

```bash
cd projectista
go build
```

The binary `projectista` will be available in the root of the project.
Rename to `pi` if you whish, and place it in your path.

# Sponsor

This project is sponsored by [illegal studio](https://illegal.studio)
