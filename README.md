# resource-exhaustion-talk

## Topics:

### Examples of (Shared) Resource Exhaustion:
1. Open file handles `sudo lsof -p <pid>`
2. Thread `ps -o nlwp <pid>`
3. IO/Network Saturation
4. Shared Application (database, shared singleton object, cloud service, etc)
5. Not sure: Try looking at syscalls `strace -fp <pid>`

### Strategies To Prevent Resource Exhaustion:
1. Ask yourself what are the resources your program depend on. If you are not sure, run a simple load test to gain an intuition.
2. Try to have your solution model the resource constraints intrinsic to your problem.
3. If your program can run on a large, or potentially unbounded input, have plan for how the program will behave.
    1. Backpressure (i.e. don't take on new work until you have the resources to do so):
        1. Worker pools
        2. Timeouts
    2. Shed load (i.e. drop old work)
        1. Staleness checks
    3. Terminate with an informative error message
