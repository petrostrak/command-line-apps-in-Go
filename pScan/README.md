#### pScan
pScan is a `CLI` tool that uses subcommands, similar to Git or Kubernetes. This tool executes a `TCP` port scan on a list of hosts similarly to the Nmap command. It allows you to add, list, and delete hosts from the list using the subcommand `hosts` . It executes the scan on selected ports using the subcommand `scan`. Users can specify the ports using a command-line flag. It also features command completion using the subcommand `completion` and manual page generation with the subcommand `docs` . Cobra helps you define the subcommand structure by associating these subcommands in a tree data structure.

The purpose of this application is to demonstrate how to use Cobra to help
you create command-line applications and use Go to create networking
applications. You can use this application to monitor your system, but
remember to never port scan systems you don’t own.