# 一个爬虫
# url = https://offershow.cn/?search_tab=2
# 爬取offer信息
import requests
import json

# 目标URL
url = "https://offershow.cn/api/od/search_job?size=210&page=4"

# 发送POST请求
response = requests.post(url)
# 检查响应状态码
if response.status_code == 200:
    # 将响应内容转换为JSON格式
    data = response.json()

    # 保存JSON数据到文件，指定ensure_ascii=False
    with open('data.json', 'w', encoding='utf-8') as f:
        json.dump(data, f, ensure_ascii=False)
else:
    print("请求失败，状态码：", response.status_code)