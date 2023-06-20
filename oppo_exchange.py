
import requests
import threading
import random
import datetime
import time
verify = True
headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36",
    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
    "Cookie":
    'newopkey=nLDUbFjitwaIihCc-Y9DvSOlBfLDR0ALZzZoW4gF1MfnCt7txoFa6SGxVsOJ176wy_1CpqoFm2M; opkey=eyJpdiI6IkN3eXd2Q0U0NThGelpzVC95ZDI2SEE9PSIsInZhbHVlIjoiQ293NjI1Yk40ZWlhQ0x4blNJcDY4dmI3M1pycWQ4T2lUdFRPNjA1dXUyakVHU3p2UjdsK1l5M1JkMWRHN1I1UiszM2MrSDFOVDJVWU92ZGZheC93bFFyUDYzNTEraTRUS0dJbURoU05IQVFzbU41RTdXT2pqczFlTGNCUDM5SVEiLCJtYWMiOiJiMmNlZTBjYzdmMzlkNTExZDA5OTM5OWIxNGZiMDI3OTEyNDdhZjBjMzNiNTk0YTRhZDNiZmQ5OTM2OWE3NTBlIn0=; NEWOPPOSID=eyJpdiI6ImRJYSsrVU1GNU9MSi9YZGNWN01TdkE9PSIsInZhbHVlIjoialRkUml1Y1JxK1Zha051d0NBRmxwWlAyVlVtM1NqaGM4TkdlUW1NTElOaGZSckZzL0FrUE9jRFNXQ1JLdHp6QldBT2V1VGhUeWRLTkFaZTZzb0tnc1orWmFmMFZqMWoxTmw3d2tnMjRkY1BzTk81Zit4QXlsTEpBMURaM1R6U1UiLCJtYWMiOiI0YTQzODA1OGYzMTBiZmJlNWRhMmQ4ZjQ0MjQzMmRiODU3NjUzYWU5NDljYTg0Y2MwNmY3ZDU1ODZmOGIwYTc1In0=; sa_distinct_id=TElpSm5jT01UNGZralVuRTJVUVJFZz09; HEYTAPID=nLDUbFjitwaIihCc-Y9DvSOlBfLDR0ALZzZoW4gF1MfnCt7txoFa6SGxVsOJ176wy_1CpqoFm2M; avatar=https://uc-avatar-cn.heytapimage.com/titans-usercenter-avatar-bucket-cn/default-new.png; ssoid=916918726; username=eKo8MHHiRumo+/uxnkxzYVk=; createtime=1659764856000; sensorsdata2015jssdkcross={"distinct_id":"18271b66f5aaaa-0a5b838bf989de-653b5753-2073600-18271b66f5b785","first_id":"","props":{},"identities":"eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTgyNzFiNjZmNWFhYWEtMGE1YjgzOGJmOTg5ZGUtNjUzYjU3NTMtMjA3MzYwMC0xODI3MWI2NmY1Yjc4NSJ9","history_login_id":{"name":"","value":""},"$device_id":"18271b66f5aaaa-0a5b838bf989de-653b5753-2073600-18271b66f5b785"}; sajssdk_2015_cross_new_user=1'
    }
gift_list = [4,4,4,4,4,4]

def get_user():
    user_center_url = "https://scrm.oppo.com/activities/api/activity/userInfo"
    user_center_resp = requests.get(user_center_url, headers=headers, verify=verify)
    if user_center_resp.status_code == 200:
        try:
            user_center_resp = user_center_resp.json()
            userId = user_center_resp["data"]["userId"]
            time_stamp = user_center_resp["traceID"][0:10]
            if userId:
                return time_stamp
            else:
                return False
        except Exception as e:
            print(e)
def get_gift_list():
    gift_url = "https://scrm.oppo.com/activities/api/activity/ofans/oqiExchangeList"
    gift_resp = requests.post(gift_url, headers=headers, verify=verify)
    if gift_resp.status_code == 200:
        gift_resp = gift_resp.json()
        data = gift_resp["data"]
        data_keys = data.keys()
        for key in data_keys:
            if isinstance(data[key], dict):
                gift_list = data[key]["oqiExchangeList"]
                gift_ids = []
                for gift in gift_list:
                    gift_id = gift["id"]
                    gift_ids.append(gift_id)
                    gift_name = gift["name"]
                    limit = gift["costOqiAmount"]
                    print("礼物名称为{}，礼物id为{}，兑换要求为{}".format(gift_name, gift_id, limit))

def exchange(gift_id):
    print("当前兑换的礼物id为", gift_id)
    exchange_url = "https://scrm.oppo.com/activities/api/activity/ofans/oqiExchange"
    data = {"id": "{}".format(gift_id)}
    exchange_resp = requests.post(exchange_url, headers=headers, json=data, verify=verify)
    print(exchange_resp.json())


def asset_time():
    today = str(datetime.date.today())
    first_time = today + " " + str(datetime.time(13, 59, 58))
    second_time = today + " " + str(datetime.time(17, 59, 58))
    third_time = today + " " + str(datetime.time(21, 59, 58))
    first_time = int(time.mktime(time.strptime(first_time, "%Y-%m-%d %H:%M:%S")))
    second_time = int(time.mktime(time.strptime(second_time, "%Y-%m-%d %H:%M:%S")))
    third_time = int(time.mktime(time.strptime(third_time, "%Y-%m-%d %H:%M:%S")))
    preset_time_list = [first_time, second_time, third_time]
    return preset_time_list



if __name__ == '__main__':

    user = get_user()
    if user:
        get_gift_list()
        preset_time_list = asset_time()
        while True:
            system_time = int(get_user())
            for preset_time in preset_time_list:
                if preset_time < system_time < preset_time + 100:
                    for i in range(1000):
                        gift_id = random.choice(gift_list)
                        t1 = threading.Thread(target=exchange,args=(gift_id,))
                        t1.start()
                else:
                    print("不在时间范围内")
                    time.sleep(1)
    else:
        print("获取用户异常，请检查cookie")
