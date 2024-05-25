# pv
echo 'pv='
awk '{print $1}' nginx/my-site/log/access.log | sort | uniq -c | awk '{total += $1} END {print total}'
# uv
echo 'uv='
awk '{print $1, $2}' nginx/my-site/log/access.log | sort -u | wc -l
