# REX 

Rex is a handy tool to automate the remote execution of commands over SSH.
It offers single host execution or group execution.
After successfull execution it will return an output.

An example config can be found in the repo at cmd/config.yml.dist. 
Please copy it and rename it to ~/.rex/config.yml. 

By default the program expects the config to be in the ~/.rex/config.yml. The path can be modified with -c flag.


## Single and Group Command Execution

Use the flag -s for single target

```bash
./rex -s myserver ls -la
```

Use the flag -g for multiple targets

```bash
./rex -g mygroup ls -la
```

## Templates

You can setup templates and execute then with : as prefix.
The args after the template name can be insert  in the command template with ${{0}}. 
The 0 would be the first argument after the command template name.



Example

```bash
./rex -s myserver :hello_world arg1 arg2  

#Example

./rex -s myserver :hello_world test hallo  
```


```yml
templates:
  update:
    cmd: sudo apt-get update -y && sudo apt-get dist-upgrade -y
  hello_world:
    cmd: echo "hello ${{0}} ${{1}} world"
```

