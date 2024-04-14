import requests
from urllib.parse import urlparse, urlsplit


def create_order():
    url = "https://maifanx.shop/create-order"

    headers = {
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "accept-language": "zh-CN,zh;q=0.9",
        "cache-control": "max-age=0",
        "content-type": "application/x-www-form-urlencoded",
        "cookie": "_ga=GA1.1.1103622423.1713063943; crisp-client%2Fsession%2F0f6d053e-1466-4feb-9651-0597138077fe=session_35dfd0b0-edba-4acd-9cc9-0da5b3b0cc0f; crisp-client%2Fsocket%2F0f6d053e-1466-4feb-9651-0597138077fe=1; dujiaoka_orders=eyJpdiI6IlpKWnlpOUc3NVpRZ0owRGJHSEpsc1E9PSIsInZhbHVlIjoidHJJbm9MTit1T0dvOEhDb2oyaUpGT25Lb2ZJa3NVR2VGRUk4aXpNTDJwb0ljT1QzTVJ6T0dKNUgwTGFqUEY5Zm1RTVJ1b2tzOUhiSHlJVEZ6TUJRaXZ6bjlrY3dHeEZTaEkyU2tKdElkcFl5YTFZVFJyM3lmRjBpOUlYeitmQmhTY0ZyZFhwSkpaOWYrRjdCWDRCMk1OazZkU0VtMlZCVzZYVUZiZW1NazdRPSIsIm1hYyI6IjU3OWQyNTc2NjQyOTliMTA0MDQ5MDM5MmMyZDdjMWRjYWFmNjQ5YzdmYmZhYzUwNDhiODdmOTEzODNlZDcyMDQifQ%3D%3D; XSRF-TOKEN=eyJpdiI6IkZMVGhSeXRDdVJSb0F4bGN2Uk0xRmc9PSIsInZhbHVlIjoiTVlsYTUrNDZyS3cxR1lzNGFcL2M3bUVUVDc1WmwxdkJWZ0FJaDlUUHBVWFRJS0oweld1T0dlc1ZKeWV1VnZqZWx3d2VmMm1MWmY0TjVRTTFMVStQU29UVVdNZkt4RkhxcU1rOGhoZjFudlJjZ0xqTjBiRlwveUVLd2xuQnBpdXY0cyIsIm1hYyI6ImI5MmMxODAyNWE3NDg1ZmRiMDdmMTE1NmE2NjUzMTE5NGJlMzFmZTNiMzg4Mzk1OTc3OWFlYTE4ZmNiOTU2NWQifQ%3D%3D; _session=eyJpdiI6IndwRWRWeWZIZVBpSHRJdEx5NitmMmc9PSIsInZhbHVlIjoiYm1UMmJyMHoxSGFQOUJ3ZUNrcjFVOWowZTQ5M0FuQ1wvckx5SG4yaVpjYU1yQ0NYbVhjNTVQXC96T1VlbmdBUmx0bVBXK2VkTGNJZlZEZTh3VEd5b2dDdkVzdDhOaUZuNGU4ZlV0bldPbU8xWHVVTEhQQnJlZnJSb1pGMlR6M3NlUiIsIm1hYyI6IjMwZWNkYWRkYjgwMDVhMWNlYjc2NjhiYmNmZDdhNGY0OTNhNmU4ZWVhY2JkMjZlMDdmNmZiOThlZjk3MTk3NTEifQ%3D%3D; _ga_4PLV1058L3=GS1.1.1713102629.2.1.1713103289.32.0.0",
        "origin": "https://maifanx.shop",
        "referer": "https://maifanx.shop/buy/9",
        "sec-ch-ua": '"Google Chrome";v="123", "Not:A-Brand";v="8", "Chromium";v="123"',
        "sec-ch-ua-mobile": "?0",
        "sec-ch-ua-platform": '"Chrome OS"',
        "sec-fetch-dest": "document",
        "sec-fetch-mode": "navigate",
        "sec-fetch-site": "same-origin",
        "sec-fetch-user": "?1",
        "upgrade-insecure-requests": "1",
        "user-agent": "Mozilla/5.0 (X11; CrOS x86_64 14541.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",
    }

    data = {
        "_token": "GMETAi3HsRNyp0N8WvD759OuDHYwDMxneWSh8Ubf",
        "gid": "9",
        "by_amount": "1",
        "email": "912386561@qq.com",
        "qq_account": "749172973",
        "payway": "1",
    }

    # Make the POST request with redirection allowed (default behavior)
    response = requests.post(url, headers=headers, data=data, allow_redirects=True)

    # Access the final URL after redirection
    final_url = response.url
    # print("Final URL:", final_url)
    print(get_order_id(final_url))

    # If you need to inspect the status codes and URLs in the redirection chain:
    # for resp in response.history:
    #     print("Redirected URL:", resp.url)
    #     print("Status code:", resp.status_code)

    # Print the final response text
    # print("Response body from final URL:", response.text)


def get_order_id(url):
    # Use urlsplit to handle URLs robustly
    parsed_url = urlsplit(url)
    path_parts = parsed_url.path.split("/")
    path_parts = [part for part in path_parts if part]
    # print(path_parts)

    if len(path_parts) > 1:
        last_element_before_slash = path_parts[-1]
    else:
        last_element_before_slash = None
    return last_element_before_slash


if __name__ == "__main__":
    create_order()
