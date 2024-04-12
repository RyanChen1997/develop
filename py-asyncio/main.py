import asyncio


async def coroutine_example():
    await asyncio.sleep(1)
    print("zhihu ID: Zarten")


coro = coroutine_example()

loop = asyncio.get_event_loop()
loop.run_until_complete(coro)
loop.close()
