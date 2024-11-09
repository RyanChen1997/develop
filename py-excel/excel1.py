import pandas as pd
import re

"""
@王五(wangwu) -> wangwu@huolala.cn
"""

# 读取 Excel 表格
input_file = '/Users/ryan/Downloads/input.xlsx'  # 输入文件的路径，确保文件名和路径正确
df = pd.read_excel(input_file)

# 从 "通行证/账户" 列中解析括号内的内容
def extract_content(account):
    match = re.search(r'\((.*?)\)', account)
    if match:
        return match.group(1)
    return ''

# 获取括号内的内容并生成新的列
df['新内容'] = df['通行证/账户'].apply(extract_content)
df['新内容'] = df['新内容'] + '@huolala.cn'

# 保存到新的 Excel 文件
output_file = 'output.xlsx'  # 输出文件的路径
df[['新内容']].to_excel(output_file, index=False)
