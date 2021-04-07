script_template = '''
  zarubaRunContainerTask:
    icon: 🐳
    description: Run zarubaServiceName (containerized)
    extend: core.startDockerContainer
    dependencies:
    - zarubaBuildImageTask
    timeout: 1h
    env:
      <<: *zarubaServiceNameEnv
    lconfig:
      ports: *zarubaServiceNamePorts
    config:
      imageName: &zarubaServiceNameImage zarubaImageName
      imageTag: latest
      containerName: &zarubaServiceNameContainer zarubaContainerName
      rebuild: true
      localhost: host.docker.internal
      expose: lconfig.ports
    
  
  zarubaStopContainerTask:
    icon: 🐳
    description: Stop zarubaServiceName's container
    extend: core.stopDockerContainer 
    config:
      containerName: *zarubaServiceNameContainer


  zarubaRemoveContainerTask:
    icon: 🐳
    description: Remove zarubaServiceName's container
    extend: core.removeDockerContainer 
    config:
      containerName: *zarubaServiceNameContainer
  

  zarubaBuildImageTask:
    icon: 🐳
    description: Build zarubaServiceName's image
    extend: core.buildDockerImage
    location: zarubaTaskLocation
    timeout: 1h
    config:
      imageName: *zarubaServiceNameImage


  zarubaPushImageTask:
    icon: 🐳
    description: Push zarubaServiceName's image
    extend: core.pushDockerImage
    dependencies:
    - zarubaBuildImageTask
    timeout: 1h
    config:
      imageName: *zarubaServiceNameImage
'''

def get_script_template() -> str:
    return script_template