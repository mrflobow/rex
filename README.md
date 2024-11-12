# Remote Execution Automation

This is a handy tool to automate the remote execution of commands via ssh.
It offers a command templating to quicker execute commands


For example config look at cmd/config.yml.dist

## Templates

You can setup templates and execute then with : as prefix.
The args after the template name can be insert  in the command template with ${{0}}. 
The 0 would be the first argument after the command template name.



Example

```bash
./rex -server myserver :hello_world arg1 arg2  

#Example

./rex -server myserver :hello_world test hallo  
```


```yml
templates:
  update:
    cmd: sudo apt-get update -y && sudo apt-get dist-upgrade -y
  hello_world:
    cmd: echo "hello ${{0}} ${{1}} world"
```

