Feature:
  Background:
    * url baseUrl
    * path 'template'
    * configure headers = { Authorization: '#(teacherAuth())'}

  Scenario: Create a valid template application
    * def content = read('nginx-deployment-template.yml.txt')
    Given request
    """
    {
      "name": "Nginx Deployment 1",
      "description": "Static Webserver 1",
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
    Then status 201


  Scenario: Create a second valid template application
    * def content = read('nginx-deployment-template.yml.txt')
    Given request
    """
    {
      "name": "Nginx Deployment 2",
      "description": "Static Webserver 2",
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
    Then status 201
