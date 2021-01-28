Feature: Template creation endpoint

  Background:
    * url baseUrl
    * path 'template'
    * configure headers = { Authorization: '#(auth(karate.tags))' }
    * def nginxDeployment = read('nginx-deployment-template.yml.txt')

  @teacher
  Scenario: Create a valid template application
    Given request
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
    When method POST
    Then status 201

  @teacher
  Scenario: Create a template with no description
    Given request
    """
    {
      "name": "Nginx Deployment",
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
    When method POST
    Then status 201

  @teacher
  Scenario: Create a template with no name and content
    Given request
    """
    {
      "description": "Static Webserver"
    }
    """
    When method POST
    Then status 422
    And match response contains { message: 'This field cannot be empty', key: 'name', value: '' }
    And match response contains { message: 'This field cannot be empty', key: 'content', value: '' }

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
    And match response contains {message:'This field needs to be between 1 and 255 characters', key: 'name', value:'#ignore'}
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
    And match response contains { message:'This field cannot be empty' ,value: '' ,key: 'wildcard' }
    And match response contains { message:'This field cannot be empty' ,value: '' ,key: 'label' }

  @teacher
  Scenario: Create a template where the field label and wildcard are empty
    Given request
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
          "label": "",
          "wildcard": "",
          "description": "This will be the port that will be exposed to the world"
        }
      ]
    }
    """
    When method POST
    Then status 422
    And match response contains { message: 'This field cannot be empty' , value: '' , key: 'wildcard' }
    And match response contains { message: 'This field cannot be empty' , value: '' , key: 'label' }

  @teacher
  Scenario: Create a template where the field label, wildcard and description are too long
    * def generateString = read('string-generator.js')
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
    And match response contains { message: 'This field needs to be between 0 and 255 characters', key:'label', value: '#ignore' }
    And match response contains { message: 'This field needs to be between 0 and 255 characters', key:'description', value: '#ignore' }
    And match response contains { message: 'This field needs to be between 0 and 255 characters', key:'wildcard', value: '#ignore' }
