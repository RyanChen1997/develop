import threading
import random
import string
from create_order import create_order

def task():
    print("start: ", name)
    while True:
        username_length = random.randint(5,10)
        username = ''.join(random.choices(string.ascii_lowercase, k=username_length))
        email = username + "@" + "qq.com"
        qq_account= username
        create_order(email, qq_account)

n_threads=5000
threads = []

if __name__ == '__main__':
    for i in range(n_threads):
        name = "Thread {}".format(i)
        thread = threading.Thread(target=task)
        thread.start()
        threads.append(thread)

    for thread in threads:
        thread.join()

