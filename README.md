# clicord

A Discord terminal client. Heavily work-in-progress, expect breaking changes.

## Important Disclaimer

> [!CAUTION]
> Automated user accounts or "self-bots" are against Discord's Terms of Service. I am not
> responsible for any loss caused by using "self-bots" or clicord. This client is against the Terms
> of Service and makes little effort in hiding that it is a custom client, for example by not
> sending typing indicators or other data that the regular Discord client does.

## Installation

### Building from source

```bash
git clone https://github.com/sinjs/clicord.git
cd clicord
go build .
```

### Linux clipboard support

- `xclip` or `xsel` for X11 (`apt install xclip`)
- `wl-clipboard` for Wayland (`apt install wl-clipboard`)

## Usage

1. Run the `clicord` executable.

   > [!WARNING]  
   > It is safer to log in using an authentication token instead of using the UI, since Discord
   > monitors their authentication endpoints closely and thus increases your ban risk.
   >
   > Instead, to log in using a token, provide the `token` command-line flag to the executable
   > (eg: `--token "MTI2NDU5Nzk0NTY4OTU3NTU5MQ.VGVTdA.Z28-XdheSB-HVwaWQgcGV-29uCg"`). The token is
   > stored in the default keyring.

2. Enter your email and password (and 2fa code if required) and click on the "Login" button to
   continue.

## Configuration

The configuration file allows you to configure and customize the behavior, keybindings, and theme of
the application. The file is not created automatically, it uses the default config embedded into the
application. To configure clicord, create the config file at one of these locations:

- Linux: `$XDG_CONFIG_HOME/clicord/config.toml` or `$HOME/.config/clicord/config.toml`
- MacOS: `$HOME/Library/Application Support/clicord/config.toml`
- Windows: `%AppData%/clicord/config.toml`

```toml
mouse = true # Allows mouse usage

timestamps = true # If enabled, message timestamps will be displayed
timestamps_before_author = true # If enabled, timestamps will be displayed before the author name instead of directly after it
timestamps_format = "3:04PM" # The timestamp format should be one of these values: https://pkg.go.dev/time#pkg-constants

# These settings are used to emulate a different Discord client. See https://docs.discord.sex/reference#example-client-properties-(web) for example values for different platforms
user_agent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.3"
os = "Windows"
browser = "Chrome"
device = ""

messages_limit = 50 # The amount of messages to fetch in one channel
editor = "default" # Any executable for an editor, for example `nvim`. By default it uses the $EDITOR environment variable, or if not found, `vi`

# Keybindings: These keybinds are equivalent to the result of https://pkg.go.dev/github.com/gdamore/tcell#EventKey.Name
[keys]
focus_guilds_tree = "Ctrl+G"
focus_messages_text = "Ctrl+T"
focus_message_input = "Ctrl+P"
toggle_guild_tree = "Ctrl+B"
select_previous = "Rune[k]"
select_next = "Rune[j]"
select_first = "Rune[g]"
select_last = "Rune[G]"

[keys.guilds_tree]
select_current = "Enter"

[keys.messages_text]
select_reply = "Rune[s]"
reply = "Rune[r]"
reply_mention = "Rune[R]"
delete = "Rune[d]"
yank = "Rune[y]"
open = "Rune[o]"

[keys.message_input]
send = "Enter"
editor = "Ctrl+E"
cancel = "Esc"

# Themes can change the visuals of the client
[theme]
border = true
border_color = "default"
border_padding = [0, 0, 1, 1]
title_color = "default"
background_color = "default"

[theme.guilds_tree]
auto_expand_folders = true
graphics = true

[theme.messages_text]
author_color = "pink"
reply_indicator = "╭ "
```
