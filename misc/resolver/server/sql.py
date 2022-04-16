import sqlite3
import logging
from venv import create


class SQLManager(object):
    def __init__(self) -> None:
        self.connector = sqlite3.connect("resolver.db")
        self.cursor = self.connector.cursor()
        self.create_all()

    def create_all(self):
        query = "CREATE TABLE IF NOT EXISTS cache (domain TEXT, txt TEXT);"
        self.cursor.execute(query)
        query = "insert into cache (domain, txt) values ('JDSAJbjdsakbjsd.jasdhbjk.', 'YetiCTF{ya_uzhe_ne_chelovek_ya_zver_n4y1}')"
        self.cursor.execute(query)
        self.connector.commit()
        query = "insert into cache (domain, txt) values ('google.com.', 'nice try:)))')"
        self.cursor.execute(query)
        self.connector.commit()
        return True

    def select_txt(self, domain):
        domain = domain.replace(r'\032', ' ')
        logging.info(f"DOMAIN to DB: {domain}")
        try:
            query = f"select txt from cache where domain='{domain}'"
            self.cursor.execute(query)
            response = self.cursor.fetchone()
        except Exception as e:
            self.connector.close()
            response = str(e)
        logging.info(f"Data from db: domain: {domain}; txt: {response}")
        if response is None:
            response = ["sqlite dns resolver cache: txt empty"]
        return str(''.join(response)).replace('\\' , '')