import queue
import time

import websocket
from locust import events, TaskSet, task, User, constant_pacing

import ChatProto_pb2


class WebSocketClient(object):

    def __init__(self, host):
        self.host = host
        self.ws = websocket.WebSocket()
        self.name = "WebSocketTest"

    def record_result(self, response_time, response_length=0, exception_msg=None):
        self.name = "WebScocketTest"
        if exception_msg:
            events.request_failure.fire(request_type="ws", name=self.name, response_time=response_time,
                                        exception=exception_msg,
                                        response_length=response_length)
        else:
            events.request_success.fire(request_type="ws", name=self.name, response_time=response_time,
                                        response_length=response_length)

    def connect(self, burl, request_name='ws'):
        self.name = request_name
        start_time = time.time()
        try:
            self.conn = self.ws.connect(url=burl)
        except websocket.WebSocketTimeoutException as e:
            total_time = int((time.time() - start_time) * 1000)
            self.record_result(response_time=total_time, exception_msg=e)
        except BrokenPipeError as e:
            total_time = int((time.time() - start_time) * 1000)
            self.record_result(response_time=total_time, exception_msg=e)
        else:
            total_time = int((time.time() - start_time) * 1000)
            self.record_result(response_time=total_time)
        return self.conn

    def recv(self):
        return self.ws.recv()

    def send(self, msg):
        self.ws.send(msg)

    def rec_msg(self, expect_str=None, time_out=500, forever=False, time_out_per=60, run_user=None):
        pass


class SupperSC(TaskSet):

    def on_start(self):
        self.url = 'ws://127.0.0.1:8080/ws'
        self.data = {}
        self.client.connect(self.url)

    @task(2)
    def Send(self):
        msg = ChatProto_pb2.ChatRequest()
        msg.type = "talk"
        msg.content = "hello go"
        msg.userName = "tom"
        self.client.send(msg.SerializeToString())
        while True:
            res = self.client.recv()
            if res != "":
                events.request_success.fire(
                    request_type="WSR",
                    name="Send",
                    response_time=100,
                    response_length=100,
                    exception=None,
                    context={},
                )
            else:
                events.request_failure.fire(
                    request_type="WSR",
                    name="Send",
                    response_time=100,
                    response_length=100,
                    exception=None,
                    context={},
                )



class WSUser(User):
    host = '127.0.0.1:8080'  #待测主机
    wait_time = constant_pacing(1)  # 单个用户执行间隔时间
    tasks = [SupperSC]

    def __init__(self, *args, **kwargs):
        super(WSUser, self).__init__(*args, **kwargs)
        self.client = WebSocketClient(self.host)