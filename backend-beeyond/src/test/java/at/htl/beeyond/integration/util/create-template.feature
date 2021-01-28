Feature:

  Background:
    * url baseUrl
    * path 'template'
    * configure headers = {Authorization: '#(auth(karate.tags))'}

  @teacher
  Scenario: Create a valid template application
    * def content = read('classpath:at/htl/beeyond/integration/util/nginx-deployment-template.yml.txt')
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
    When method POST
    And status 201
    Given url responseHeaders['Location'][0]
    When method GET
    Then def template = response
