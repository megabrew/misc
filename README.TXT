Compiling/installing
--------------------
These tools can be installed together by using "go install github.com/megabrew/misc/...@latest", or individually by using the command with the tool's name in place of "...". Installing them in this fashion will build and place the executables in the user's "go/bin" directory to be used systemwide across projects. If it's necessary to use these tools in local project directories, it should be straight-forward to compile them individually using "go build (tool).go". Alternatively, clone the whole repo and run "go build -o . ./..." (once again, to build an individual tool replace "..." with the name of the subdirectory containing it).

The Go toolchain (https://go.dev/dl/) must be installed on your system; on Windows, the necessary directories will be added to your PATH automatically.
