Feature:

  Background:
    * url baseUrl
    * path 'application/custom'
    * def nginxDeployment = read('classpath:at/htl/beeyond/integration/util/nginx-deployment.yml.txt')
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Create a valid custom application
    Given request
    """
    {
      "note": "Nginx Deployment",
      "content": "#(nginxDeployment)",
      "namespace": "stuetz"
    }
    """
    When method POST
    Then status 204
