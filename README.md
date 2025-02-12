![](https://badge.byted.org/ci/passed/ad/creative_tcpp_server_i18n/ref/master)
![](https://badge.byted.org/ci/coverage/ad/creative_tcpp_server_i18n)

## Tiktok Creative AIGC Suite

### Build

```shell script
./build.sh
```

### Run

### Use Makefile to update kitex
```
#更新客户端代码，需要指定thrift_file, 比如更新ttcx_business代码如下
make k_client thrift=../service_rpc/ad/tcpp/ad_ttcx_business.thrift
```

### Generate Kitex Client
```
#统一使用KiteX进行代码生成
kitex -module code.byted.org/ad/creative_tcpp_server_i18n xx.thrift
```

### Update DB Model
```
make gen
```
