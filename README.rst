yet another golang logging lib

benchmark::

    go run benchmark/log_benchmark.go
    qps of     dummy1: 2727908735, runtime :3.67 s
    qps of     dummy2:  552468326, runtime :1.81 s
    qps of   variadic:    6151799, runtime :1.63 s
    qps of    logging:    4003802, runtime :2.50 s
    qps of      golog:    5986678, runtime :1.67 s
    qps of     beelog:    1170780, runtime :8.54 s


