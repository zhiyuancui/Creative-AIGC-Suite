

## Creative AIGC Suite
### Summary
Creative is one of the most critical factors for success on stream media, and with recent advancements in Generative AI, there's renewed focus on making the process of creating more accessible, efficient and simple, empowering the world to create meaningful content at scale.   
Generative AI refers to AI systems and models that can be used to create new content, including audio, images, text, and videos, and may also be referred to as AIGC (AI-Generated Content).   
Within this project, you'll find the latest Generative AI products, capabilities and solutions to help your clients and partners supercharge creative productivity.   

### What is Creative AIGC Suite?
Creative AIGC Suite is a suite of generative AI solutions that elevates your content creation journey.  
With Creative AIGC Suite, everything from helping to write a script to producing a video and optimizing assets is effortless and efficient—fueling business success with content that strikes a chord.
Everything we're building for Creative AIGC Suite is designed to help make branded content. It can help brands:  
* Level the creative playing field: We are enabling the creative journey from concept to completion.  
* Boost productivity: By leveraging generative AI, businesses can supercharge their productivity in content creation. It makes it more possible to create at scale, integrate hyper-personalization, and make constant iterations. 
* Uncover insights: Generative AI goes beyond just creating content; it also identifies cultural and industry trends as well as campaign learnings, enabling brands to make informed decisions and refinements to their creative strategies.

### Creative Assistant 
Creative Assistant is a virtual assistant designed to intelligently collaborate with advertisers, creative partners and content creators on creative direction.  
The Assistant draws information from a wealth of creative knowledge to provide quality responses covering creative inspiration, insights, script generation and refinement, best practices and solution recommendations. 

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
