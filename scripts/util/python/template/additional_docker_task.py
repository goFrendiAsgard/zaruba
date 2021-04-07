script_template = '''
  zarubaStopContainerTask:
    icon: 🐳
    description: Stop zarubaServiceName's container
    extend: core.stopDockerContainer 
    config:
      containerName: zarubaContainerName
'''

def get_script_template() -> str:
    return script_template