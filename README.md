# DoubleBooked

Given a csv input of appointments with a start time and an end time, this program returns a csv output of pairs of Appointments which have a conflict.

## Getting Started

Clone the repository locally. Within the workspace, input the following command:
```
go install driver
```

cd to the /bin directory, and execute the driver:

```
cd bin
./driver
```
(Note that the current state of the product uses relative paths, so running this from a separate directory may cause filename errors when fetching input files)

Look in the workspace/file directory for a new csv file containing all pairs of conflicting appointments!

## Config

Take a look at the config file, and make any necessary changes to the input file name, output file name, or time format of the input file. For more information on time formatting in Golang, check Go's documentation on [the Time package](https://golang.org/pkg/time/#pkg-constants).

