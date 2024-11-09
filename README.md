# Installation
simply run `go build -o ignore_tool main.go` to create the go binary. Then depending on your os added it to path.
- `sudo cp ignore_tool /usr/local/bin`
- `sudo chmod +x /usr/local/bin/my_app`

# Usage
- `--list` to get the supported languages
- `<lang> > .gitignore` to create the `.gitignore`
- `<lang>` print the content that will be stored in the `.gitignore`

# Example

```Bash
  # list all supported languages
  ignore_tool --list

  # create a .gitingore file for go
  ignore_tool Go > .gitignore
```

  
