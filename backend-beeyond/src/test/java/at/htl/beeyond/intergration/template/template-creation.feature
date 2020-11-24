Feature: Template creation endpoint

  Background:
    * url baseUrl
    * path 'template'

  @teacher
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
   When method POST
   Then status 204


  @teacher
  Scenario: Create a template with no description

  @teacher
  Scenario: Create a template with no name and content

  @teacher
  Scenario: Create a template with a too long name and description

  @teacher
  Scenario: Create a template where at least 1 field is missing

  @teacher
  Scenario: Create a template where the field label and wildcard are missing

  @teacher
  Scenario: Create a template where the field label and wildcard are empty

  @teacher
  Scenario: Create a template where the field label, wildcard and description are too long
