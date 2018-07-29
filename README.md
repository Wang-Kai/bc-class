### bc-class

### API 文档

#### 创建 Deloyment
**POST**:   /create/deployment

Request:
```js
{
    "name":"bc-class",  // deployment 名字，确保唯一性
    "labels":{
        "app":"bc-class"    // deployment 标签，其值最好与名字一致
    },
    "pod":{
        "labels":{
            "app":"bc-class"    // pod 标签，其值最好与 deployment 名字一致
        },
        "containers":[      // 容器数组
            {
                "name":"novnc",  // required
                "image":"uhub.service.ucloud.cn/safehouse/novnc",    // required
                "command":[
                    "/bin/sh"
                ],
                "args":[
                    "-c",
                    "/usr/src/app/noVNC/utils/launch.sh --vnc localhost:5091"
                ],
                "containerPorts":[
                    {
                        "container_port":6080    // required
                    }
                ]
            },
            {
                "name":"ubuntu-xfce-vnc",
                "image":"uhub.service.ucloud.cn/safehouse/ubuntu-xfce-vnc",
                "containerPorts":[
                    {
                        "container_port":5901
                    }
                ]
            }
        ]
    }
}
```

Response: 
```js
{
    "code": 200,
    "message": "Create successful"
}
```

#### 罗列所有课程环境
**GET**:    /list/deployment

Response: 
```js
{
    code: 200,
    message: 'ok',
    data: [
        {
            name: 'deployment name',
            available: 60
        }
    ]
}
```

#### 罗列一个 deployment 下面的 pod
**GET**:    /list/pod/:deployment

Response:
```js
{
    "code":200,
    "data":[
        {
            "name":"etherum-64b46fb57f-qvdwk",
            "ip":"2002:ac1c:b401:1::2"
        },
        {
            "name":"etherum-64b46fb57f-r98ng",
            "ip":"2002:ac1c:b401:1::3"
        },
        {
            "name":"etherum-64b46fb57f-szhzl",
            "ip":"2002:ac1c:b401:1::3"
        }
    ]
}
```

#### 删除一个 Pod
**DELETE**:    /pod/:name

Response:
```js
{
    code: 200,
    message: 'Delete successful'
}
```

#### 扩容 deployment
**GET**:    /scale/:deployment/:amount

Response:
```js
{
    code: 200,
    message: 'Scale successful'
}
```



