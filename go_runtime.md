# Go GMP Runtime & Scheduler

```text
+---------------------------------------------------------------------------------------------------+
|                                        GO RUNTIME ENGINE                                          |
+---------------------------------------------------------------------------------------------------+
                                                  |
                                                  v
     [ User Code: go worker() ] ──> Placed in `runnext` Slot (High-Priority Cache)
                                                  | (If full -> pushed to LRQ -> overflow to GRQ)
                                                  v
        +-----------------------------------------------------------------------------------+
        |                         GLOBAL RUN QUEUE (GRQ)                                    |
        |                         [ Protected by Mutex Lock ]                               |
        +-----------------------------------------------------------------------------------+
                   ^                                                         |
                   | (Overflow / Preempted / 1/61 Tick Check)                | (Scheduled Batch Pull)
                   |                                                         v
        +-----------------------------------------------------------------------------------+
        | LOGICAL PROCESSORS (P)   [ Total Count = GOMAXPROCS ]                             |
        |                                                                                   |
        |   +-------------------------------+               +-------------------------------+   |
        |   | P1 Context                    |               | P2 Context                    |   |
        |   | - runnext: G10                |               | - runnext: nil                |   |
        |   | - LRQ: [G2, G3, ... (max 256)]|               | - LRQ: [G5, G6, ...]          |   |
        |   | - Memory Alloc Cache (mcache) |               | - Memory Alloc Cache (mcache) |   |
        |   +-------------------------------+               +-------------------------------+   |
        +-----------------------------------------------------------------------------------+
                          |                                               |
                          v (Bound M to P)                                v (Bound M to P)
        +-----------------------------------------------------------------------------------+
        | OS THREADS (M)           [ Managed by Kernel, up to 10,000 threads ]              |
        |                                                                                   |
        |   +-------------------------------+               +-------------------------------+   |
        |   | M1 (OS Thread)                |               | M2 (OS Thread)                |   |
        |   +-------------------------------+               +-------------------------------+   |
        +-----------------------------------------------------------------------------------+
                          |                                               |
         +----------------+----------------+                              |
         |                                 |                              |
         v                                 v                              v
  [ Network I/O ]                 [ Disk I/O / Cgo ]             [ Active CPU Execution ]
         |                                 |                              |
         v                                 v                              v
  Netpoller                      Handoff (`handoffp`)            Executing Goroutine ($G$)
  (`epoll`/`kqueue`)             (M1 blocks in kernel;            State: `_Grunning`
  Non-blocking socket            P1 detaches and is              on physical CPU core
  management                     adopted by idle M3)
  ```
