package docs

func init() {

	Definition["signal"] = "# `signal`\n\n> Sends a signal RPC\n\n## Description\n\n`signal` sends an operating system RPC (known as \"signal\") to a specified\nprocess, identified via it's process ID (\"pid\").\n\nThe following quote from [Wikipedia explains what signals](https://en.wikipedia.org/wiki/Signal_(IPC))\nare:\n\n> Signals are standardized messages sent to a running program to trigger\n> specific behavior, such as quitting or error handling. They are a limited\n> form of inter-process communication (IPC), typically used in Unix, Unix-like,\n> and other POSIX-compliant operating systems.\n>\n> A signal is an asynchronous notification sent to a process or to a specific\n> thread within the same process to notify it of an event. Common uses of\n> signals are to interrupt, suspend, terminate or kill a process.\n\n### Listing supported signals\n\nSignals will differ from one operating system to another. You can retrieve a\nJSON map with supported signals by running `signal` without any parameters.\n\n## Usage\n\n**Send a signal:**\n\n1. The first parameter is the process ID (int)\n2. The second parameter is the signal name (str). This will be all in\n   UPPERCASE and prefixed \"SIG\"\n\n```\nsignal pid SIGNAL\n```\n\n**List supported signals:**\n\n```\nsignal -> <stdout>\n```\n\n## Examples\n\n**Send a signal:**\n\n```\nfunction signal.SIGUSR1.trap {\n    bg {\n        exec <pid:GLOBAL.SIGNAL_TRAP_PID> $MUREX_EXE -c %(\n            event signalTrap example=SIGUSR1 {\n                out \"SIGUSR1 received...\"\n            }\n\n            out \"waiting for signal...\"\n            sleep 3\n        )\n    }\n    sleep 2 # just in case `exec` hasn't started yet\n    signal $GLOBAL.SIGNAL_TRAP_PID SIGUSR1\n}\n\ntest unit function signal.SIGUSR1.trap %{\n    StdoutMatch: \"waiting for signal...\\nSIGUSR1 received...\\n\"\n    DataType:    str\n    ExitNum:     0\n}\n```\n\n**List supported signals:**\n\n```\n» signal\n{\n    \"SIGABRT\": \"aborted\",\n    \"SIGALRM\": \"alarm clock\",\n    \"SIGBUS\": \"bus error\",\n    \"SIGCHLD\": \"child exited\",\n    \"SIGCONT\": \"continued\",\n    \"SIGFPE\": \"floating point exception\",\n    \"SIGHUP\": \"hangup\",\n    \"SIGILL\": \"illegal instruction\",\n    \"SIGINT\": \"interrupt\",\n    \"SIGIO\": \"I/O possible\",\n    \"SIGKILL\": \"killed\",\n    \"SIGPIPE\": \"broken pipe\",\n    \"SIGPROF\": \"profiling timer expired\",\n    \"SIGPWR\": \"power failure\",\n    \"SIGQUIT\": \"quit\",\n    \"SIGSEGV\": \"segmentation fault\",\n    \"SIGSTKFLT\": \"stack fault\",\n    \"SIGSTOP\": \"stopped (signal)\",\n    \"SIGSYS\": \"bad system call\",\n    \"SIGTRAP\": \"trace/breakpoint trap\",\n    \"SIGTSTP\": \"stopped\",\n    \"SIGTTIN\": \"stopped (tty input)\",\n    \"SIGTTOU\": \"stopped (tty output)\",\n    \"SIGURG\": \"urgent I/O condition\",\n    \"SIGUSR1\": \"user defined signal 1\",\n    \"SIGUSR2\": \"user defined signal 2\",\n    \"SIGVTALRM\": \"virtual timer expired\",\n    \"SIGWINCH\": \"window changed\",\n    \"SIGXCPU\": \"CPU time limit exceeded\",\n    \"SIGXFSZ\": \"file size limit exceeded\"\n}\n```\n\n## Flags\n\n* `SIGINT`\n    **\"Signal interrupt\"** -- equivalent to pressing `ctrl`+`c`\n* `SIGQUIT`\n    **\"Signal quit\"** -- requests the process quits and performs a core dump\n* `SIGTERM`\n    **\"Signal terminate\"** -- request for a processes termination. Similar to `SIGINT`\n* `SIGUSR1`\n    **\"Signal user 1\"** -- user defined\n* `SIGUSR2`\n    **\"Signal user 2\"** -- user defined\n\n## Detail\n\nThe interrupts listed above are a subset of what is supported on each operating\nsystem. Please consult your operating systems docs for details on each signal\nand what their function is.\n\n### Windows Support\n\nWhile Windows doesn't officially support signals, the following POSIX signals\nare emulated:\n\n```go\nvar interrupts = map[string]syscall.Signal{\n\t\"SIGHUP\":  syscall.SIGHUP,\n\t\"SIGINT\":  syscall.SIGINT,\n\t\"SIGQUIT\": syscall.SIGQUIT,\n\t\"SIGILL\":  syscall.SIGILL,\n\t\"SIGTRAP\": syscall.SIGTRAP,\n\t\"SIGABRT\": syscall.SIGABRT,\n\t\"SIGBUS\":  syscall.SIGBUS,\n\t\"SIGFPE\":  syscall.SIGFPE,\n\t\"SIGKILL\": syscall.SIGKILL,\n\t\"SIGSEGV\": syscall.SIGSEGV,\n\t\"SIGPIPE\": syscall.SIGPIPE,\n\t\"SIGALRM\": syscall.SIGALRM,\n\t\"SIGTERM\": syscall.SIGTERM,\n}\n```\n\n### Plan 9 Support\n\nPlan 9 is not supported.\n\n### Catching incoming signals\n\nSignals can be caught (often referred to as \"trapped\") in Murex with an event:\n`signalTrap`. Read below for details.\n\n## See Also\n\n* [Interactive Shell](../user-guide/interactive-shell.md):\n  What's different about Murex's interactive shell?\n* [Terminal Hotkeys](../user-guide/terminal-keys.md):\n  A list of all the terminal hotkeys and their uses\n* [`MUREX_EXE` (path)](../variables/MUREX_EXE.md):\n  Absolute path to running shell\n* [`bg`](../commands/bg.md):\n  Run processes in the background\n* [`event`](../commands/event.md):\n  Event driven programming for shell scripts\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`signalTrap`](../events/signaltrap.md):\n  Trap OS signals"

}
