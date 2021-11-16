Feature: Template creation endpoint

  Background:
    * url baseUrl
    * path 'template'
    * configure headers = { Authorization: '#(auth(karate.tags))' }
    * def nginxDeployment = read('classpath:at/htl/beeyond/integration/util/nginx-deployment-template.yml.txt')
    * def insertTemplate = read('classpath:at/htl/beeyond/integration/util/create-template.feature')
    * def insertTemplateResponse = call insertTemplate
    * def template = insertTemplateResponse.template

  @teacher
  Scenario: Patch the newly created template
    Given path '/' + insertTemplateResponse.template.id
    And request
    """
    {
      "name": "Nginx Deployment",
      "description": "Static Webserver",
      "content": "#(nginxDeployment)",
      "fields": [
        {
          "label": "Server count",
          "wildcard": "replica",
          "description": "How many server should there be?"
        },
        {
          "label": "Port of your saver",
          "wildcard": "port",
          "description": "This will be the port that will be exposed to the world"
        }
      ]
    }
    """
    When method PATCH
    Then status 204