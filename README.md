# jump
Jump is a command-line tool for quickly navigating to directories you use often.

[![asciicast](https://asciinema.org/a/Q6EhmGOHmxSnGSYahelpTuEnh.svg)](https://asciinema.org/a/Q6EhmGOHmxSnGSYahelpTuEnh)

## Installation
Installation is a little bit trickier than it should be, as it is [very difficult and dodgy](https://stackoverflow.com/a/2375174/7409288) to change the directory of the parent process. Start off by installing the `jump-config` package using `go get`:

```bash
go get -u gitlab.com/ollybritton/jump/jump-config
```

You should now have access to the `jump-config` command in your shell, assuming you have `$GOPATH/bin` added to your `$PATH`.

Since this command won't actually be able to change the directory, we now need to add a bash/zsh function to your `.zshrc` or `.bashrc`. To do this, you can either call

```bash
jump-config init --language [your language, such as zsh] [path to rc file]
```

This will tack a `jump ()` function onto the end of your file. To activate these changes, you need to `source` the file you specified.

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

## To Add

* Figure out how to actually format Go code - it looks like utter rubbish at the moment.
* I think that the bash.sh doesn't actually work, which is a bit useless.
* Add more options for the init phase, like choosing if they want colors.
