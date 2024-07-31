import os
from crewai import Agent

os.environ["OPENAI_API_BASE"]='http://localhost:11434/v1'

os.environ["OPENAI_MODEL_NAME"]='crewai-llama3'
# custom_model_name from the bash script

os.environ["OPENAI_API_KEY"] = "NA"

database_specialist_agent = Agent(
  role = "Database specialist",
  goal = "Provide data to answer business questions using SQL",
  backstory = '''You are an expert in SQL, so you can help the team
  to gather needed data to power their decisions.
  You are very accurate and take into account all the nuances in data.''',
  allow_delegation = False,
  verbose = True
)

tech_writer_agent = Agent(
  role = "Technical writer",
  goal = '''Write engaging and factually accurate technical documentation
    for data sources or tools''',
  backstory = '''
  You are an expert in both technology and communications, so you can easily explain even sophisticated concepts.
  You base your work on the factual information provided by your colleagues.
  Your texts are concise and can be easily understood by a wide audience.
  You use professional but rather an informal style in your communication.
  ''',
  allow_delegation = False,
  verbose = True
)

CH_HOST = 'http://localhost:8123' # default address

def get_clickhouse_data(query, host = CH_HOST, connection_timeout = 1500):
  r = requests.post(host, params = {'query': query},
    timeout = connection_timeout)
  if r.status_code == 200:
      return r.text
  else:
      return 'Database returned the following error:\n' + r.text
