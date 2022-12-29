FROM apache/airflow:2.3.0

# Extra packages
USER root
RUN apt-get update \
  && apt-get install -y --no-install-recommends \
         vim cron wget\
  && apt-get autoremove -yqq --purge \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*
USER airflow

# Config files
COPY tests/airflow_dev.cfg ./airflow.cfg
COPY tests/requirements.txt ./requirements.txt
COPY config_airflow_v2/default_variables.json ./default_variables.json
COPY config_airflow_v2/default_connections.json ./default_connections.json

# requirements
RUN pip --no-cache-dir install --upgrade awscli s3cmd
RUN pip install -r requirements.txt



