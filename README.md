# jump
Jump is a command-line tool for quickly navigating to directories you use often.

[![asciicast](https://asciinema.org/a/Q6EhmGOHmxSnGSYahelpTuEnh.svg)](https://asciinema.org/a/Q6EhmGOHmxSnGSYahelpTuEnh)

## Installation
Installation is a little bit trickier than it should be, as it is [very difficult and dodgy](https://stackoverflow.com/a/2375174/7409288) to change the directory of the parent process. For this reason, you will need to add some code to your `.bashrc` or `.zshrc`. The command `jump-config` does the heavily lifting, parsing and searching through the config file, and the code you add to your shell config file simply displays a prompt and takes you there.

Start off by installing the `jump-config` package using `go get`:

```bash
go get -u gitlab.com/ollybritton/jump/jump-config
```

You should now have access to the `jump-config` command in your shell, assuming you have `$GOPATH/bin` added to your `$PATH`.

Since this command won't actually be able to change the directory, we now need to add a bash/zsh function to your `.zshrc` or `.bashrc`.

### Automatic Setup
The following command

```bash
jump-config init --language [your shell language, default zsh] [path to rc file]
```

Will tack a `jump ()` function onto the end of your file. To activate these changes, you need to `source` the file you specified. Please run `jump-config init --help` for more options, such as enabling and disabling color output.

### Manual Setup
If you want to manually add the function yourself, you can use the command

```bash
jump-config func --language [your shell language, default zsh]
```

which will print the code for the function. You could then do

```bash
jump-config func --language zsh >> ~/.zshrc
```

to automatically append the function yourself.

That's it! You should now be able to use the `jump` command.

## Usage
You use the `jump` command like this:

```bash
jump [alias]
```

where `alias` is an alias you have defined in the `~/.jump-config.yaml` file. An alias is just a short-hand code for a longer directory. An example of a valid `.jump-config.yaml` file would be:

```yaml
aliases:
    code: " ~/Documents/Code"
    go: "~/go/src/gitlab.com/ollybritton"
    python: "~/Documents/Code/Python"
```

It does not need to match the alias fully in order to know where to jump to. For example, say you have an alias named `python` which points to `~/Documents/Code/Python`. You need not do this every time:

```bash
jump python
```

You can instead shorten it to

```bash
jump py
```

If there are multiple aliases that begin with those letters, it will prompt you to select the correct one.

Once you have defined at least one alias, you can then use the short command `jump-config add [name][alias]` to quickly add an alias. For example

```bash
jump-config add root ~
```

Will add the code

```yaml
    root: "~"
```

