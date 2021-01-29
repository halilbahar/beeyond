Feature: Custom application creation endpoint

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
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a custom application with no note
    Given request
    """
    {
      "content": "#(nginxDeployment)"
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a custom application with no content
    Given request { }
    When method POST
    Then status 422
    And match response contains { "message": "This field cannot be empty", "value": "", "key": "content" }

  @student
  Scenario: Create a custom application with blank content
    Given request
    """
    {
      "content": ""
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field cannot be empty", "value": "", "key": "content" }

  @student
  Scenario: Create a custom application with invalid kubernetes content

  @student
  Scenario: Create a custom application with a too long note
    * def generateString = read('string-generator.js')
    Given request
    """
    {
      "note": "#(generateString())",
      "content": "#(content)"
    }
    """
    When method POST
    Then status 422
    And match response contains { "message": "This field needs to be between 0 and 255 characters", "value": "#ignore", "key":"note"}

  @student
  Scenario: Create a custom application with id set

  @student
  Scenario: Create a custom application with application status set

  @student
  Scenario: Create a custom application with owner set

  @student
  Scenario: Create a custom application with created at set
