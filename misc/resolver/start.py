import threading
import logging
from time import sleep

from dnslib.server import DNSServer, DNSLogger

from server.server import Resolver
from server.sql import SQLManager


def callback(domain):
    logging.debug(f"[*] Callback called with request: {domain}")
    orm = SQLManager()
    txt = orm.select_txt(domain)
    return txt

def main():
    try:
        resolver = Resolver(callback)
        hooks = "-send,-recv,-request,-reply,-truncated,-error,-data"
        dns_logger = DNSLogger(log=hooks)
        server = DNSServer(resolver, port=12634, logger=dns_logger, address="0.0.0.0", tcp=False)
        server.start_thread()
    except Exception as e:
        logging.error(str(e))
    while 1:
        sleep(1)


if __name__ == "__main__":
    logging.basicConfig(level="DEBUG")
    main()