

## Creative AIGC Suite
### Summary
Creativity is a crucial factor in the success of streaming media, and advancements in Generative AI are making content creation more accessible, efficient, and seamless. These innovations empower creators to produce meaningful content at scale.  

Generative AI refers to AI-powered systems and models that generate new content—such as audio, images, text, and videos—also known as AI-Generated Content (AIGC).  

This project features the latest Generative AI products, capabilities, and solutions to help your clients and partners elevate creative productivity.  


### What is Creative AIGC Suite?
Creative AIGC Suite is a powerful suite of generative AI solutions designed to enhance your content creation journey.  
From scriptwriting to video production and asset optimization, Creative AIGC Suite streamlines every step, making content creation effortless and efficient—helping businesses produce impactful content that drives success. 
Built specifically for branded content, Creative AIGC Suite empowers brands to:

* Democratize creativity – Simplifying the creative process from concept to execution, making high-quality content creation more accessible.
* Maximize efficiency – Leveraging generative AI to scale content production, enable hyper-personalization, and facilitate continuous iteration.
* Unlock strategic insights – Going beyond content generation to identify cultural and industry trends, extract campaign learnings, and refine creative strategies for better decision-making.

### Creative AI Assistant 
Creative Assistant is a virtual assistant designed to collaborate intelligently with to shape creative direction.  

Drawing from extensive creative expertise, it provides valuable insights, inspiration, script generation and refinement, best practices, and tailored solution recommendations—empowering users to produce high-quality, impactful content with ease.

### Creative Studio 
Creative Studio is an AI-powered video generation platform that transforms minimal user input into high-quality videos in minutes.  

Easily convert product information or a URL into a compelling video, enhance storytelling with a digital avatar narrating your script, and expand your reach by localizing content through AI-driven translation and dubbing capabilities.

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
