# WeiXinBackEnd
微信后端开发，使用golang

# 请求大全
{"method":"GET","url":"/user"}
{
    "method":"PUT",
    "url":"/user",
    "json_data":{
        "username":"许一涵",
        "signature":"我很帅"
    }
}


{"method":"GET","url":"/salary/?page_size=5&page=1&company=狐&city=北京"}
{"method":"GET","url":"/salary/getById?id=1"}
{"method":"GET","url":"/salary/getByUserId?page_size=5&page=1&user_id=obGiG6n3SPlTapeLcCVx2VAg1la8"}
{
    "method":"POST",
    "url":"/salary/create",
    "json_data":{
        "Company":"华为",
        "City":"北京",
        "Salary":"20w",
        "Major":"软件工程",
        "CategoryFirst":"技术/开发",
        "CategorySecond":"前端开发"
    }
}

{
    "method":"DELETE",
    "url":"/salary?id=?"
}

{
    "method":"POST",
    "url":"/salary/creates",
    "json_data":{
        "salaries":[
            {
                "Company":"字节跳动",
                "City":"北京",
                "Salary":"230w",
                "Major":"软件工程",
                "CategoryFirst":"技术/开发",
                "CategorySecond":"前端开发"
            },
            {
                "Company":"字节跳",
                "City":"北京",
                "Salary":"300w",
                "Major":"软件工程",
                "CategoryFirst":"技术/开发",
                "CategorySecond":"前端开发"
            }
        ]
        
    }
}