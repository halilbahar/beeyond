Feature:

  Background:
    * url baseUrl
    * path 'template'
    * configure headers = {Authorization: '#(teacherAuth())'}
    * print headers
    * print auth

  Scenario: Create a valid template application
    * def content = read('nginx-deployment-template.yml.txt')
    Given request
        """
        {
          "name": "Nginx Deployment",
          "description": "Static Webserver",
          "content": "#(content)",
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
    * print headers
    When method POST
    * print 'headers:', karate.prevRequest.headers
    And status 201
    And url responseHeaders['Location'][0]
    And method GET
    And request
    And def template = response


