# Golang + spawned Windows docker container

This demo aims to demonstrate that windows docker container spawned as part of the same console process
group of another process are not responding to/not being terminated by CTRL_BREAK_EVENT.

It probably has something to do with not being able to attach the spawning process's console to
the docker container.

Process hierarchy:

```
test.exe -> PS script -> Docker container
```

The PS script is launched by creating a new console process group. The hope was that the docker container
would automatically become part of the same process group and would be terminated by the CTRL_BREAK_EVENT

The result now is that the docker container stays running.


## Useful links

- [Console Process Group])(https://docs.microsoft.com/en-us/windows/console/console-process-groups)
- [Identifying child processes in powershell](http://www.boldevin.com/?p=89)
- [Process explorer](https://docs.microsoft.com/en-us/sysinternals/downloads/process-explorer)
