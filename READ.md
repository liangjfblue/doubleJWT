
长短token，短用于资源接口请求，长用于请求刷新token方案。
- token（用于一般的资源请求）
- refreshToken（用于请求刷新token，返回两个新的token（短和长））

测试步骤：
- 1、创建数据库和建表（使用create.sql脚本。在于script目录）
- 2、测试时修改config.yaml的两个时间（shortTime 和 longTime），比如短的10s，长的30s。
- 3、调用 register 接口注册账号
- 4、调用 login 接口登录
- 5、根据 login 得到的 token ，调用 comment 接口评论
- 6、等待 shortTime 后，token失效，使用 refreshtoken 接口再次得到 token 和 refreshToken 。
- 7、若 refreshToken过期失效，再次调用login 接口登录得到 token 和 refreshToken 。

[POST] ```http://172.16.0.66:8099/v1/user/register```

请求：

    {
        "username":"laingjf",
        "password":"123456"
    }

返回：

    {
        "code": 1,
        "message": "OK",
        "data": {
            "uuid": "5e4eb0f2453343e98919b930d8fcee52"
        }
    }

[POST] ```http://172.16.0.66:8099/v1/user/login```

请求：

    {
        "username":"laingjf",
        "password":"123456"
    }

返回：

    {
        "code": 1,
        "message": "OK",
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY5NzA5NDcsImlhdCI6MTU2Njk3MDkyNywiaWQiOjEsIm5iZiI6MTU2Njk3MDkyNywidXNlcm5hbWUiOiJsYWluZ2pmIn0.pB5-TERuhSLogg-kdarK5J29LedfZJutaWi9IxeECs0",
            "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY5NzA5ODcsImlhdCI6MTU2Njk3MDkyNywiaWQiOjEsIm5iZiI6MTU2Njk3MDkyNywidXNlcm5hbWUiOiJsYWluZ2pmIn0.YkMm-a5Krp9b4Hsgxpod0y7KrLGMJJPlqkj83SKefmQ"
        }
    }

[POST] ```http://172.16.0.66:8099/v1/service/comment```

headers Authorization ：Bearer token

请求：

    {
        "topic_id":10,
        "topic_type":"sport",
        "content":"我来评论了，你好帅啊",
        "from_uid":1
    }

返回：

    {
        "code": 1,
        "message": "OK",
        "data": null
    }


[POST] ```http://172.16.0.66:8099/v1/base/refreshtoken```

headers Authorization ：Bearer refreshToken

请求：

    {
        "code": 1,
        "message": "OK",
        "data": {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY5NzA5NDcsImlhdCI6MTU2Njk3MDkyNywiaWQiOjEsIm5iZiI6MTU2Njk3MDkyNywidXNlcm5hbWUiOiJsYWluZ2pmIn0.pB5-TERuhSLogg-kdarK5J29LedfZJutaWi9IxeECs0",
            "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjY5NzA5ODcsImlhdCI6MTU2Njk3MDkyNywiaWQiOjEsIm5iZiI6MTU2Njk3MDkyNywidXNlcm5hbWUiOiJsYWluZ2pmIn0.YkMm-a5Krp9b4Hsgxpod0y7KrLGMJJPlqkj83SKefmQ"
        }
    }




