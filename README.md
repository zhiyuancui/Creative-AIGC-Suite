

## Creative AIGC Suite
### Summary
Creative is one of the most critical factors for success on stream media, and with recent advancements in Generative AI, there's renewed focus on making the process of creating more accessible, efficient and simple, empowering the world to create meaningful content at scale.   
Generative AI refers to AI systems and models that can be used to create new content, including audio, images, text, and videos, and may also be referred to as AIGC (AI-Generated Content).   
Within this project, you'll find the latest Generative AI products, capabilities and solutions to help your clients and partners supercharge creative productivity.   

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
