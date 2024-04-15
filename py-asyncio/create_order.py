import requests
from urllib.parse import urlparse, urlsplit


def create_order(email, qq_account):
    url = "https://maifanx.shop/create-order"

    headers = {
        "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
        "accept-language": "zh-CN,zh;q=0.9",
        "cache-control": "max-age=0",
        "content-type": "application/x-www-form-urlencoded",
        "cookie": "_ga=GA1.1.435105058.1713146903; XSRF-TOKEN=eyJpdiI6IjRab1B1QU1kcENTQmlyb09hTUJpQ2c9PSIsInZhbHVlIjoiSXVGZG9RMGJUMVYwV0Z1SkVib0F3ZXFaaHBuY0FTTkZFZ1pZbERtR3V5T1wvXC9jSUx1ZWxHY3I5MnM0QlhCN3VYV3NpTWVPNXQwRFVscUpNeTZPc2oyVFl4THdLZzBMNUpwTkZKWXFDOEFEVVVLODNqc2xhcmpJdUw5cWNVV0ZhbCIsIm1hYyI6ImM5YThjNDc4NmVjZGRkOTVjYzVlMjIyNjM3MmNmYzFkYzFmYzFjYTUzNTZhNDI1NjA2NTA2MmYyYWNmMDg2NGMifQ%3D%3D; _session=eyJpdiI6ImdDWWNBVmMxck00NHlJOWMzZnlLb2c9PSIsInZhbHVlIjoiME5IVEVzeTNKb3NXSm1RcVJ4blZmMENUREM4RlVpZkZxUERxeUQzTVJOWk1PQmFlcU1DNjNWcDgxTmNuSUJ6cnVEbGVmZ1JPcEwxTUVPa3VSdU8zSEd1akZ2MTVJTGRObG5aUU1iK3p6WlwvZ0dIZG1UOVZDOHdKSHY2VXdpVXdSIiwibWFjIjoiYmUxMDNkYmY2MWEwOTFkMTA5NmI5ODQ2OWY5MTdhMDRkNTRhNzI0ODY2OTU5OWNjZTFmNjExNDY0OTI0YmE3YSJ9; _ga_4PLV1058L3=GS1.1.1713146903.1.1.1713146917.46.0.0; crisp-client%2Fsession%2F0f6d053e-1466-4feb-9651-0597138077fe=session_69eb4608-2e36-4cb3-9039-c04378c1f1a7",
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
        "_token": "hcXG95CgJ9ootECIayB4gbQ0njRjWM2TO4y4EG4p",
        "gid": "9",
        "by_amount": "1",
        "email": email,
        "qq_account": qq_account,
        "payway": "1",
    }

    # Make the POST request with redirection allowed (default behavior)
    response = requests.post(url, headers=headers, data=data, allow_redirects=True)
    print(response)

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
    create_order("10293i10l@qq.com", "123123141")
