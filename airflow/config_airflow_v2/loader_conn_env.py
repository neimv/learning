#! /usr/bin/env python

import os

from airflow import settings
from airflow.models import Connection, Variable


def main():
    all_envs = os.environ
    conn_envs = [env for env in all_envs if env.startswith('AIRFLOW_CONN_')]
    var_envs = [env for env in all_envs if env.startswith('AIRFLOW_VAR_')]

    for conn_name in conn_envs:
        conn = os.getenv(conn_name)
        name = conn_name.replace('AIRFLOW_CONN_', '').lower()

        try:
            conn = Connection(conn_id=name, uri=conn)
            session = settings.Session()
            session.add(conn)
            session.commit()
        except Exception as e:
            print(f'error: {e}')

    for var_name in var_envs:
        val = os.getenv(var_name)
        name = var_name.replace('AIRFLOW_VAR_', '').lower()

        try:
            var = Variable(key=name, val=val)
            session = settings.Session()
            session.add(var)
            session.commit()
        except Exception as e:
            print(f'error: {e}')


if __name__ == '__main__':
    main()
