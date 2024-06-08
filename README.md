# Xpass password manager
A simple password manager I've written for personal use in GO. Terminal only, no central servers, total encryption. Please note, if you lose your passphrase you will lose access to your passwords (i.e. unrecoverable!).

## Usage
If you want to use it, then you can run the program like any other GO program. Note, you have to set the env variable XPASS\_LOCKER to have a folder path leading to where you want to store your locker. After doing that, you can manage your passwords with the following commands:  
 - init _"lockername"_: creates a new locker (.aes) file to store passwords in.
 - ls _"optional: lockername"_: See the lockers or password ids in a lockerfile.
 - unlock _"lockername"_: Prompts a secure input to unlock a lockername. 
 - add _"lockername"_ _"password id"_: Adds a username/password combination in a locker with a password id.
 - get _"lockername"_ _"password id"_: Gets a username/password combination and returns them to the standard output.
 - cat _"lockername"_: Print the raw locker contents.

## Next steps
 - Add password to clipboard.
