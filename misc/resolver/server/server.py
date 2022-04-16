from ast import parse
from asyncio.log import logger
from dnslib.dns import RR
import logging
import random


class Resolver(object):
    def __init__(self, callback) -> None:
        self.callback = callback

    def resolve(self, request, handler):
        logging.debug(f"==>Request: {request}")
        reply = request.reply()
        response = request.questions[0].get_qname()
        parsed_question = str(request.questions[0]).split()
        logging.info(f"PARSED Q: {parsed_question}")
        if parsed_question[2] == "A":
            domain = parsed_question[0][1:]
            response = f"{domain} 60 A {self.__get_random_A()}"
            logging.info(f"REPLY DATA: {domain}, response {response}")
        elif parsed_question[2] == "TXT":
            domain = parsed_question[0][1:]
            txt_record = self.callback(domain)
            response = f"{domain} 60 TXT {txt_record}"
        logging.info(f"PARSED Q: {parsed_question}")
        logging.info(f"TEST: {request.questions[0]}")
        logging.debug(f"<==Replay: {response}")
        # "abc.def. 60 AAAA 2001:0db8:11a3:09d7:1f34:8a2e:07a0:765d\nabc.def. 60 A 127.0.0.1\nabc.def. 60 TXT testesttest"
        reply.add_answer(*RR.fromZone(response))
        return reply

    def __get_random_A(self):
        return str(random.randint(2,255)) + "." \
            + str(random.randint(2,255)) + "." \
            + str(random.randint(2,255)) + "." \
            + str(random.randint(2,255))