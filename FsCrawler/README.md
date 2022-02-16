#### File System Crawler
The fsClawler tool has two main goals: descending into a directory tree to look
for files that match a specified criteria and executing an action on these files.

It accepts these command-line parameters:

* `-root`: The root of the directory tree to start the search. The default is the
current directory.
* `-list`: List files found by the tool. When specified, no other actions will
be executed.
* `-ext`: File extension to search. When specified, the tool will only match
files with this extension.
* `-size`: Minimum file size in bytes. When specified, the tool will only
match files whose size is larger than this value.
* `-del`: Adding the ability to delete
the files it finds.
* `-archive`: Before deleting files that are consuming too much space, you might want to
back them up in a compressed form so you can keep them around in case
you need them later.