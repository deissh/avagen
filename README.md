
![Example](https://i.imgur.com/4tk6a5m.png)

# Avagen
Generate avatars with initials from names.

```bash
Bombarding http://127.0.0.1:8080?name=q&type=png for 10s using 1000 connection(s)
[==========================================================================] 10s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec     79390.31   18672.76  212498.18
  Latency       12.59ms    22.79ms      1.26s
  HTTP codes:
    1xx - 0, 2xx - 789992, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:   109.92MB/s
```

## API

Path: `/`
Desc: Generate new avatars or return from cache

Params
 * `name` Nick or user name
 * `size` Image size (default 128px)

## Docs

ToDo

#### Glyph metrics

![](https://www.freetype.org/freetype2/docs/tutorial/metrics.png)
