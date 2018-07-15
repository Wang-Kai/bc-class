### bc-class

### API 文档

#### 罗列所有课程环境
URL:    /list/deployment
Method: GET

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

#### 处理一个用户的接入
URL:    /access/:deployment/:user
Method: GET

Response:
```js
{
    code: 200,
    message: 'ok',
    data: {
        name: "bc-class-74b5bdb46f-44575",
        ip: "2002:ac12:fed0:1::24"
    }
}
```



