Feature: Custom application creation

  Background:
    * url baseUrl
    * path 'application/custom'
    * def nginxDeployment = read('classpath:at/htl/beeyond/integration/util/nginx-deployment.yml.txt')
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @student
  Scenario: Create a valid custom application
    Given request
    """
    {
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "moritz",
      "class": "5AHIF",
      "to": "26.02.2023",
      "purpose": "SYP"
    }
    """
    When method POST
    Then status 201
    Given url responseHeaders['Location'][0]
    When method GET
    Then def runningApplication = response