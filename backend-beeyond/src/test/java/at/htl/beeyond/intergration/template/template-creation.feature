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
    * def content = read('nginx-deployment-template.yml.txt')
    Given request
    """
    {
      "name": "Nginx Deployment",
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
  Scenario: Create a template with no name and content
    * def content = read('nginx-deployment-template.yml.txt')
    Given request
    """
    {
      "description": "Static Webserver"
    }
    """
    When method POST
    Then status 422

  @teacher
  Scenario: Create a template with a too long name and description
    * def generateString = read('string-generator.js')
    * def content = read('nginx-deployment-template.yml.txt')
    Given request
    """
    {
      "name": #(generateString()),
      "description": #(generateString()),
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
    Then status 422
    And match response contains {message:'This field needs to be between 0 and 255 characters', key: 'name', value:'#ignore'}
    And match response contains {message:'This field needs to be between 0 and 255 characters', key: 'description', value:'#ignore'}

  @teacher
  Scenario: Create a template where at least 1 field is missing
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
        }
      ]
    }
    """
    When method POST
    Then status 422
    And match response contains {message:'Missing fields: [port], obsolete fields: []',value:'',key:''}

  @teacher
  Scenario: Create a template where the field label and wildcard are missing
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
          "description": "This will be the port that will be exposed to the world"
        }
      ]
    }
    """
    When method POST
    Then status 422
    And match response contains {message:'This field cannot be empty',value:'',key:'fields[1].wildcard'}
    And match response contains {message:'This field cannot be empty',value:'',key:'fields[1].label'}

  @teacher
  Scenario: Create a template where the field label and wildcard are empty
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
          "label": "",
          "wildcard": "",
          "description": "This will be the port that will be exposed to the world"
        }
      ]
    }
    """
    When method POST
    Then status 422
    And match response contains {message:'This field cannot be empty',value:'',key:'fields[1].wildcard'}
    And match response contains {message:'This field cannot be empty',value:'',key:'fields[1].label'}

  @teacher
  Scenario: Create a template where the field label, wildcard and description are too long
    * def generateString = read('string-generator.js')
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
          "label": #(generateString()),
          "wildcard": #(generateString()),
          "description": #(generateString())
        }
      ]
    }
    """
    When method POST
    Then status 422
    And match response contains {message:'This field needs to be between 0 and 255 characters', key:'fields[1].label', value:'#ignore'}
    And match response contains {message:'This field needs to be between 0 and 255 characters', key:'fields[1].description', value:'#ignore'}
    And match response contains {message:'This field needs to be between 0 and 255 characters', key:'fields[1].wildcard', value:'#ignore'}
