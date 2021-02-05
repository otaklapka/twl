# Terminal Work Log
A trivial CLI app to log work.
![screen](https://i.imgur.com/39ziNV7.png)

The purpose of this app is to store quick notes about work during the day as easy as possible.
At the end of the day you could estimate the spent time and track it to right tasks in your time tracking tool.
The other day it will serve as list for standup.


##Subcommands
* `ls` lists records, default is today
    * `[-d <dd.mm.yyyy>]` date to list records of
    * `[-l]` list last day containing some records
* `log <message>` logs new message with current timestamp
    * `[-t <dd.mm.yyyy hh:mm>]` custom time
* `set <id>` update record
    * `[-m <new message>]` update message
    * `[-t <new time>]` update time
* `rm <id>` remove record
* `tdiff <id|hh:mm> <id|hh:mm>` count time diff between records or time

##Usage
* Download binary of desired platform
* Make it executable
* Add to PATH or alias

The app will create database file next to the executable