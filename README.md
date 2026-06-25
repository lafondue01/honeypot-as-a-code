# HoneyPot-as-a-Code

The goal is only to print the time and the IP address of the attacker when they try to connect to the server.

## Execute the code

1. cd /honey-pot-as-a-code/src
2. got mod init honeypot-as-a-code/main
3. go mod tidy
4. go mod run .

Then go on an other terminal and execute:

5. telnet <@IP> <port>
Finally write a message and send it.

The connection will close and the server will print the message from the attacker.
