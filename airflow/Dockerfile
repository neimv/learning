FROM apache/airflow:2.2.2

ENV AIRFLOW_HOME=/opt/airflow
# Extra packages
USER root
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
         vim cron wget\
  && apt-get autoremove -yqq --purge \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
USER airflow


WORKDIR $AIRFLOW_HOME

# Config files
COPY config_airflow_v2/airflow.cfg ./airflow.cfg
COPY config_airflow_v2/requirements.txt ./requirements.txt

# requirements
RUN pip --no-cache-dir install --upgrade awscli s3cmd
RUN pip install -r requirements.txt

COPY dags/ ./dags
COPY plugins/ ./plugins
#COPY tests /opt/airflow/tests
