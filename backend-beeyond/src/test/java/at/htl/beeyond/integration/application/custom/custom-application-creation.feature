Feature: Custom application creation endpoint

  Background:
    * url baseUrl
    * path 'application/custom'

  @student
  Scenario: Create a valid custom application
    * def content = read('content-nginx-custom-application.yml.txt')
    Given request
    """
    {
      "note": "Nginx Deployment",
      "content": "#(content)",
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a custom application with no note
    * def content = read('content-nginx-custom-application.yml.txt')
    Given request
    """
    {
      "content": "#(content)"
    }
    """
    When method POST
    Then status 204

  @student
  Scenario: Create a custom application with no content
    * def content = read('content-nginx-custom-application.yml.txt')
    Given request
    """
    {}
    """
    When method POST
    Then status 422
    And match response contains [{"message":"This field cannot be empty","value":"","key":"content"}]

  @student
  Scenario: Create a custom application with blank content
    * def content = read('content-nginx-custom-application.yml.txt')
    Given request
    """
    {
      "content": ""
    }
    """
    When method POST
    Then status 422
    And match response contains [{"message":"This field cannot be empty","value":"","key":"content"}]

  @student
  Scenario: Create a custom application with a too long note
    * def content = read('content-nginx-custom-application.yml.txt')
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
    And match response contains [{"message":"This field needs to be between 0 and 255 characters","value":"#ignore","key":"note"}]