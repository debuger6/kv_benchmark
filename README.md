# kv_benchmark

## usage
### compile
```shell script
go build
```

### run
eg.

_tikv_
```shell script
./kv_benchmark -h  # for help
./kv_benchmark -addr 127.0.0.1:2379 -db tikv -oc 100000 -tc 16
INFO[0000] [pd] create pd client with endpoints [127.0.0.1:2379] 
INFO[0000] [pd] leader switches to: http://127.0.0.1:2379, previous:  
INFO[0000] [pd] init cluster id 6780150801926342024     
Takes(s): 3.00, Total_op: 43331, OPS: 14442.64, Min_lat(us): 156, Avg_lat(us): 1106, max_lat(us): 22615
Takes(s): 6.00, Total_op: 85443, OPS: 14239.48, Min_lat(us): 156, Avg_lat(us): 1122, max_lat(us): 22615
Takes(s): 7.10, Total_op: 100000, OPS: 14075.31, Min_lat(us): 156, Avg_lat(us): 1125, max_lat(us): 22615
```
_pika_
```shell script
./kv-benchmark -addr 127.0.0.1:9221 -db pika -oc 100000 -tc 16
Takes(s): 3.00, Total_op: 41416, OPS: 13803.53, Min_lat(us): 141, Avg_lat(us): 1157, max_lat(us): 13670
Takes(s): 6.00, Total_op: 84549, OPS: 14090.97, Min_lat(us): 141, Avg_lat(us): 1134, max_lat(us): 13670
Takes(s): 7.09, Total_op: 100000, OPS: 14099.44, Min_lat(us): 141, Avg_lat(us): 1131, max_lat(us): 16813
```