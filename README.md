### bc-class

### API 文档

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



