#### colStats
The program will receive two optional input parameters each
with a default value:

* `col`: The column on which to execute the operation. It defaults to 1.
* `op`: The operation to execute on the selected column. Initially, this tool
will support two operations: `sum` , which calculates the sum of all values
in the column, and `avg` , which determines the average value of the
column.

In addition to the two optional flags, this tool accepts any number of file
names to process. If the user provides more than one file name, the tool
combines the results for the same column in all files.